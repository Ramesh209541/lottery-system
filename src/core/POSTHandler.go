package main

import (
	"encoding/json"
	"fmt"
	log "logger"
	"net/http"
)

type ticketData struct {
	Quantity       int
	LinesPerTicket int
}

func ticketCreator(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t ticketData
	err := decoder.Decode(&t)
	if err != nil {
		fmt.Fprintf(w, "invalid json data \n")
		log.Error.Println("Error while decoding json data")
		return
	}
	if t.Quantity <= 0 || t.LinesPerTicket <= 0 {
		fmt.Fprintf(w, "invalid json data \n")
		log.Error.Println("quantity or lines per ticket should not be zero")
		return
	}
	err = insertDatatoDB(t.Quantity, t.LinesPerTicket)
	if err != nil {
		log.Error.Println("insertDatatoDB is failed")
		fmt.Fprintf(w, "creating ticket is failed")
		return
	}
	fmt.Fprintf(w, "creating new ticket is successful")
}
