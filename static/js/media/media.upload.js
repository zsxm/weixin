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
		console.log(ctype);
		if(ctype=="news"&&copt=="list"){
			
		}else{
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