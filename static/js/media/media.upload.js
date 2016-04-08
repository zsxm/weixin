$(function(){
	
	var loadsOk=[true,true,true,true,true];
	$(".nav-tabs li").click(function(){
		var a=$(this).find("a");
		var t=a.attr("href");
		var i=a.attr("index");
		var url="/media/"
		if(t=="#panel-news"&&loadsOk[i]){
			url+="media.news.html";
		}else if(t=="#panel-video"&&loadsOk[i]){
			url="media.video.html";
		}else if(t=="#panel-image"&&loadsOk[i]){
			url="media.image.html";
		}else if(t=="#panel-voice"&&loadsOk[i]){
			url="media.voice.html";
		}else if(t=="#panel-thumb"&&loadsOk[i]){
			url="media.thumb.html";
		}else{
			return;
		}
		loads(t,url,i);
	});
	var loads=function(t,url,i){
		$(t).load(url+"?"+Math.random(),{},function(d){
			loadsOk[i]=false;
		});
	}
	loads("#panel-news","/media/media.news.html?"+Math.random(),0);
});