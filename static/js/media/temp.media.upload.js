$(function(){
	var url="/media/index";
	$(".dropdown-menu li").click(function(){
		var a=$(this).find("a");
		var c=a.attr("c");
		var t=a.attr("href");
		var ctype=a.attr("ctype");
		var copt=a.attr("copt");
		var action=a.attr("action");
		var page=t.substr(1,t.length).replace("-",".").replace("-",".");
		console.log(c,t,url,page,action);
		if(c=="true"||copt=="list"){
			a.attr("c","false");
			load(t,url,page,ctype,copt,action);
		}
	});
	var load=function(t,url,page,ctype,copt,action){
		$(t).load(url+"?"+Math.random(),{page:page,action:action,saveType:0,ctype:ctype,copt:copt},function(result){
			console.log(result);
		});
	}
	//load("#media-video-list",url,"media.video.list","video","list","")
	$("#defaultClick").click();
});