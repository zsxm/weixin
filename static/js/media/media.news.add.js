var showNewsForm=function(t,id){
	$(".form-horizontal").hide();
	$("#newsForm_"+id).fadeIn(300);
	if(t!="add"){
		$("html,body").animate({scrollTop:$("#tipsMsg").offset().top},500);
	}
}
var deleteNewsForm=function(id){
	if(newsFormLen()>1){
		$(".form-horizontal").hide();
		var otherObj;
		otherObj=$("#newsForm_"+id).prev();
		console.log(otherObj.length);
		if(otherObj.length==0){
			otherObj=$("#newsForm_"+id).next();
		}
		$("#newsForm_"+id).remove();
		$("#_thumb_"+id).remove();
		$(otherObj).fadeIn(300);
	}
}
var newsFormLen=function(){
	return $(".form-horizontal").length;
}
$(function(){
	//CKEDITOR.config.readOnly = true;//禁用或启用ckeditor
	//CKEDITOR.instances.content.setData("正文");//设置内容
	//CKEDITOR.instances.content.getData();//获取内容
	//CKEDITOR.instances.content.document.getBody().getText();//获得纯文本

	
	var newsFormEditor=function(id){
		var content='content_'+id;
		var editor1=CKEDITOR.replace(content,{
			toolbar:[
				['Font','Bold','Underline','Italic','TextColor','RemoveFormat','FontSize']
			]
		});
		var defaultText="正文";
		CKEDITOR.instances[content].on("instanceReady",function(){
			this.document.on("click", function(){
				var txt=this.getBody().getText();
				if($.trim(txt)==defaultText){
					this.getBody().setText("");
				}
			});
			this.on("blur",function(){
				var txt=this.document.getBody().getText();
				var html=this.document.getBody().getHtml();
				if($.trim(txt)==""&&html.indexOf("img")==-1){
					this.document.getBody().setHtml("<h1 style='color:#888888'>"+defaultText+"</h1>");
				}
			});
			this.document.getBody().setHtml("<h1 style='color:#888888'>"+defaultText+"</h1>");
		});
		//插入内容
		CKEDITOR.instances[content].on("instanceReady",function(){
			//this.document.getBody().setHtml('<img src="/static/image/login/1.jpg" border="0" title="Hello" />');
		});
	}

	//上传控件
	var newsFormFileInput=function(id){
		var url="/media/upload";
		$("#newsFile_"+id).fileinput({
			uploadUrl: url,//上传地址
			enctype: 'multipart/form-data',
			overwriteInitial: true,//重新选择文件后是否覆盖 initialPreview
			language:"zh",//设置语言
			previewFileType:["image"],//只选择image类型文件
			allowedFileExtensions:["jpg","jpeg", "png", "gif"],
	        initialCaption: "本地上传封面",
			dropZoneTitle:"封面拖拽到这里...",
			browseLabel: "本地上传封面",//选择按钮文字 
			//isPreviewable:false,//关闭预览
			//removeLabel: '删除所有',//移除按钮文字
			showRemove:false//隐藏移动按钮
		}).on("fileuploaded",function(event,data,prvid,index){
			var result=data.response;
			console.log(id,result);
			if (result.code=="0") {
				var da=result.data;
				var dir=da.DirName;
				var fn=da.FileNameId[0];
				var filePath=dir+"/"+fn;
				$("#newsFilePath_"+id).val(filePath);
				$("#newsThumb_"+id).attr("src","/"+filePath);	
				$("#thumb_media_id_"+id).val("");//封面图片mediaId
			}else{
				$.alertmsg("#tipsMsg","danger",result.codemsg);
			}
		}).on("fileremoved",function(event,prvid,index){//未上传时点击删除
			console.log("fileremoved",prvid);
			$("#newsFilePath_"+id).val("");
			$("#thumb_media_id_"+id).val("");
		}).on("filesuccessremove",function(event,prvid,index){//上传成功后点击删除
			console.log("removed",prvid);
			$("#newsFilePath_"+id).val("");
			$("#thumb_media_id_"+id).val("");
		});
	}
	//媒体新闻form验证
	var newsFormValidate=function(id){
		$("#newsForm_"+id).validate({
			rules:{
				title:{
					required:true
				}
			},
			messages:{
				title:{
					required:"标题不能为空且长度不能超过64字"
				}
			},
			beforSubmit:function(a){
				console.log("aaaaaaaaa",a);
			},
			submitHandler:function(form){
				var content='content_'+id;
				var load=$.loadding.New("正在保存请稍候...", 6, -1, "#videoReleaseBtn");
				var text=CKEDITOR.instances[content].document.getBody().getText();
				text=$.trim(text);
				console.log(text);
				if(text==""||text=="正文"||text.length>20000){
					$.alertmsg("#tipsMsg","danger","正文不能为空且长度不能超过20000字");
					return ;
				}
//				load.Show();
//				$("#newsForm_"+id).ajaxSubmit({
//					dataType : "json",
//					success : function(result){
//						load.Hide();
//						if(result.code=="0"){
//							$.alertmsg("#tipsMsg","success","素材保存成功");
//						}else{
//							$.alertmsg("#tipsMsg","danger",result.codemsg);
//						}
//					}
//				});
			}
		});
	}
	var newsFormTemp = '{{#each .}}'
					//+'<form id="newsForm_{{id}}" action="{{action}}" method="post" class="form-horizontal" role="form">'
				+'<div id="newsForm_{{id}}" class="form-horizontal">'
					+'<div class="form-group">'
				    +'  	<div class="col-sm-4">'
				    +'     		<input type="text" class="form-control" id="newsTitle_{{id}}" name="title" placeholder="标题">'
				    +'  	</div>'
				    +'  	<div class="col-sm-4">'
				    +'    		<input type="text" class="form-control" id="newsAuthor_{{id}}" name="author" placeholder="作者">'
				    +'  	</div>'
				    +'  	<div>'
					+'			<input id="show_cover_pic_{{id}}" name="show_cover_pic" type="checkbox"/>'
				    +'     		<lable>是否显示封面</lable>'
				    +'  	</div>'
				   	+'</div>'
					+'<div class="form-group">'
				    +'  	<div class="col-sm-12">'
				    +'     		<input type="text" class="form-control" id="contentSourceUrl_{{id}}" name="content_source_url" placeholder="原文链接">'
				    +'  	</div>'
				   	+'</div>'
					+'<div class="form-group">'
				    +'  	<div class="col-sm-12">'
					+'			<textarea id="content_{{id}}" name="content"></textarea>'
				    +'  	</div>'
				   	+'</div>'
					+'<div class="form-group">'
				    +'  	<div class="col-sm-12">'
				    +'     		<input type="file" id="newsFile_{{id}}" name="file" value="" accept="image/*" class="file-loading"/>'
					+'			<input type="hidden" name="ctype" value="news"/>'
					+'			<input type="hidden" id="newsFilePath_{{id}}" value=""/>'
					+'			<input type="hidden" id="thumb_media_id_{{id}}" name="thumb_media_id" value=""/>'
					+'			<input type="hidden" name="saveType" value="{{saveType}}"/>'
				    +'  	</div>'
				   	+'</div>'
			    	//+'</form>'
				+'</div>'
					+'{{/each}}';
	
	var newsThumbTemp='{{#each .}}<div id="_thumb_{{id}}" class="file-preview-frame" style="width:100px;height:100px;">'
				+'	<img id="newsThumb_{{id}}" src="" class="file-preview-image" '
				+'	style="width:100px;height:100px;background-image: url(/static/image/comm/numbers.jpg) 0 300px no-repeat;">'
				+'	<div class="file-thumbnail-footer">'
				+'		<div id="titleMsg_{{id}}"></div>'
				+'	</div>'
				+'	<div class="file-actions">'
				+'		<div class="file-footer-buttons">'
				+'			<button id="{{id}}" onclick="showNewsForm(\'edit\',\'{{id}}\')" type="button" class="kv-file-edit btn btn-xs btn-default" title="编辑"><i class="glyphicon glyphicon-edit text-info">编辑</i></button>'
				+'			<button id="{{id}}" onclick="deleteNewsForm(\'{{id}}\')" type="button" class="kv-file-remove btn btn-xs btn-default" title="删除"><i class="glyphicon glyphicon-trash text-danger">删除</i></button>'
				+'		</div>'
				+'	</div>'
				+'</div>{{/each}}';
	var newsFormTitleKeyUp=function(id){
		$("#newsTitle_"+id).keydown(function(){
			$("#titleMsg_"+id).html($(this).val());
		});
		$("#newsTitle_"+id).keyup(function(){
			$("#titleMsg_"+id).html($(this).val());
		});
	}
	var createNewsForm=function(){
		var id=new Date().getTime();
		var formLen=newsFormLen();
		if(formLen<8){
			var data=[{"id":id,"action":newsAction,"saveType":newsSaveType}];
			//生成form
			var template = Handlebars.compile(newsFormTemp);
			var content = template(data);
			$("#_newsForms").append(content);
			//alert($("#_newsForms").html());
			$('#show_cover_pic_'+id).checkboxpicker();
			newsFormFileInput(id);//生成inputfile
			//newsFormValidate(id);//生成验证
			newsFormEditor(id);//生成编辑内容
			newsFormTitleKeyUp(id);//title添加keyup事件
			//生成右侧拦 图片缩略
			template= Handlebars.compile(newsThumbTemp);
			content = template(data);
			$("#_newsThumb").append(content);
			showNewsForm("add",id);//显示当前创建的form
		}
	}
	$("#addNews").click(function(){
		createNewsForm();
	});
	createNewsForm();//默认创建一个
});