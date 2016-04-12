$(function(){
	var saveType=$("#voiceForm #saveType").val();
	var url="/media/upload";
	
	$("#voiceFile").fileinput({
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
		previewFileType:["audio"],//只选择 audio 类型文件
		allowedFileExtensions:["mp3","wma","wav","amr"],
        initialCaption: "请选择上传素材",
		dropZoneTitle:"语音文件拖拽到这里...",
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
			$("#voiceFilePath").val(filePath);				
		}else{
			$.alertmsg("#tipsMsg","danger",result.Codemsg);
		}
	}).on("fileremoved",function(event,prvid,index){//未上传时点击删除
		console.log("fileremoved",prvid);
		$("#voiceFilePath").val("");
	}).on("filesuccessremove",function(event,prvid,index){//上传成功后点击删除
		console.log("filesuccessremove",prvid);
		$("#voiceFilePath").val("");
	});
	
	$("#voiceForm").validate({
		rules:{
			localName:{
				required:true
			}
		},
		messages:{
			localName:{
				required:"请上传语音文件"
			}
		},
		submitHandler:function(form){
			var load=$.loadding.New("正在保存请稍候...", 6, -1, "#videoReleaseBtn");
			load.Show();
			$("#voiceForm").ajaxSubmit({
				dataType : "json",
				success : function(result){
					load.Hide();
					if(result.Code=="0"){
						$.alertmsg("#tipsMsg","success","素材保存成功");
					}else{
						$.alertmsg("#tipsMsg","danger",result.Codemsg);
					}
				}
			});
		}
	});
})