var menu={
	leve1:3,
	leve2:5,
	leve1TxtMax:4,
	leve2TxtMax:8,
	init:function(){
		var menus=$("#btns .dropup");
		//微信接口查询菜单,如果没有创建菜单显示默认的新创建按钮
		//if(length==0){
			$("#defaultBtn").click(function(){
				menu.genLeve1Btn();
				if(menu.leve1BtnsSize()==3){
					$(this).hide();
				}
			});
		//}
		menu.ctrolOffOn("-1");
		menu.deleteMenu();
		menu.nameBlur();
		menu.menuTypeChange();
		menu.urlBlur();
		menu.keyBlur();
		menu.saveBtn();
	},
	genLeve1Btn:function(){//生成一级菜单
		var btns=$("#btns");
		var id=new Date().getTime();
		var html='<div id="pmenu_'+id+'" class="btn-group dropup open">'
					+'<button id="leve1_'+id+'" class="btn btn-default" onclick="menu.leve1Edit($(this))">菜单名称</button>&nbsp;&nbsp;'
					+'<button data-toggle="dropdown" class="btn btn-default dropdown-toggle"><span class="caret"></span></button>'
					+'<ul class="dropdown-menu">'
					+'	<li onclick="menu.genLeve2Btn(\''+id+'\',this)"><a href="javascript:void(0)">ADD+</a></li>'
					+'</ul>'
				+'</div>';
		btns.append(html);
		menu.leve1Edit($("#leve1_"+id));
	},
	genLeve2Btn:function(pid,o){//生成二级菜单
		var obj=$(o);
		var id=new Date().getTime();
		var html='<li id="leve2_'+id+'" class="leve2btn" onclick="menu.leve2Edit($(this))"><a id="tleve2_'+id+'" href="javascript:void(0)">子菜单名称</a></li>';
		obj.before(html);
		if(menu.leve2BtnsSize(o)==5){
			obj.remove();
		}else{
			obj.before('<li id="dleve2_'+id+'" class="divider"></li>');
		}
		menu.leve2Edit($("#leve2_"+id));
		menu.removeLeve1Attr(pid);
	},
	leve1BtnsSize:function(){//一级菜单数量
		return $("#btns .dropup").length;
	},
	leve2BtnsSize:function(o){//某一个一级菜单下二级菜单数量
		var obj=$(o);
		return obj.parent().find(".leve2btn").length;
	},
	leve1Edit:function(o){//一级编辑
		var obj=$(o);
		if(menu.leve2BtnsSize(o)>0){
			menu.ctrol().hide();
		}else{
			menu.ctrol().show();
		}
		var id=obj.parent().attr("id").split("_")[1];
		var attrs=menu.getAttrs("leve1_"+id);
		menu.setCtrol(id,obj.text(),attrs.menutype,attrs.url,attrs.key);
		menu.setDeleteMenu(id,"删除菜单");
	},
	leve2Edit:function(o){//二级编辑
		var obj=$(o);
		var id=obj.attr("id").split("_")[1];
		menu.ctrol().show();
		var attrs=menu.getAttrs("tleve2_"+id);
		menu.setCtrol(id,obj.find("a").text(),attrs.menutype,attrs.url,attrs.key);
		menu.setDeleteMenu(id,"删除子菜单");
	},
	getAttrs:function(id){//获得属性值 return json
		var o=$("#"+id);
		var menutype=o.attr("menutype");
		menutype=menutype==undefined||menutype==""?"-1":menutype;
		var url=o.attr("url");
		url=url==undefined?"":url;
		var key=o.attr("key");
		key=key==undefined?"":key;
		var json={"menutype":menutype,"url":url,"key":key};
		return json;
	},
	setCtrol:function(menuid,name,type,url,key){//设置一些默认值
		var nameo=$("#name");
		nameo.val(name);
		nameo.attr("menuid",menuid);
		$("#menuType").attr("menuid",menuid);
		var mtselect=document.getElementById("menuType");
		for(var i=0;i<mtselect.length;i++){
			if(mtselect[i].value==type){
				mtselect[i].selected=true;
			}
		}
		var urlo=$("#url");
		urlo.val(url);
		urlo.attr("menuid",menuid);
		var keyo=$("#key");
		keyo.val(key);
		keyo.attr("menuid",menuid);
		menu.ctrolOffOn(menuid,type);
	},
	resetCtrol:function(){//行署
		menu.setCtrol("","",-1,"","");
	},
	removeLeve1Attr:function(id){//移除一级菜单所有设置属性
		id="leve1_"+id;
		menu.removeMenuForAttr(id,"menutype")
		menu.removeMenuForAttr(id,"url")
		menu.removeMenuForAttr(id,"key")
	},
	removeMenuForAttr:function(id,attr){//移除菜单指定设置的属性
		var m=$("#"+id);
		m.removeAttr(attr);
	},
	ctrol:function(){
		return $("#ctrol");
	},
	setDeleteMenu:function(id,txt){//设置删除菜单按钮
		var btn=$("#deleteMenuBtn");
		btn.text(txt);
		btn.attr("menuid",id);
	},
	deleteMenu:function(){//删除菜单按钮
		$("#deleteMenuBtn").click(function(){
			var btn=$("#deleteMenuBtn");
			var menuid=btn.attr("menuid");
			$("#leve2_"+menuid).remove();
			$("#dleve2_"+menuid).remove();
			$("#pmenu_"+menuid).remove();
			menu.resetCtrol();
		});
	},
	nameBlur:function(){
		$("#name").blur(function(){
			var o=$(this);
			var menuid=o.attr("menuid");
			$("#leve1_"+menuid).text(o.val());
			$("#tleve2_"+menuid).text(o.val());
		});
	},
	menuTypeChange:function(){
		$("#menuType").change(function(){
			var o=$(this);
			var menuid=o.attr("menuid");
			$("#leve1_"+menuid).attr("menutype",o.val());
			$("#tleve2_"+menuid).attr("menutype",o.val());
			menu.ctrolOffOn(menuid,o.val());
		});
	},
	urlBlur:function(){
		$("#url").blur(function(){
			var o=$(this);
			var menuid=o.attr("menuid");
			$("#leve1_"+menuid).attr("url",o.val());
			$("#tleve2_"+menuid).attr("url",o.val());
		});
	},
	keyBlur:function(){
		$("#key").blur(function(){
			var o=$(this);
			var menuid=o.attr("menuid");
			$("#leve1_"+menuid).attr("key",o.val());
			$("#tleve2_"+menuid).attr("key",o.val());
		});
	},
	ctrolOffOn:function(menuid,menuTypeValue){//控制显示隐藏表单控件,并重置被引藏的控制值
		switch(menuTypeValue){
			case "click":
				$("#divKey").show();
				$("#divUrl").hide();
				menu.removeMenuForAttr("leve1_"+menuid,"url");
				menu.removeMenuForAttr("tleve2_"+menuid,"url");
			break;
			case "view":
				$("#divUrl").show();
				$("#divKey").hide();
				menu.removeMenuForAttr("leve1_"+menuid,"key");
				menu.removeMenuForAttr("tleve2_"+menuid,"key");
			break;
			default:
				$("#divUrl").hide();
				$("#divKey").hide();
				menu.removeMenuForAttr("leve1_"+menuid,"key");
				menu.removeMenuForAttr("tleve2_"+menuid,"key");
				menu.removeMenuForAttr("leve1_"+menuid,"url");
				menu.removeMenuForAttr("tleve2_"+menuid,"url");
		}
	},
	saveBtn:function(){//保存,获得所有配置的菜单，封装为json数据，请求发送到后台
		
	}
	
};
menu.init();

$(function(){
	var url="/menu/type/list";
	var pubnum_temp='<option value="-1">请选择</option>{{#each .}}<option value="{{value}}">{{name}}</option>{{/each}}';
	$.get(url,{r:"ajax"},function(result){
		var template = Handlebars.compile(pubnum_temp);
		var content = template(result);
		$("#menuType").html(content);
	});
})