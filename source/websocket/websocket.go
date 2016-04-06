package websocket

import (
	"errors"
	"weixin/source/pubnum/api"

	"github.com/zsxm/scgo/chttp"
	"github.com/zsxm/scgo/session"
	"github.com/zsxm/scgo/websocket"
)

type Connection struct {
	WS      *websocket.Conn // websocket 连接器
	Send    chan []byte     // 发送信息的缓冲 channel
	Cpubnum api.CachePubnum
}

func (c *Connection) reader() {
	defer c.WS.Close()
	for {
		buf := make([]byte, 100)
		n, err := c.WS.Read(buf)
		if err != nil {
			loger.Error("reader", err)
			break
		}
		loger.Info("reader", c.WS.Session().Id())
		cpubnum, err := pubnum(c.WS.Session())
		msg := Message{Value: buf[:n], Key: cpubnum.Name}
		H.Message <- msg
	}
}

func (c *Connection) writer() {
	defer c.WS.Close()
	for msg := range c.Send {
		_, err := c.WS.Write(msg)
		if err != nil {
			loger.Error("writer", err)
			break
		}
		loger.Info("writer", c.WS.Session().Id())
	}
}

type Message struct {
	Key   string
	Value []byte
}

type hub struct {
	Connections map[string]*Connection // 注册的连接器
	Message     chan Message           // 连接器中发入的信息
	Register    chan *Connection       // 连接器中注册请求
	Unregister  chan *Connection       // 连接器中注销请求
}

var H = hub{
	Message:     make(chan Message),
	Register:    make(chan *Connection),
	Unregister:  make(chan *Connection),
	Connections: make(map[string]*Connection),
}

func (h *hub) run() {
	for {
		select {
		case c := <-h.Register:
			h.Connections[c.Cpubnum.Name] = c
			loger.Info("register key:", c.Cpubnum.Name)
			for k, _ := range h.Connections {
				loger.Info("connections keys:", k)
			}
		case c := <-h.Unregister:
			if _, ok := h.Connections[c.Cpubnum.Name]; ok {
				loger.Info("unregister", c.Cpubnum.Name)
				delete(h.Connections, c.Cpubnum.Name)
				close(c.Send)
			}
			loger.Info("connection count", len(h.Connections))
		case msg := <-h.Message:
			c := h.Connections[msg.Key]
			if c != nil {
				select {
				case c.Send <- msg.Value:
					loger.Info("send message", c.Cpubnum.Name)
				default:
					loger.Info("close unregister", c.Cpubnum.Name)
					delete(h.Connections, c.Cpubnum.Name)
					close(c.Send)
				}
			}
		}
	}
}

func init() {
	control.Add("/websocket/to", toPage)
	control.AddWS("/socket/conn", webSocketConnServer)
	go H.run()
}

func toPage(c chttp.Context) {
	c.HTML("/websocket/websocket", nil)
}

func webSocketConnServer(ws *websocket.Conn) {
	cpubnum, err := pubnum(ws.Session())
	if err == nil {
		c := &Connection{Send: make(chan []byte, 256), WS: ws, Cpubnum: cpubnum}
		H.Register <- c
		defer func() {
			H.Unregister <- c
		}()
		go c.writer()
		c.reader()
	} else {
		id, _ := ws.Session().Get("id")
		cpubnum.Name = id
		c := &Connection{Send: make(chan []byte, 256), WS: ws, Cpubnum: cpubnum}
		H.Register <- c
		defer func() {
			H.Unregister <- c
		}()
		go c.writer()
		msg := Message{Value: []byte(err.Error()), Key: cpubnum.Name}
		H.Message <- msg
		loger.Error("webSocketConnServer:", err.Error())
	}
}

func pubnum(session session.Interface) (api.CachePubnum, error) {
	var cpubnum api.CachePubnum
	smp, err := session.GetMap()
	if err != nil {
		return cpubnum, errors.New("未登录")
	}
	userid := smp.Get("id")
	if userid == "" {
		return cpubnum, errors.New("未登录")
	}
	pubnumid := api.GetCachePubNumId(userid)
	if pubnumid == "" {
		return cpubnum, errors.New("请启用一个公众号")
	}
	cpubnum = api.CachePubNum(pubnumid)
	if cpubnum.Name == "" {
		return cpubnum, errors.New("请获取启用的公众号Token")
	}
	return cpubnum, nil
}
