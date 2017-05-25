$(document).ready(function(){
	$("#paymentOptions").hide();
	$("#itemSelection").show();

	$("#buynow").click(function(){
		$("#itemSelection").hide();
		$("#paymentOptions").show();
	});

	$("#addcart").click(function(){
		$("#cartmessage").text("Details added to cart.");
	});

	$("#addbtn").click(function(){
		$("addressMessage").text("details are saved.");
	});

	$("#next").click(function(){
		$("#itemSelection").hide();
		$("#paymentOptions").hide();
		$("#card").hide();
		$("#add").show();
	})
});