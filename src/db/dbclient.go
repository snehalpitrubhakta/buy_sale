package db

import (
	"database/sql"
	"logger"

	_ "github.com/go-sql-driver/mysql"
)

type DbConfig struct {
	Host     *string
	Port     *int
	User     *string
	Password *string
	DbName   *string
}

func (config *DbConfig) GetConnection(dbName string) *sql.DB {
	//hostName := *config.Host
	//port := *config.Port
	//userName := *config.User
	//password := *config.Password
	//dbInfo := strings.Join([]string{"host='", hostName, "' port=", strconv.Itoa(port), " user='", userName, "' password='", password, "' dbname='", dbName, "' sslmode=disable"}, "")
	conn, err := sql.Open("mysql", "projectpdp:pdp@tcp(127.0.0.1:3306)/"+dbName)
	if err != nil {
		logger.Error(err.Error())
		conn = nil
	}
	return conn
}
