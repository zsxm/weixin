function enable(id){
	var url="/pubnum/enable";
	$.get(url,{pubnumid:id},function(result){
		if(result.Code=="0"){
			window.location.reload();
		}
	});
}