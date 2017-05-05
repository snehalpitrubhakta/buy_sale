$(document).ready(function(){
	console.log("grainType");
	$("#paymentOptions").hide();
	$("#itemSelection").show();

	$("#buynow").click(function(){
		$("#itemSelection").hide();
		$("#paymentOptions").show();
	});
});