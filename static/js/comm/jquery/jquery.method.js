//添加的公共提示框和层
$(function(){
	$.extend({
		alertmsg:function(ele,type,msg){
			$.bootstrapGrowl(msg, {
				ele: ele, // which element to append to
				type: type, // (null, 'info', 'danger', 'success')
				offset: {from: 'top', amount: 20}, // 'top', or 'bottom'
				align: 'center', // ('left', 'right', or 'center')
				width: 'auto', // (integer, or 'auto')
				delay: 3000, // Time while the message will be displayed. It's not equivalent to the *demo* timeOut!
				allow_dismiss: true, // If true then will display a cross to close the popup.
				stackup_spacing: 10 // spacing between consecutively stacked growls.
			});
		},//type=1 4 5 6
		loadding:{
			New:function(msg,type,time,btn){
				var doc = $(window.parent.document);
				var iframeLen=doc.find("iframe").length;
				if(iframeLen==0){
					doc=$(window.document)
				}
				var back='<div id="_back_mask" style="position: absolute; top: 0px; filter: alpha(opacity=60); background-color: #777;z-index: 9999; left: 0px;opacity:0.5;-moz-opacity:0.5;display:none;"></div>';
				back = $(back);
				back.css("height",doc.height());   
        		back.css("width",doc.width());
				if(doc.find("#_back_mask").length==0){
					doc.find("body").append(back);
				}
				console.log(iframeLen,doc.height(),doc.width(),$(document).height(),$(document).width());
				var loadding = {
					doc:doc,
					btn:btn,
					msg:msg,
					type:type,
					time:time,
					btnType:false,
					Show:function(){
						if(!loadding.time){
							loadding.time=-1;
						}
						if(loadding.btn&&loadding.btn!=""){
							$(loadding.btn).addClass("disabled");
							var btn=$(loadding.btn).attr("type");
							if(btn=="submit"){
								loadding.btnType=true;
								$(loadding.btn).attr("type","button");
							}
						}
						ZENG.msgbox.show(loadding.msg, loadding.type, loadding.time);
						loadding.doc.find("#_back_mask").show();
					},
					Hide:function(){
						ZENG.msgbox._hide();
						if(loadding.btn&&loadding.btn!=""){
							$(loadding.btn).removeClass("disabled");
							if(loadding.btnType){
								loadding.btnType=false;
								$(loadding.btn).attr("type","submit");
							}
						}
						loadding.doc.find("#_back_mask").hide();
					}
				}
				return loadding;
			}
		}
	});
	$.validator.setDefaults({
		ignore : "",
		invalidHandler: function(form, validator) {
			$.each(validator.invalid,function(key,value){
				$.alertmsg("#tipsMsg","danger",value);
				return false;
			}); 
		},
		errorPlacement: function(error, element) {
		}
	});
})