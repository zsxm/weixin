//被修改的
(function() {
	var c;
	c = jQuery;
	var docu = $(window.parent.document);
	var iframeLen=docu.find("iframe").length;
	if(iframeLen==0){
		docu=$(window.document)
	}
	c.bootstrapGrowl = function(f,a,doc) {
		if(doc==undefined){
			doc=docu;
		}
		var b, e, d;
		a = c.extend({},
		c.bootstrapGrowl.default_options, a);
		b = c("<div>");
		b.attr("class", "bootstrap-growl alert");
		a.type && b.addClass("alert-" + a.type);
		a.allow_dismiss && (b.addClass("alert-dismissible"), b.append('<button class="close" data-dismiss="alert" type="button"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>'));
		b.append(f);
		a.top_offset && (a.offset = {
			from: "top",
			amount: a.top_offset
		});
		d = a.offset.amount;
		doc.find(".bootstrap-growl").each(function() {
			return d = Math.max(d, parseInt(c(this).css(a.offset.from)) + c(this).outerHeight() + a.stackup_spacing)
		});
		e = {
			position: "body" === a.ele ? "fixed": "absolute",
			margin: 0,
			"z-index": "9999",
			display: "none"
		};
		e[a.offset.from] = d + "px";
		b.css(e);
		"auto" !== a.width && b.css("width", a.width + "px");
		doc.find(a.ele).append(b);
		switch (a.align) {
		case "center":
			b.css({
				left:
				"50%",
				"margin-left": "-" + b.outerWidth() / 2 + "px"
			});
			break;
		case "left":
			b.css("left", "20px");
			break;
		default:
			b.css("right", "20px")
		}
		b.fadeIn();
		0 < a.delay && b.delay(a.delay).fadeOut(function() {
			return c(this).alert("close")
		});
		return b
	};
	c.bootstrapGrowl.default_options = {
		ele: "body",
		type: "info",
		offset: {
			from: "top",
			amount: 20
		},
		align: "right",
		width: 250,
		delay: 4E3,
		allow_dismiss: !0,
		stackup_spacing: 10
	}
}).call(this);