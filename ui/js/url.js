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
			requestJSON = JSON.stringify(requestObj);
			var url = "http://localhost:81/apiserver?query="+requestJSON;
			PDP.callWithAjax("GET",url,"json",null,"Please wait...!!","Internal Error","application/json",window.renderResponse,self);
			})
			window.renderResponse = function(response, obj){
					console.log(response);
			};
	})
	
	