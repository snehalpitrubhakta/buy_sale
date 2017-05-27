$(document).ready(function(){
		var self = this;		
		event.preventDefault();
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
			var data = {
				"UserName":"shital",
				"ItemName":item,
				"Quantity":quantity,
				"Quality":quality,
				"Address":""
			}
			dataObj = JSON.stringify(data);			
			var requestObj = {
					"Api": "addcart",
					"Data": dataObj
			}
			requestJSON = JSON.stringify(requestObj);
			var url = "http://localhost:3030/apiserver?query="+requestJSON;
			console.log(url)
			PDP.callWithAjax("GET",url,"json",null,"Please wait...!!","Internal Error","application/json",window.renderResponse,self);
		});
		
		$("#buynow").on('click',function(){
			var itembuy = document.getElementById("items").value;
			var quantitybuy = document.getElementById("quantity").value;
			if(document.getElementById('high').checked){
				var qualitybuy = document.getElementById("high").value;
			}
			if(document.getElementById('medium').checked){
				var qualitybuy = document.getElementById("medium").value;
			}
			if(document.getElementById('low').checked){
				var qualitybuy = document.getElementById("low").value;
			}

			$("#placeOrder").on('click',function(){
			var address = document.getElementById("addr").value;
			var addressData = {
				"UserName":"shital",
				"ItemName":itembuy,
				"Quantity":quantitybuy,
				"Quality":qualitybuy,		
				"Address":address		
			}
			addObj = JSON.stringify(addressData);
			var requestObj = {
					"Api": "order",
					"Data": addObj
			}
			requestJSONbuy = JSON.stringify(requestObj);
			var urlbuy = "http://localhost:3030/apiserver?query="+requestJSONbuy;
			console.log(urlbuy);
			PDP.callWithAjax("GET",urlbuy,"json",null,"Please wait...!!","Internal Error","application/json",window.renderResponse,self);
		});		
		});

		

		window.renderResponse = function(response, obj){
			console.log(response);
		};
	});
	
	