package main

import (
	"fmt"
	log "logger"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func listofTickets(w http.ResponseWriter, r *http.Request) {
	err := getListOfTicketsfromDB(w)
	if err != nil {
		log.Error.Printf("getListOfTicketsfromDB is failed\n")
		fmt.Fprint(w, "\ngetListOfTicketsfromDB is failed err:", err)
		return
	}
}

func getTicketData(w http.ResponseWriter, r *http.Request) {
	ticketID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(ticketID)
	if err != nil {
		fmt.Fprintf(w, "\ninvalid ticket id \n")
		log.Error.Println("error while converting string to int")
		return
	}
	err = getticketFromDB(id, w)
	if err != nil {
		fmt.Fprintf(w, "\ngetticketFromDB is failed to get the ticket \n")
		log.Error.Println("getticketFromDB is failed")
		return
	}
}
