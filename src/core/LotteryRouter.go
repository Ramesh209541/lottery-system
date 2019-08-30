package main

import (
	"errors"
	log "logger"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func createRouter() (*chi.Mux, error) {
	r := chi.NewRouter()
	if r == nil {
		log.Error.Println("error while creating new router is nil")
		return nil, errors.New("error while creating new router")
	}
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Route("/ticket", func(r chi.Router) {
		r.Post("/", ticketCreator)
		r.Get("/", listofTickets)
		r.Get("/{id}", getTicketData)
		r.Put("/{id}", amendDatatoTicket)
	})
	r.Route("/status", func(r chi.Router) {
		r.Put("/{id}", statusCheckout)
	})
	return r, nil

}
