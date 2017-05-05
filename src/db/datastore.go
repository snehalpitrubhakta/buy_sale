package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type dbConfig struct {
	Host     *string
	Port     *int
	User     *string
	Password *string
	DbName   *string
}

func main() {
	dbuser := "projectpdp"
	password := "pdp"
	port := 3306
	dbName := "projectpdp"
	dbInfo := strings.Join([]string{dbuser, ":", password, "@tcp(127.0.0.1", ":", strconv.Itoa(port), ")/", dbName}, "")
	conn, err := sql.Open("mysql", dbInfo)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()
	//InsertUserData(conn, "shital", "shital@gmail.com", "2312", "8806686669", "jalgaon")
	//InsertIntoSoldItems(conn, 1, "i1", "shital")
	//InsertIntoCart(conn, "shital", "i2")
	//InertIntoPaymentDetails(conn, "shital", "i2", "cod")

}

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

func InsertIntoSoldItems(conn *sql.DB, orderid int, itemid string, username string) {
	stmt, err := conn.Prepare("insert into solditems(orderid,itemid,username) values(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(orderid, itemid, username)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}

func InsertIntoCart(conn *sql.DB, username string, itemid string) {
	stmt, err := conn.Prepare("insert into cart(username,itemid) values(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(username, itemid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}

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
