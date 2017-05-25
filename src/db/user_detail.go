package db

import (
	"database/sql"
	"logger"
	"strconv"
	"strings"
)

type User_detail struct {
	UserId    int64 `json:",string"`
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func (detail *User_detail) SaveUserDetail(conn *sql.DB) int64 {
	id := detail.UserId
	fname := detail.FirstName
	lname := detail.LastName
	email := detail.Email
	password := detail.Password
	sqlStatement := strings.Join([]string{"INSERT INTO userinfo (userid,firstname,lastname,email,password) VALUES(", strconv.FormatInt(id, 10), ",'", fname, "','", lname, "','", email, "','", password, "');"}, "")
	_, err := conn.Exec(sqlStatement)
	if err != nil {
		logger.Fatal("%s", err)
		return 0
	}
	return 1
}

func Authenticate(conn *sql.DB, email string, password string) bool {
	query := strings.Join([]string{"SELECT count(userid) as count FROM userinfo WHERE email = '", email, "' AND password = '", password, "';"}, "")
	logger.Info("%s", query)
	rows, err := conn.Query(query)
	if err != nil {
		logger.Debug("%s", err)
	}
	defer rows.Close()
	for rows.Next() {
		var email string
		err := rows.Scan(&email)
		if err != nil {
			logger.Fatal("%s", err)
			return false
		}
		count, err := strconv.Atoi(email)
		if count == 0 {
			return false
		}
	}
	return true
}

func GetUSerDetailsById(conn *sql.DB, userId int64) *User_detail {
	details := &User_detail{}
	QueryForData := strings.Join([]string{"SELECT userid,firstname,lastname,email,password  FROM userinfo WHERE userid = '", strconv.FormatInt(userId, 10), "';"}, "")
	data, err := conn.Query(QueryForData)
	if err != nil {
		logger.Debug("%s", err)
	}
	defer data.Close()
	for data.Next() {
		err = data.Scan(&details.UserId, &details.FirstName, &details.LastName, &details.Email, &details.Password)
		if err != nil {
			logger.Error("%s", err)
		}
	}
	return details
}

func (detail *User_detail) UpdateUserDetail(conn *sql.DB) bool {
	id := detail.UserId
	fname := detail.FirstName
	lname := detail.LastName
	email := detail.Email
	password := detail.Password
	sqlStatement := strings.Join([]string{"UPDATE userinfo SET firstname='", fname, "',lastname='", lname, "',email='", email, "',password='", password, "'  WHERE userid = '", strconv.FormatInt(id, 10), "';"}, "")
	_, err := conn.Exec(sqlStatement)
	if err != nil {
		logger.Fatal("%s", err)
		return false
	}
	return true
}
