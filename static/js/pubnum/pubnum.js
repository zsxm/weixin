function enable(id){
	var url="/pubnum/enable";
	$.get(url,{pubnumid:id},function(result){
		console.log(result);
		if(result.code=="0"){
			window.location.reload();
		}
	});
}
$(function(){
	
	$("#pubnumForm").validate({
		rules:{
			name:{
				required:true
			},
			appid:{
				required:true
			},
			appsecret:{
				required:true
			},
			token:{
				required:true
			}
		},
		messages:{
			name:{
				required:"请填写姓名"
			},
			appid:{
				required:"请填写appid"
			},
			appsecret:{
				required:"请填写appsecret"
			},
			token:{
				required:"请填写token"
			}
		},
		submitHandler:function(form){
			var load=$.loadding.New("正在保存请稍候...", 6, -1, "#saveReleaseBtn");
			load.Show();
			$("#pubnumForm").ajaxSubmit({
				dataType : "json",
				success : function(result){
					load.Hide();
					if(result.code=="0"){
						$.alertmsg("#tipsMsg","success","公众号添加成功");
					}else{
						$.alertmsg("#tipsMsg","danger",result.codemsg);
					}
				}
			});
		}
	});
})