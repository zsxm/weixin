$(function(){
	var pubnum_temp='<option value="-1">请选择</option>{{#each .}}<option value="{{id}}">{{name}}</option>{{/each}}';
	var url="/pubnum/get/id";
	$.get(url,{},function(result){
		var template = Handlebars.compile(pubnum_temp);
		var content = template(result);
		$("#pubnumid").html(content);
	});
	
	$("#get").click(function(){
		var url="/token/api/get/token";
		var pubnumid=$("#pubnumid").val();
		$.get(url,{pubnumid:pubnumid},function(result){
			if(result.code=="0"){
				$("#token").val(result.token);
			}else{
				alert(result.codemsg);
			}
		});
	});
})