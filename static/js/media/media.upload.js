$(function(){
	
	var loadsOk=[true,true,true,true,true];
	$(".nav-tabs li").click(function(){
		var a=$(this).find("a");
		var t=a.attr("href");
		var i=a.attr("index");
		var url="/media/index";
		var page,action="/media/temp/upload";
		if(t=="#panel-news"&&loadsOk[i]){
			page="media.news";
			action="/media/upload";
		}else if(t=="#panel-video"&&loadsOk[i]){
			page="media.video";
			action="/media/release/voice";
		}else if(t=="#panel-image"&&loadsOk[i]){
			page="media.image";
			action="/media/release/common";
		}else if(t=="#panel-voice"&&loadsOk[i]){
			page="media.voice";
			action="/media/release/common";
		}else if(t=="#panel-thumb"&&loadsOk[i]){
			page="media.thumb";
			action="/media/release/common";
		}else{
			return;
		}
		loads(t,url,page,action,i);
	});
	var loads=function(t,url,page,action,i){
		$(t).load(url+"?"+Math.random(),{page:page,action:action,saveType:1},function(d){
			loadsOk[i]=false;
		});
	}
	loads("#panel-news","/media/index?"+Math.random(),"media.news","/media/upload",0);
});