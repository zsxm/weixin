$(function(){
	var STYLE_SETTING = 'style="width:{width};height:{height};"';
	var INPUT='<div><input type="text" id="tmpfiguretitle" name="tmpfiguretitle" class="form-control" filepath="noupload"/></div>';
	var url="/media/upload";
	$("#imageFile").fileinput({
		uploadUrl: url,//上传地址
		initialPreview:[],//默认加载文件
		initialPreviewConfig:[],//默认加载文件配置
		enctype: 'multipart/form-data',
		//minFileCount: 2,//至少有两个文件
		overwriteInitial: false,//重新选择文件后是否覆盖 initialPreview
		language:"zh",//设置语言
		//showUpload:false,//隐藏上传按钮
		maxFileCount: 10,//选择文件数量
		//maxImageWidth: 600,//文件宽度
       // maxFileSize: 256,//文件大小 2048kb 
		//elErrorContainer: "#errorBlock",
		previewFileType:["image"],//只选择image类型文件
		allowedFileExtensions:["jpg","jpeg", "png", "gif","bmp","mp3","wma","wav","amr"],
        initialCaption: "请选择上传素材",
		browseLabel: '选择文件',//选择按钮文字 
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