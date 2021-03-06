//模版点击后添加到编辑器中
var tempClick=function(t){
	var o=$(t);
	var id=$("#activityForm").val();
	var content="content_"+id;
	UE.getEditor(content).execCommand('insertHtml', o.html())
}
//模版类型按钮点击
var tempBtnClick=function(t,id){
	var o=$(t);
	var pid=o.attr("id");
	var url="/media/news/add/temp";
	var temp=$("#newsFormTemplate_"+id);
	var tempList=temp.find("#tempList");
	tempList.hide();
	$.get(url,{pid:pid},function(result){
		var template = Handlebars.compile('{{#each .}}<div id="temp_'+id+'_{{id}}" onclick="tempClick(this)" style="cursor:pointer;border:1px solid gray;margin: 3px 3px 3px 3px ;">{{{content}}}</div>{{/each}}');
		var content = template(result);
		tempList.html(content);
		tempList.fadeIn(300);//显示
	});
}

//显示活动的form
var showNewsForm=function(t,id){
	$(".form-horizontal").hide();//隐藏
	$("#newsForm_"+id).fadeIn(300);//显示
	if(t!="add"){
		$("html,body").animate({scrollTop:$("#tipsMsg").offset().top},500);
	}
	$("#activityForm").val(id);//设置当前活动的表单id
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
//获得图文数量
var newsFormLen=function(){
	return $(".form-horizontal").length;
}
//预览
var previewEditor=function(){
	var id=$("#activityForm").val();
	var content='content_'+id;
	var content=UE.getEditor(content).getContent()
	console.log(content);
}

//清空
var clearEditor=function(){
	var id=$("#activityForm").val();
	var content='content_'+id;
	UE.getEditor(content).setContent('', false);
}

$(function(){
	$.bootstrapGrowl('<button id="previewEditorBtn" onclick="previewEditor()" class="btn btn-default">预览</button>', {
		ele: 'body',
		type: 'info',
		offset: {from: 'bottom', amount: 220},
		align: 'right',
		width: 85,
		delay: -1,
		allow_dismiss: false,
		stackup_spacing: 10
	},$(window.document));
	$.bootstrapGrowl('<button id="clearEditorBtn" onclick="clearEditor()" class="btn btn-default">清空</button>', {
		ele: 'body',
		type: 'success',
		offset: {from: 'bottom', amount: 160},
		align: 'right',
		width: 85,
		delay: -1,
		allow_dismiss: false,
		stackup_spacing: 10
	},$(window.document));
	
	var toolbars={
		toolbars: [
			['fullscreen', 'source', 'undo', 'redo','bold', 'italic', 'underline', 'fontborder', 'strikethrough', 'superscript', 'subscript', '|', 
			'removeformat', 'formatmatch', 'autotypeset', 'blockquote', 'pasteplain', '|', 
			'simpleupload', 'insertimage', 'emotion', 'scrawl', 'insertvideo', 'music', 'attachment', 'map', 'gmap', 'webapp', 'pagebreak', 'template', 'background', '|',
			'forecolor', 'backcolor', 'insertorderedlist', 'insertunorderedlist', 'selectall', 'cleardoc']//,'preview'
		]
	}
		
	var newsFormEditor=function(id){
		var content='content_'+id;
    	var ue = UE.getEditor(content,toolbars);
	}

	//上传控件
	var newsFormFileInput=function(id){
		var load;
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
			showPreview:false,//关闭预览
			removeLabel: "删除",//移除按钮文字
			showRemove:true//隐藏移除按钮
		}).on("fileuploaded",function(event,data,prvid,index){
			load.Hide()
			console.log(2);
			var result=data.response;
			//console.log(id,result);
			if (result.code=="0") {
				var da=result.data;
				var dir=da.DirName;
				var fn=da.FileNameId[0];
				var filePath=dir+"/"+fn;
				load=$.loadding.New("正在保存到微信服务器...", 6, -1);
				load.Show();
				//本地上传成功后，再上传到微信服务器
				var url="/media/release/common";
				$.post(url,{localName:filePath,ctype:"image"},function(result){
					load.Hide()
					console.log(result)
					if (result.code=="0") {
						$("#thumb_media_id_"+id).val(result.data);//封面图片mediaId
						$("#newsFilePath_"+id).val(filePath);
						$("#newsThumb_"+id).attr("src","/"+filePath);	
						$.alertmsg("#tipsMsg","success","封面保存成功");
					}else{
						$.alertmsg("#tipsMsg","danger",result.codemsg);
					}
					console.log(result);
				});
			}else{
				$.alertmsg("#tipsMsg","danger",result.codemsg);
			}
		}).on("filepreupload",function(event,prvid,index){//上传之前
			load=$.loadding.New("正在上传...", 6, -1, "#newsReleaseBtn");
			load.Show();
		}).on("fileclear",function(event,prvid,index){//移除按钮点击事件
			$("#newsFilePath_"+id).val("");
			$("#thumb_media_id_"+id).val("");
			$("#newsThumb_"+id).attr("src","/static/image/comm/golang.png");
		}).on("fileremoved",function(event,prvid,index){//未上传时点击删除
			//console.log("fileremoved",prvid);
			$("#newsFilePath_"+id).val("");
			$("#thumb_media_id_"+id).val("");
			$("#newsThumb_"+id).attr("src","/static/image/comm/golang.png");	
		}).on("filesuccessremove",function(event,prvid,index){//上传成功后点击删除
			//console.log("removed",prvid);
			$("#newsFilePath_"+id).val("");
			$("#thumb_media_id_"+id).val("");
			$("#newsThumb_"+id).attr("src","/static/image/comm/golang.png");	
		});
	}
	//验证所有相同的title
	$.validator.addMethod('myTitles', function(value,element,param){
		var bool=true;
		var titles=$("input[name='title']");
		for(var i=0;i<titles.length;i++){
			var title=$(titles[i]);
			var val=$.trim(title.val());
			if(val==""||val.length==64){
				var id=title.attr("id");
				id=id.split("_")[1];
				var form=$("#newsForm_"+id);
				var dp=form.css("display");
				if(dp=="none"){
					showNewsForm('edit',id);
				}
				bool=false;
				break;
			}
		}
		return bool;
	}, '');
	
	$("#newsMediaForm").validate({
		debug:false,
		onfocusout:false,
		onkeyup:false,
		rules:{
			title:{
				myTitles:true
			}
		},
		messages: {
			title:{
				myTitles:"标题不能为空且长度不能超过64字"
			}
		},
		submitHandler:function(form){
			var bool=true;
			var contents=$("textarea[name='content']");
			for(var i=0;i<contents.length;i++){
				var content=$(contents[i]);
				var id=content.attr("identif");
				id=id.split("_")[1];
				var content='content_'+id;
				if(!UE.getEditor(content).hasContents()){
					var form=$("#newsForm_"+id);
					var dp=form.css("display");
					if(dp=="none"){
						showNewsForm('edit',id);
					}
					$.alertmsg("#tipsMsg","danger","正文不能为空且长度不能超过20000字");
					bool=false;
					break
				}
			}
			if(bool){
				var load=$.loadding.New("正在保存请稍候...", 6, -1, "#newsReleaseBtn");
				load.Show();
				$("#newsMediaForm").ajaxSubmit({
					dataType : "json",
					success : function(result){
						load.Hide();
						if(result.code=="0"){
							$.alertmsg("#tipsMsg","success","素材保存成功");
						}else{
							$.alertmsg("#tipsMsg","danger",result.codemsg);
						}
						$("html,body").animate({scrollTop:$("#tipsMsg").offset().top},500);
					}
				});
			}else{
				return false;
			}
		}
	});
	//fomr
	var newsFormTemp = '{{#each .}}'
					//+'<form id="newsForm_{{id}}" action="{{action}}" method="post" class="form-horizontal" role="form">'
				+'<div id="newsForm_{{id}}" class="form-horizontal">'
					+'<div class="form-group">'
				    +'  	<div class="col-sm-4">'
				    +'     		<input type="text" class="form-control" id="newsTitle_{{id}}" name="title" maxlength="64" placeholder="标题" value="测试标题数据">'
				    +'  	</div>'
				    +'  	<div class="col-sm-4">'
				    +'    		<input type="text" class="form-control" id="newsAuthor_{{id}}" name="author" maxlength="20" placeholder="作者" value="测试作者数据">'
				    +'  	</div>'
				    +'  	<div>'
					+'			<input id="show_cover_pic_{{id}}" name="show_cover_pic" type="checkbox"/>'
				    +'     		<lable>是否显示封面</lable>'
				    +'  	</div>'
				   	+'</div>'
					+'<div class="form-group">'
				    +'  	<div class="col-sm-12">'
				    +'     		<input type="text" class="form-control" id="contentSourceUrl_{{id}}" name="content_source_url" maxlength="255" placeholder="原文链接" value="http://www.baidu.com">'
				    +'  	</div>'
				   	+'</div>'
					+'<div class="form-group">'
				    +'  	<div class="col-sm-12">'
					+'			<div id="template_{{id}}" class="col-md-5 column">模版内容</div>'
					+'			<div class="col-md-7 column"><div class="row">'
					+'				<script id="editor_{{id}}" type="text/plain" style="width:1024px;height:500px;"></script>'
					+'				<textarea identif="content_{{id}}" id="content_{{id}}" name="content">测试模版内容数据</textarea>'
				    +'  		</div></div>'
				    +'  	</div>'
				   	+'</div>'
					+'<div class="form-group">'
				    +'  	<div class="col-sm-12">'
				    +'     		<input type="file" id="newsFile_{{id}}" name="file" value="" accept="image/*" class="file-loading"/>'
					+'			<input type="hidden" name="ctype" value="news"/>'
					+'			<input type="hidden" id="newsFilePath_{{id}}" name="newsFilePath" value=""/>'
					+'			<input type="hidden" id="thumb_media_id_{{id}}" name="thumb_media_id" value=""/>'
					+'			<input type="hidden" name="saveType" value="{{saveType}}"/>'
				    +'  	</div>'
				   	+'</div>'
					+'<div class="form-group">'
				    +'  	<div class="col-sm-12">'
					+'			<button id="{{id}}" onclick="newsMediaSelect(this)" type="button" ctype="image" class="btn btn-block btn-default">封面 素材库选择</button>'
				    +'  	</div>'
				   	+'</div>'
			    	//+'</form>'
				+'</div>'
					+'{{/each}}';
	//右侧栏
	var newsThumbTemp='{{#each .}}<div id="_thumb_{{id}}" class="file-preview-frame" style="width:100px;height:100px;">'
				+'	<img id="newsThumb_{{id}}" src="/static/image/comm/golang.png" class="file-preview-image" '
				+'	style="width:100px;height:100px;">'
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
	//标题事件			
	var newsFormTitleKeyUp=function(id){
		$("#newsTitle_"+id).keydown(function(){
			var val=$(this).val();
			if(val.length>10){
				val=val.substr(0,10);
			}
			$("#titleMsg_"+id).html(val);
		});
		$("#newsTitle_"+id).keyup(function(){
			var val=$(this).val();
			if(val.length>10){
				val=val.substr(0,10);
			}
			$("#titleMsg_"+id).html(val);
		});
	}
	
	//模版加载
	var createNewsFormTemplate=function(id){
		var temp=$("#newsFormTemplate_"+id);
		if(temp.attr("temp")!="ok"){
			var url="/media/index?"+Math.random();
			$("#template_"+id).load(url,{page:"media.news.add.edit.temp"},function(d){
				temp=$("#newsFormTemplate_");
				temp.attr("temp","ok");
				temp.attr("id","newsFormTemplate_"+id);
				var url="/media/news/add/temp";
				var btnList=temp.find("#btnList");
				$.get(url,{pid:"root"},function(result){
					var template = Handlebars.compile('{{#each .}}<button id="{{id}}" onclick="tempBtnClick(this,\''+id+'\')" class="btn btn-default" type="button">{{name}}</button>{{/each}}');
					var content = template(result);
					btnList.html(content);
					btnList.find("button:eq(0)").click();//默认执行第一个
				});
			});
		}
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
			$('#show_cover_pic_'+id).checkboxpicker();
			newsFormFileInput(id);//生成inputfile
			newsFormEditor(id);//生成编辑内容
			newsFormTitleKeyUp(id);//title添加keyup事件
			//生成右侧拦 图片缩略
			template= Handlebars.compile(newsThumbTemp);
			content = template(data);
			$("#_newsThumb").append(content);
			showNewsForm("add",id);//显示当前创建的form
			createNewsFormTemplate(id);//模版加载
		}
	}
	$("#addNews").click(function(){
		createNewsForm();
	});
	createNewsForm();//默认创建一个
});