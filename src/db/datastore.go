package db

import (
	"database/sql"
	"log"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Cart_details struct {
	UserName string
	ItemName string
	Quantity int64 `json:",string"`
	Quality  string
	Address  string
}

func (detail *Cart_details) SaveToCart(conn *sql.DB) int64 {
	username := detail.UserName
	itemname := detail.ItemName
	quantity := detail.Quantity
	quality := detail.Quality
	price := quantity * 30
	var retVal int64 = 1
	sqlStatement := strings.Join([]string{"INSERT INTO cart (username,itemname,quantity,quality,price) VALUES ('", username, "','", itemname, "',", strconv.FormatInt(quantity, 10), ",'", quality, "',", strconv.FormatInt(price, 10), ");"}, "")
	_, err := conn.Exec(sqlStatement)
	if err != nil {
		log.Fatal("%s", err)
		retVal = 0
	}
	return retVal
}

func (detail *Cart_details) InsertIntoSoldItems(conn *sql.DB) int64 {
	username := detail.UserName
	itemname := detail.ItemName
	quantity := detail.Quantity
	quality := detail.Quality
	address := detail.Address
	var retVal int64 = 1
	sqlStatement := strings.Join([]string{"INSERT INTO solditems (username,itemname,quantity,quality,address) VALUES ('", username, "','", itemname, "',", strconv.FormatInt(quantity, 10), ",'", quality, "','", address, "');"}, "")
	_, err := conn.Exec(sqlStatement)
	if err != nil {
		log.Fatal("%s", err)
		retVal = 0
	}
	return retVal
}

/*
func InertIntoPaymentDetails(conn *sql.DB, username string, itemid string, paymentType string) {
	stmt, err := conn.Prepare("insert into paymentdetails(username,itemid,paymenttype) values(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(username, itemid, paymentType)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}

/*
func InsertUserData(conn *sql.DB, username string, email string, password string, phonenumber string, address string) {
	stmt, err := conn.Prepare("insert into userdetails(username,email,password,phonenumber,address) values(?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(username, email, password, phonenumber, address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
*/
