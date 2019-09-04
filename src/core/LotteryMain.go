package main

import (
	"database/sql"
	log "logger"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dataBaseFile = "./LotteryTickets.db"
	configFile   = "config/conf.json"
)

var database *sql.DB
var ticketCount = 1

type lotteryConfigType struct {
	PortNumber string `json:"ServerPortNumber"`
}

var lotteryConfig lotteryConfigType

func main() {
	var err error
	logHandle, err := OpenLoggerFile()
	if err != nil {
		os.Exit(1)
	}
	defer logHandle().Close()
	log.Info.Println("Lottery system started")
	err = parseConfigFileValues(configFile, &lotteryConfig)
	if err != nil {
		log.Error.Println("parsing configuration is failed")
		return
	}
	log.Info.Println("Lottery system started")

	err = initializeDatabase(dataBaseFile)
	if err != nil {
		log.Error.Println("initializeDatabase is failed")
		return
	}
	router, err := createRouter()
	if err != nil {
		log.Error.Println("Not able to create router")
		return
	}

	err = http.ListenAndServe(":"+lotteryConfig.PortNumber, router)
	if err != nil {
		log.Error.Println("starting server is failed err:", err)
		return
	}
}
