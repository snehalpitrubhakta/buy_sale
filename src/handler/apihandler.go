package handler

import (
	"data"
	"database/sql"
	"db"
	"encoding/json"
	"logger"

	_ "github.com/go-sql-driver/mysql"
)

const (
	UPDATE    = "update"
	GETDETAIL = "getdetail"
	ADDCART   = "addcart"
	ORDER     = "order"

	//add handler for update and get user details.login for authentication and reg for register.
)

func HandlerApiRequest(conn *sql.DB, request *data.ApiRequest) []byte {
	var response *data.ApiResponse = nil
	cartDetailObj := &db.Cart_details{}
	err := json.Unmarshal([]byte(*request.Data), cartDetailObj)
	if err != nil {
		panic(err)
	}
	logger.Info(" API : %s", *request.Api)
	switch *request.Api {
	case ORDER:
		logger.Info("In Place order info")
		isSuccess := cartDetailObj.InsertIntoSoldItems(conn)
		if isSuccess == 0 {
			logger.Fatal("Order not placed.")
			response = &data.ApiResponse{Status: 200, Data: " Order not Placed", Error: ""}
		} else {
			logger.Info("order successfully placed.")
			response = &data.ApiResponse{Status: 200, Data: "Order Successfully Placed", Error: ""}
		}
	case ADDCART:
		logger.Info("in Add cart:apihandler.")
		isSuccess := cartDetailObj.SaveToCart(conn)
		if isSuccess == 0 {
			logger.Fatal("Data in not added in cart.")
			response = &data.ApiResponse{Status: 200, Data: " Data not added to cart.", Error: ""}
		} else {
			logger.Info("data added to cart.")
			response = &data.ApiResponse{Status: 200, Data: " Data added to cart.", Error: ""}
		}
	default:
		logger.Warn("Default case.")
		response = &data.ApiResponse{Status: 200, Data: nil, Error: ""}
	}
	resp, err := json.Marshal(response)
	if nil != err {
		return nil
	}
	return resp
}
