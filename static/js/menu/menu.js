var menu={
	leve1:3,
	leve2:5,
	leve1TxtMax:4,
	leve2TxtMax:8,
	init:function(){
		var url="/menu/get";
		$.get(url,{},function(result){//获得菜单
			if(result.code=="0"){
				var menus=result.menu;
				var btns=menus.button;
				for(var i in btns){
					var btn=btns[i];
					var name=menu.isBlank(btn.name)?"":btn.name;
					var type=menu.isBlank(btn.type)?"":btn.type;
					var key=menu.isBlank(btn.key)?"":btn.key;
					var url=menu.isBlank(btn.url)?"":btn.url;
					var subBtn=btn.sub_button
					if(btn.type==undefined&&subBtn.length>0){
						var id=menu.genId();
						menu.genLeve1Btn(id,name,type,key,url);
						for(var j in subBtn){
							var sbtn=subBtn[j];
							console.log(sbtn)
							name=menu.isBlank(sbtn.name)?"":sbtn.name;
							type=menu.isBlank(sbtn.type)?"":sbtn.type;
							key=menu.isBlank(sbtn.key)?"":sbtn.key;
							url=menu.isBlank(sbtn.url)?"":sbtn.url;
							menu.genLeve2Btn(id,menu.genId(),this,name,type,key,url);
						}
					}else{
						menu.genLeve1Btn(menu.genId(),name,type,key,url);
					}
					
					if(menu.leve1BtnsSize()==menu.leve1){
						$("#defaultBtn").hide();
					}
				}
			}else{
				alert(result.codemsg);
			}
		});
		$("#defaultBtn").click(function(){//创建菜单按钮
			var id=menu.genId();
			menu.genLeve1Btn(menu.genId(),"菜单名称","","","");
			if(menu.leve1BtnsSize()==menu.leve1){
				$(this).hide();
			}
		});
		menu.ctrolOffOn("-1");
		menu.deleteMenu();
		menu.nameBlur();
		menu.menuTypeChange();
		menu.urlBlur();
		menu.keyBlur();
		menu.saveBtn();
	},
	genId:function(){
		return Math.random().toString(36).substr(2);
	},
	isBlank:function(v){
		if(v==undefined||v=="null"||v==null||v==""){
			return true;
		}
		return false;
	},
	genLeve1Btn:function(id,name,menutype,key,url){//生成一级菜单
		var btns=$("#btns");
		var arrts="";
		arrts+=menu.isBlank(name)?"":"name='"+name+"' ";
		arrts+=menu.isBlank(menutype)?"":"menutype='"+menutype+"' ";
		arrts+=menu.isBlank(key)?"":"key='"+key+"' ";
		arrts+=menu.isBlank(url)?"":"url='"+url+"' ";
		
		var html='<div id="pmenu_'+id+'" class="btn-group dropup open">'
					+'<button id="leve1_'+id+'" class="btn btn-default" onclick="menu.leve1Edit(\''+id+'\')" '+arrts+'>'+name+'</button>&nbsp;&nbsp;'
					+'<button data-toggle="dropdown" class="btn btn-default dropdown-toggle"><span class="caret"></span></button>'
					+'<ul id="subList_'+id+'" class="dropdown-menu">'
					+'	<li id="createBtn" onclick="menu.genLeve2Btn(\''+id+'\',menu.genId(),this,\'子菜单名称\',\'\',\'\',\'\')"><a href="javascript:void(0)">ADD+</a></li>'
					+'</ul>'
				+'</div>';
		btns.append(html);
		menu.leve1Edit(id);
	},
	genLeve2Btn:function(pid,id,o,name,menutype,key,url){//生成二级菜单
		var arrts="";
		arrts+=menu.isBlank(name)?"":"name='"+name+"' ";
		arrts+=menu.isBlank(menutype)?"":"menutype='"+menutype+"' ";
		arrts+=menu.isBlank(key)?"":"key='"+key+"' ";
		arrts+=menu.isBlank(url)?"":"url='"+url+"' ";
		var html='<li id="leve2_'+id+'" class="leve2btn" onclick="menu.leve2Edit(\''+pid+'\',\''+id+'\')">'
				+'	<a id="tleve2_'+id+'" href="javascript:void(0)" '+arrts+'>'+name+'</a>'
				+'</li>';
		var subList=$("#subList_"+pid);
		var createBtn=subList.find("#createBtn");
		createBtn.before(html);
		if(menu.leve2BtnsSize(pid)==menu.leve2){
			subList.find("#createBtn").hide();
		}else{
			createBtn.before('<li id="dleve2_'+id+'" class="divider"></li>');
		}
		menu.leve2Edit(pid,id);
		menu.removeLeve1Attr(pid);
	},
	leve1BtnsSize:function(){//一级菜单数量
		return $("#btns .dropup").length;
	},
	leve2BtnsSize:function(id){//某一个一级菜单下二级菜单数量
		var subList=$("#subList_"+id)
		return subList.find(".leve2btn").length;
	},
	leve1Edit:function(id){//一级编辑
		if(menu.leve2BtnsSize(id)>0){
			menu.ctrol().hide();
		}else{
			menu.ctrol().show();
		}
		var attrs=menu.getAttrs("leve1_"+id);
		menu.setCtrol(id,attrs.name,attrs.type,attrs.url,attrs.key);
		menu.setDeleteMenu("",id,"1","删除菜单");
		$("#name").attr("maxlength",menu.leve1TxtMax);
	},
	leve2Edit:function(pid,id){//二级编辑
		menu.ctrol().show();
		var attrs=menu.getAttrs("tleve2_"+id);
		menu.setCtrol(id,attrs.name,attrs.type,attrs.url,attrs.key);
		menu.setDeleteMenu(pid,id,"2","删除子菜单");
		$("#name").attr("maxlength",menu.leve2TxtMax);
	},
	getAttrs:function(id){//获得属性值 return json
		var o=$("#"+id);
		var name=o.attr("name");
		name=name==undefined?"":name;
		var menutype=o.attr("menutype");
		menutype=menutype==undefined||menutype==""?"-1":menutype;
		var url=o.attr("url");
		url=url==undefined?"":url;
		var key=o.attr("key");
		key=key==undefined?"":key;
		var json={"name":name,"type":menutype,"url":url,"key":key};
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
	resetCtrol:function(){//重置
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
	setDeleteMenu:function(pid,id,level,txt){//设置删除菜单按钮
		var btn=$("#deleteMenuBtn");
		btn.text(txt);
		btn.attr({"menuid":id,"level":level,"pid":pid});
	},
	deleteMenu:function(){//删除菜单按钮
		$("#deleteMenuBtn").click(function(){
			var btn=$(this);
			var menuid=btn.attr("menuid");
			var level=btn.attr("level");
			var pid=btn.attr("pid");
			var id=btn.attr("menuid");
			$("#leve2_"+menuid).remove();
			$("#dleve2_"+menuid).remove();
			$("#pmenu_"+menuid).remove();
			btn.removeAttr("menuid");
			menu.resetCtrol();
			menu.showAddBtn(level,pid,id);
		});
	},
	showAddBtn:function(level,pid,id){//显示添加按钮
		if(level=="1"){
			if(menu.leve1BtnsSize()==menu.leve1-1){
				$("#defaultBtn").show();
			}
		}else if(level=="2"){
			if(menu.leve2BtnsSize(pid)==menu.leve2-1){
				var subList=$("#subList_"+pid);
				var createBtn=subList.find("#createBtn");
				var prevObj=createBtn.prev();
				if(prevObj.attr("class")!="divider"){
					var prevId=prevObj.attr("id").split("_")[1];
					createBtn.before('<li id="dleve2_'+prevId+'" class="divider"></li>');
				}
				createBtn.show();
			}
		}
	},
	nameBlur:function(){
		$("#name").blur(function(){
			var o=$(this);
			var menuid=o.attr("menuid");
			$("#leve1_"+menuid).text(o.val());
			$("#leve1_"+menuid).attr("name",o.val());
			$("#tleve2_"+menuid).text(o.val());
			$("#tleve2_"+menuid).attr("name",o.val());
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
			case "scancode_waitmsg":
			case "scancode_push":
			case "pic_sysphoto":
			case "pic_photo_or_album":
			case "pic_weixin":
			case "location_select":
//			case "media_id":
//			case "view_limited":
				$("#divKey").show();
				$("#divUrl").hide();
				menu.removeMenuForAttr("leve1_"+menuid,"url");
				menu.removeMenuForAttr("tleve2_"+menuid,"url");
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
	fjson:function(json){//过滤菜单json属性判断类型设置相应参数 return new json
		var jsn={};
		jsn.type=json.type;
		jsn.name=json.name;
		var type=json.type;
		if(type=="click"){
			jsn.key=json.key;
		}else if(type=="view"){
			jsn.url=json.url;
		}else if(type=="scancode_waitmsg"||type=="scancode_push"||type=="pic_sysphoto"||type=="pic_photo_or_album"||type=="pic_weixin"||type=="location_select"){
			jsn.key=json.key;
			jsn.sub_button=[];
		}
		return jsn;
	},
	saveBtn:function(){//保存
		//获得所有配置的菜单，封装为json数据，请求发送到后台
		//未验证填写的数据或选择的类型
		$("#saveBtn").click(function(){
			var btns=$("#btns .dropup");//所有一级菜单
			var menus={};
			menus.button=new Array();
			btns.each(function(i,v){
				var o=$(v)
				var pid=o.attr("id").split("_")[1];
				var btnList=o.find("#subList_"+pid+" .leve2btn a");
				var btnLen=btnList.length;
				if(btnLen==0){
					var attrs=menu.getAttrs("leve1_"+pid);
					var jsn=menu.fjson(attrs);
					menus.button.push(jsn);
				}else{
					var sbtnsList={};
					var name=$("#leve1_"+pid).attr("name");
					var sbtns=new Array();
					btnList.each(function(i,v){
						var o=$(v)
						var id=o.attr("id").split("_")[1];
						var attrs=menu.getAttrs("tleve2_"+id);
						var jsn=menu.fjson(attrs);
						sbtns.push(jsn);
					});
					sbtnsList.name=name;
					sbtnsList.sub_button=sbtns;
					menus.button.push(sbtnsList);
				}
			});
			var data=JSON.stringify(menus);
			
			console.log(data);
			var url="/menu/save";
			$.post(url,{data:data},function(result){
				console.log(result);
				if(result.code=="0"){
					alert("创建成功");
				}else{
					alert(result.codemsg);
				}
			})
		});
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