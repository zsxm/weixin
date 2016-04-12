$(function(){
	var loadsOk=[true,true,true,true];
	$(".nav-tabs li").click(function(){
		var a=$(this).find("a");
		var t=a.attr("href");
		var i=a.attr("index");
		var url="/media/index";
		var page,action="/media/temp/release";
		if(t=="#panel-video"&&loadsOk[i]){
			page="media.video";
		}else if(t=="#panel-image"&&loadsOk[i]){
			page="media.image";
		}else if(t=="#panel-voice"&&loadsOk[i]){
			page="media.voice";
		}else if(t=="#panel-thumb"&&loadsOk[i]){
			page="media.thumb";
		}else{
			return;
		}
		loads(t,url,page,action,i);
	});
	var loads=function(t,url,page,action,i){
		$(t).load(url+"?"+Math.random(),{page:page,action:action,saveType:0},function(d){
			loadsOk[i]=false;
		});
	}
	loads("#panel-video","/media/index?"+Math.random(),"media.video","/media/temp/release",0);
});