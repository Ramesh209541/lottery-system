package main

import (
	"encoding/json"
	"fmt"
	log "logger"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type amendType struct {
	LinesCount int
}

func amendDatatoTicket(w http.ResponseWriter, r *http.Request) {
	ticketID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(ticketID)
	if err != nil {
		fmt.Fprintf(w, "\ninvalid index")
		log.Error.Println("error while converting string to int")
		return
	}
	decoder := json.NewDecoder(r.Body)
	var t amendType
	err = decoder.Decode(&t)
	if err != nil {
		fmt.Fprintf(w, "\ninvalid json data")
		log.Error.Println("err in decoding json data is:", err)
		return
	}
	err = amendLinetoticketInDB(id, t.LinesCount, w)
	if err != nil {
		fmt.Fprintf(w, "\namending tickets is failed")
		log.Error.Println("amendLinetoticketInDB is failed")
		return
	}
}

func statusCheckout(w http.ResponseWriter, r *http.Request) {
	ticketID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(ticketID)
	if err != nil {
		fmt.Fprintf(w, "\ninvalid ticket id \n")
		log.Error.Println("error while converting string to int")
		return
	}

	err = statusCheckoutinDB(id, w)
	if err != nil {
		fmt.Fprintf(w, "/nchecking status of ticket is failed")
		log.Error.Println("statusCheckoutinDB is failed")
		return
	}

}
