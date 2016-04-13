function enable(id){
	var url="/pubnum/enable";
	$.get(url,{pubnumid:id},function(result){
		console.log(result);
		if(result.code=="0"){
			window.location.reload();
		}
	});
}