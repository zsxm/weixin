$(function(){
	//同步永久素材数据
	$("#syncLineListBtn").click(function(){
		var load=$.loadding.New("正在同步数据,请稍候...", 6, -1, "#syncLineListBtn");
		load.Show();
		var url="/media/sync/line/list";
		$.get(url,{},function(result){
			load.Hide();
			if(result.code=="0"){
				$.alertmsg("#tipsMsg","success","数据同步成功");
			}else{
				$.alertmsg("#tipsMsg","danger",result.codemsg);
			}
		});
	});
	
	$("#getLocalListBtn").click(function(){
		var url="/media/temp/local/get";
		mediaList(url);
	});
	
	$("#getLineListBtn").click(function(){
		var url="/media/line/get";
		mediaList(url);
	});
	
	var mediaList=function(url){
		var mediaType=$("#mediaType").val();
		var local_temp='<thead><tr><th>MediaId</th><th>创建时间</th><th>类型</th><th>操作</th></tr></thead>';
		local_temp+='<tbody>'
			+'{{#each data}}'
			+'	<tr>'
			+'		<td><a href="#modal-container-detail" onclick="show(this)" title="{{title}}" id="{{mediaId}}" type="{{ctype}}" created="{{uyymdhms created}}" local="{{localName}}" url="{{url}}" data-toggle="modal">{{mediaId}}</a></td>'
			+'		<td>{{uyymdhms created}}</td><td>{{ctype}}</td>'
			+'		<td>{{#if (eq saveType 1)}}<a href="javascript:void(0);" onclick="deleteMedia(this,\'{{mediaId}}\')">删除</a>{{/if}}</td>'
			+'	</tr>'
			+'{{/each}}'
			+'</tbody>';
		$.get(url,{mediaType:mediaType},function(result){
			console.log(result);
			if(result.code=="0"){
				var template = Handlebars.compile(local_temp);
				var content = template(result);
				$("#mediaList").html(content);
			}else{
				$.alertmsg("#tipsMsg","danger",result.codemsg);
			}
		});
	}
})

var show=function(t){
	var o=$(t);
	$("#title").html(o.attr("title"));
	$("#mediaId").html(o.attr("id"));
	$("#created").html(o.attr("created"));
	var url=o.attr("url");
	if(url.indexOf("http")!=-1){
		url='<a href="'+url+'" target="_blank">查看</a>';
	}
	$("#url").html(url);
	$("#local").html(o.attr("local"));
	$("#type").html(o.attr("type"));
}

var deleteMedia=function(o,id){
	var url="/media/delete";
	$.confirm({
	    title: '删除提示!',
	    content: '删除该素材将无法恢复.是否确认删除？',
		confirmButton:"确认",
		cancelButton:"取消",
	    confirm: function(){
			var load=$.loadding.New("正在删除数据,请稍候...", 6, -1, "");
			load.Show();
	 		$.get(url,{id:id},function(result){
				load.Hide();
				if(result.code=="0"){
					$(o).parent().parent().remove();
					$.alertmsg("#tipsMsg","success","删除成功");
				}else{
					$.alertmsg("#tipsMsg","danger",result.codemsg);
				}
			});
	    },
	    cancel: function(){
			
	    }
	});

}