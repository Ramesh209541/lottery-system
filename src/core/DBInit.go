package main

import (
	"database/sql"
	log "logger"
)

func initializeDatabase(DBfile string) error {
	var err error
	database, err = sql.Open("sqlite3", dataBaseFile)
	if err != nil {
		log.Error.Println("error opening sql data base is:", err)
		return err
	}
	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS ticket (id INTEGER NOT NULL PRIMARY KEY, status BOOL)")
	if err != nil {
		log.Error.Println("error preparing the ticket table err:", err)
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		log.Error.Println("executing of data base is failed err:", err)
		return err
	}
	statement, err = database.Prepare("CREATE TABLE IF NOT EXISTS ORDERS ( ID2 int,value int,outcome int ,  FOREIGN KEY (ID2) REFERENCES ticket (id));")
	if err != nil {
		log.Error.Println("error preparing the ticket line table:", err)
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		log.Error.Println("executing of data base is failed err:", err)
		return err
	}

	statement, err = database.Prepare("CREATE TABLE IF NOT EXISTS ticketcount (count int)")
	if err != nil {
		log.Error.Println("error preparing the table of ticket count:", err)
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		log.Error.Println("executing of data base is failed err:", err)
		return err
	}

	ticketcountrows, err := database.Query("select count from ticketcount")
	if err != nil {
		log.Error.Println("error querrying", err)
		return err
	}
	for ticketcountrows.Next() {
		err = ticketcountrows.Scan(&ticketCount)
		if err != nil {
			log.Error.Println("scan of ticketcount failed", err)
			return err
		}
	}
	if ticketCount == 1 {
		statement, err := database.Prepare("INSERT INTO ticketcount (count) VALUES (?)")
		if err != nil {
			log.Error.Println("error preparing data base  is:", err)
			return err
		}
		_, err = statement.Exec(ticketCount)
		if err != nil {
			log.Error.Println("executing of data base is failed err:", err)
			return err
		}
	}
	return nil
}
