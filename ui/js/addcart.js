$(document).ready(function(){
	$("#addressOption").hide();
	$("#add").hide();
	$("#card").hide();	
	$("#addcart").click(function(){
		$("#itemSelection").hide();
		$("#addressOption").show();	
	});
	$("#back2").click(function(){
		$("#addressOption").hide();		
		$("#itemSelection").show();		
	});
	$("#cod").click(function(){
		$("#paymentOptions").hide();
		$("#add").show();
		console.log(this.id);
	});
	$("#back3").click(function(){
		$("#add").hide();		
		$("#paymentOptions").show();	
		console.log(this.parent.id);
		console.log("-----------------------");	
	});
	$("#back1").click(function(){
		$("#paymentOptions").hide();		
		$("#itemSelection").show();		
	});
	$("#credit").click(function(){
		$("#paymentOptions").hide();
		$("#card").show();
	});
	$("#back4").click(function(){
		$("#card").hide();		
		$("#paymentOptions").show();		
	});

});