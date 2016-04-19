var comm={};
comm.closeTab=function(o,id){
	o=$(o);
	var tabPanel=$("#tabPanel_"+id);
	var panel=$("#panel_"+id);
	var cls=tabPanel.attr("class");
	if(cls=="active"){
		var tobj=tabPanel.prev().length==1?tabPanel.prev():tabPanel.next().length==1?tabPanel.next():"";
		if(tobj!=""){
			var tid=tobj.attr("id").split("_")[1];
			console.log(tid)
			$("#tab_"+tid).tab("show");
			$("#tab_"+tid).click();
		}
	}
	$("#"+id).removeAttr("id");
	tabPanel.remove();
	panel.remove();
}
$(function(){
	var tabTitle=$("#tabTitle");
	var tabContent=$("#tabContent");
	$(".panel-body a").click(function(){
		var o=$(this);
		var href=o.attr("href");
		var id="",bool=false;
		if(o.attr("id")==undefined){
			id=new Date().getTime()+""+parseInt(Math.random()*100);
			o.attr("id",id);
		}else{
			id=o.attr("id");
			bool=true
		}
		if(!bool){
			var title='<li id="tabPanel_'+id+'"><a id="tab_'+id+'" href="#panel_'+id+'" data-toggle="tab">'+o.attr("title")+' <b onclick="comm.closeTab(this,\''+id+'\')" style="color:blue;cursor:pointer;" title="关闭">X</b></a></li>';
			tabTitle.append(title);
			var content='<div class="tab-pane" id="panel_'+id+'">'
						+'	<iframe id="iframe_'+id+'" scrolling="auto" frameborder="0" src="'+href+'" style="width:100%;height:80%;"></iframe>'
						+'</div>';
			tabContent.append(content);
		}
		$("#tab_"+id).tab("show");
		return false;
	});
})