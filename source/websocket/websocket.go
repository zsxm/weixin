package websocket

import (
	"errors"
	"weixin/source/pubnum/api"

	"github.com/zsxm/scgo/chttp"
	"github.com/zsxm/scgo/session"
	"github.com/zsxm/scgo/websocket"
)

type Connection struct {
	WS        *websocket.Conn // websocket 连接器
	Send      chan []byte     // 发送信息的缓冲 channel
	SessionId string          //存放连接的sessionId 多个
	Cpubnum   api.CachePubnum //存放公众号基本信息
}

type Message struct {
	SendId string //发送者ID
	ReceId string //接收者ID
	Value  []byte //发送消息
}

type hub struct {
	Connections  map[string]*Connection            // 注册的连接器
	Connections1 map[string]map[string]*Connection // 注册的连接器 key=userid value=[key=SessionId value=Connection]
	Message      chan Message                      // 连接器中发入的信息
	Register     chan *Connection                  // 连接器中注册请求
	Unregister   chan *Connection                  // 连接器中注销请求
}

var H = hub{
	Message:      make(chan Message),
	Register:     make(chan *Connection),
	Unregister:   make(chan *Connection),
	Connections:  make(map[string]*Connection),
	Connections1: make(map[string]map[string]*Connection),
}

func (c *Connection) reader() {
	for {
		buf := make([]byte, 100)
		n, err := c.WS.Read(buf)
		if err != nil {
			loger.Error("reader:", err)
			break
		}
		loger.Info("reader:", c.WS.Session().Id())
		cpubnum, err := pubnum(c.WS.Session()) //当前登录用户获得pubnum
		msg := Message{Value: buf[:n], ReceId: cpubnum.Userid, SendId: c.Cpubnum.Userid}
		H.Message <- msg
	}
	c.WS.Close()
	H.Unregister <- c
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

func (h *hub) run() {
	for {
		select {
		case c := <-h.Register:
			loger.Info("register key:", c.Cpubnum.Userid)
			if conns, ok := h.Connections1[c.Cpubnum.Userid]; ok {
				conns[c.SessionId] = c
			} else {
				conn := make(map[string]*Connection)
				conn[c.SessionId] = c
				h.Connections1[c.Cpubnum.Userid] = conn
			}
			for connsk, conns := range h.Connections1 {
				loger.Info("now connections keys:", connsk)
				for k, _ := range conns {
					loger.Info("\t", connsk, "-->connections key:", k)
				}
			}
			loger.Info("-------------Register-----------", h.Connections1)
		case c := <-h.Unregister:
			if conns, ok := h.Connections1[c.Cpubnum.Userid]; ok {
				for k, _ := range conns {
					if k == c.SessionId {
						delete(conns, k)
						close(c.Send)
						loger.Info("unregister userid=", c.Cpubnum.Userid, "sessionid=", k)
					}
				}
			}
			loger.Info("connection count:", len(h.Connections1))
			for connsk, conns := range h.Connections1 {
				loger.Info("\t", connsk, "-->connection count:", len(conns))
			}
			if len(h.Connections1[c.Cpubnum.Userid]) == 0 {
				delete(h.Connections1, c.Cpubnum.Userid)
			}
			loger.Info("-------------Unregister-----------", h.Connections1)
		case msg := <-h.Message:
			if conns, ok := h.Connections1[msg.ReceId]; ok {
				for _, conn := range conns {
					select {
					case conn.Send <- msg.Value:
						loger.Info("send message:", msg.ReceId)
					default:
						loger.Info("close unregister:", msg.ReceId)
						delete(conns, conn.SessionId)
						close(conn.Send)
					}
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
		c := &Connection{Send: make(chan []byte, 256), WS: ws, Cpubnum: cpubnum, SessionId: ws.Session().Id()}
		H.Register <- c
		go c.writer()
		c.reader()
	} else {
		id, _ := ws.Session().Get("id")
		cpubnum.Name = id
		c := &Connection{Send: make(chan []byte, 256), WS: ws, Cpubnum: cpubnum, SessionId: ws.Session().Id()}
		H.Register <- c
		go c.writer()
		msg := Message{Value: []byte(err.Error()), ReceId: cpubnum.Userid}
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
