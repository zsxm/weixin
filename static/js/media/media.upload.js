$(function(){
	//素材文件添加
	$(".dropdown-menu li").click(function(){
		var a=$(this).find("a");
		var c=a.attr("c");
		var t=a.attr("href");
		var ctype=a.attr("ctype");
		var copt=a.attr("copt");
		var action=a.attr("action");
		var page=t.substr(1,t.length).replace("-",".").replace("-",".");
		var url="/media/index";
		if(ctype=="news"&&copt=="list"){//图文列表
			
		}else{//图文添加和其它
			if(ctype=="news"){
				$("#previewEditorBtn").parent().show();
				$("#clearEditorBtn").parent().show();
			}else{
				$("#previewEditorBtn").parent().hide();
				$("#clearEditorBtn").parent().hide();
			}
			if(c=="true"||copt=="list"){
				a.attr("c","false");
				load(t,url,page,ctype,copt,action);
			}
		}
	});
	var load=function(t,url,page,ctype,copt,action){
		$(t).load(url+"?"+Math.random(),{page:page,action:action,saveType:1,ctype:ctype,copt:copt});
	}
	$("#defaultClick").click();
});