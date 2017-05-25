$(document).ready(function(){
		var self = this;		
		event.preventDefault();
		var data = {};
		$("#addcart").on('click',function(){
			var item = document.getElementById("items").value;
			var quantity = document.getElementById("quantity").value;
			if(document.getElementById('high').checked){
				var quality = document.getElementById("high").value;
			}
			if(document.getElementById('medium').checked){
				var quality = document.getElementById("medium").value;
			}
			if(document.getElementById('low').checked){
				var quality = document.getElementById("low").value;
			}
			var requestObj = {
					"Api": "addcart",
					"ItemName":item,
					"Quantity":quantity,
					"Quality":quality
					}
			console.log(requestObj);
		
			})
	})
	
	