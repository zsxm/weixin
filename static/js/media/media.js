$(function(){
	$("#getLocalListBtn").click(function(){
		var url="/media/temp/local/get";
		var mediaType=$("#mediaType").val();
		if(mediaType=="-1"){
			alert("请选择素材类型");
			return;
		}
		var local_temp='<thead><tr><th widht="100px;">MediaId</th><th>创建时间</th><th>类型</th><th>路径</th><th>操作</th></tr></thead>';
		local_temp+='<tbody>'
			+'{{#each .}}'
			+'	<tr>'
			+'		<td>{{mediaId}}</td><td>{{uyymdhms created}}</td><td>{{ctype}}</td><td>{{localName}}</td><td>操作</td>'
			+'	</tr>'
			+'{{/each}}'
			+'</tbody>';
		$.get(url,{type:"local",mediaType:mediaType},function(result){
			var template = Handlebars.compile(local_temp);
			var content = template(result);
			$("#mediaList").html(content);
		});
	});
})