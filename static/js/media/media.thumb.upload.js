$(function(){
	var url="/media/upload";
	$("#thumbFile").fileinput({
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
		allowedFileExtensions:["jpg","jpeg", "png", "gif","bmp"],
        initialCaption: "请选择上传素材",
		dropZoneTitle:"缩略图文件拖拽到这里...",
		browseLabel: '选择文件',//选择按钮文字 
		//removeLabel: '删除所有',//移除按钮文字
		showRemove:false//隐藏移动按钮
	}).on("fileuploaded",function(event,data,prvid,index){
		var result=data.response;
		if (result.Code=="0") {
			var da=result.Data;
			var dir=da.DirName;
			var fn=da.FileNameId[0];
			var filePath=dir+"/"+fn;
			$("#thumbFilePath").val(filePath);				
		}else{
			alert(result.Codemsg);	
		}
	}).on("fileremoved",function(event,prvid,index){//未上传时点击删除
		console.log("fileremoved",prvid);
		$("#thumbFilePath").val("");
	}).on("filesuccessremove",function(event,prvid,index){//上传成功后点击删除
		console.log("filesuccessremove",prvid);
		$("#thumbFilePath").val("");
	});
	
	//保存
	$("#thumbReleaseBtn").click(function(){
		var thumbFilePath=$("#thumbFilePath").val();
		if(thumbFilePath==""){
			alert("请上传图片");
			return;
		}
		$("#thumbForm").ajaxSubmit({
			dataType : "json",
			success : function(result){
				console.log(result);
				if(result.Code=="0"){
					alert("素材保存成功");
				}else{
					alert(result.Codemsg);
				}
			}
		});
	});
})