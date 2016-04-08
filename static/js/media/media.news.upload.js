$(function(){
	//CKEDITOR.config.readOnly = true;//禁用或启用ckeditor
	//CKEDITOR.instances.content.setData("正文");//设置内容
	//CKEDITOR.instances.content.getData();//获取内容
	//CKEDITOR.instances.content.document.getBody().getText();//获得纯文本
	var editor1=CKEDITOR.replace('content',{
		toolbar:[
			['Font','Bold','Underline','Italic','TextColor','RemoveFormat','FontSize']
		]
	});
	var defaultText="正文";
	CKEDITOR.instances.content.on("instanceReady",function(){
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
	CKEDITOR.instances.content.on("instanceReady",function(){
		//this.document.getBody().setHtml('<img src="/static/image/login/1.jpg" border="0" title="Hello" />');
	});
	
	var STYLE_SETTING = 'style="width:{width};height:{height};"';
	var INPUT='<div><input type="text" id="tmpfiguretitle" name="tmpfiguretitle" class="form-control" filepath="noupload"/></div>';
	var url="/media/upload";
	$("#newsFile").fileinput({
		uploadUrl: url,//上传地址
		initialPreview:[],//默认加载文件
		initialPreviewConfig:[],//默认加载文件配置
		enctype: 'multipart/form-data',
		//minFileCount: 2,//至少有两个文件
		overwriteInitial: true,//重新选择文件后是否覆盖 initialPreview
		language:"zh",//设置语言
		//showUpload:false,//隐藏上传按钮
		maxFileCount: 1,//选择文件数量
		//maxImageWidth: 600,//文件宽度
       // maxFileSize: 256,//文件大小 2048kb 
		//elErrorContainer: "#errorBlock",
		previewFileType:["image"],//只选择image类型文件
		allowedFileExtensions:["jpg","jpeg", "png", "gif"],
        initialCaption: "请选择封面",
		dropZoneTitle:"封面拖拽到这里...",
		browseLabel: '选择封面',//选择按钮文字 
		//removeLabel: '删除所有',//移除按钮文字
		showRemove:false//隐藏移动按钮
		/*,previewTemplates:{//模版
				image:'<div class="file-preview-frame" id="{previewId}" data-fileindex="{fileindex}">'
				+'<img src="{data}" class="file-preview-image" title="{caption}" title="{caption}" ' + STYLE_SETTING + '>{footer}配图标题'+INPUT+'</div>'
		}*/
	}).on("fileuploaded",function(event,data,prvid,index){
		if (data.response.Code=="0") {
			alert("素材上传成功");						
		}else{
			alert(data.response.Codemsg);	
		}
	})/*.on("fileremoved",function(event,prvid,index){//未上传时点击删除
		//alert(event+"  "+prvid+"  "+index);
	})*/;
})