var newsTemplates={
	image:'{{#each data}}<div id="{{mediaId}}" src="{{localName}}" onclick="selectMedia(this,\'image\',\'{{../st}}\')" title="选择" class="file-preview-frame" style="width:100px;height:100px;cursor:pointer;">'
			+'	<img src="/{{localName}}" class="file-preview-image" style="width:100px;height:100px;">'
			+'</div>{{/each}}',
	video:'{{#each data}}<div id="{{mediaId}}" src="{{localName}}" onclick="selectMedia(this,\'video\',\'{{../st}}\')" title="选择" class="file-preview-frame" style="width:100px;height:100px;cursor:pointer;">'
			+'	<video width="100px" height="100px" controls="">'
			+'		<source src="/{{localName}}" type="video/mp4">'
			+'	</video>'
			+'</div>{{/each}}',
	voice:'{{#each data}}<div id="{{mediaId}}" src="{{localName}}" onclick="selectMedia(this,\'voice\',\'{{../st}}\')" title="选择" class="file-preview-frame" style="width:100px;height:100px;cursor:pointer;">'
			+'	<audio controls="">'
			+'		<source src="/{{localName}}" type="audio/wav">'
			+'	</audio>'
			+'</div>{{/each}}'
};
var selectMedia=function(o,ctype,st){
	o=$(o);
	var mediaId=o.attr("id");
	var id=$("#activityForm").val();
	var src=o.attr("src");
	if(st=="cover"){//封面选择
		$("#thumb_media_id_"+id).val(mediaId);//封面id
		$("#newsThumb_"+id).attr("src","/"+src);
	}else if(st=="content"){//内容
		var content="content_"+id;
		//CKEDITOR.instances[content].document.getBody().setHtml('<img src="/'+src+'" border="0" title="Hello" />');//设置html代码 覆盖所有
		var html="";
		if(ctype=="image"){
			html='<img src="/'+src+'"/>';
		}else if(ctype=="video"){
			html='<video controls=""><source src="/'+src+'" type="video/mp4"></video>';
		}else if(ctype=="voice"){
			html='<audio controls=""><source src="/'+src+'" type="audio/wav"></audio>';
		}
		
        UE.getEditor(content).execCommand('insertHtml', html)
		//var element = CKEDITOR.dom.element.createFromHtml(html);
		//CKEDITOR.instances[content].insertElement(element);//插入html代码 追加
	}
	$("#modal-container-mediaselect").modal("hide");//隐藏素材容器
	//console.log(mediaId,id);
}
//封面 素材库选择按钮
var newsMediaSelect=function(o){
	var o=$(o);
	var ctype=o.attr("ctype");
	lineGetList("cover",ctype)
}
//获得 mediatype类型的素材
var lineGetList=function(st,ctype){
	var url="/media/line/get";
	$.get(url,{mediaType:ctype},function(result){
		result.st=st;
		//console.log(result);
		if(result.code=="0"){
			var template;
			if(ctype=="image"){
				template = Handlebars.compile(newsTemplates.image);
			}else if(ctype=="video"){
				template = Handlebars.compile(newsTemplates.video);
			}else if(ctype=="voice"){
				template = Handlebars.compile(newsTemplates.voice);
			}
			var content = template(result);
			$("#mediaSelect").html(content);
			$("#modal-container-mediaselect").modal("show");//显示素材容器
		}else{
			$.alertmsg("#tipsMsg","danger",result.codemsg);
		}
	});
}
$(function(){
	$("#leftBar_video,#leftBar_image,#leftBar_voice").click(function(){
		var o=$(this);
		var ctype=o.attr("ctype");
		lineGetList("content",ctype);
	});
})