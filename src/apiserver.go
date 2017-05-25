package main

import (
	"config"
	"data"
	"database/sql"
	"encoding/json"
	"handler"
	"logger"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var conn *sql.DB

func handleRequest(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("query")
	apiRequest := &data.ApiRequest{}
	err := json.Unmarshal([]byte(query), apiRequest)
	if nil != err {
		w.Write([]byte(err.Error()))
		return
	}
	response := handler.HandlerApiRequest(conn, apiRequest)
	w.Write(response)

}

func main() {
	ConfigFile := "src/config/config.json"
	if !config.Init(ConfigFile) {
		panic("failed to initialise.")
	}
	loggerConfig := &logger.LoggerConfig{LogFile: *config.AppConfig.LogFile, LogLevel: logger.DEBUG}
	logger.Init(loggerConfig)
	conn = config.AppConfig.DbConfig.GetConnection("projectpdp")
	logger.Debug("connection created")
	logger.Info("API server started successfully.")
	http.HandleFunc("/apiserver", handleRequest)
	err := http.ListenAndServe(":3030", nil)
	logger.Debug("%s", err)
	logger.Debug("API server started successfully. with port:3030")
}
