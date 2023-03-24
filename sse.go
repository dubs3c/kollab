package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Broker struct {

	// Events are pushed to this channel by the main events-gathering routine
	Notifier chan *EventPath

	// New client connections
	newClients chan chan EventPath

	// Closed client connections
	closingClients chan chan EventPath

	// Client connections registry
	clients map[chan EventPath]bool
}

func NewSSEServer() (broker *Broker) {
	// Instantiate a broker
	broker = &Broker{
		Notifier:       make(chan *EventPath, 1),
		newClients:     make(chan chan EventPath),
		closingClients: make(chan chan EventPath),
		clients:        make(map[chan EventPath]bool),
	}

	// Set it running - listening and broadcasting events
	go broker.listen()

	return
}

func (a *Api) SSEHandler(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "pathId")
	if id != "" {
		_, err := uuid.Parse(id)
		if err != nil {
			RespondWithError(w, 404, "ID not recognized")
			return
		}
	}

	// Make sure that the writer supports flushing.
	//
	flusher, ok := w.(http.Flusher)

	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Each connection registers its own message channel with the Broker's connections registry
	messageChan := make(chan EventPath)

	// Signal the broker that we have a new connection
	a.Broker.newClients <- messageChan

	// Remove this client from the map of connected clients
	// when this handler exits.
	defer func() {
		a.Broker.closingClients <- messageChan
	}()

	// Listen to connection close and un-register messageChan
	// notify := rw.(http.CloseNotifier).CloseNotify()
	notify := r.Context().Done()

	go func() {
		<-notify
		a.Broker.closingClients <- messageChan
	}()

	for {

		// Write to the ResponseWriter
		// Server Sent Events compatible
		fmt.Fprintf(w, "data: %s\n\n", <-messageChan)

		// Flush the data immediatly instead of buffering it for later.
		flusher.Flush()
	}

}

func (broker *Broker) listen() {
	for {
		select {
		case s := <-broker.newClients:

			// A new client has connected.
			// Register their message channel
			broker.clients[s] = true
			log.Printf("Client added. %d registered clients", len(broker.clients))
		case s := <-broker.closingClients:

			// A client has dettached and we want to
			// stop sending them messages.
			delete(broker.clients, s)
			log.Printf("Removed client. %d registered clients", len(broker.clients))
		case event := <-broker.Notifier:

			// We got a new event from the outside!
			// Send event to all connected clients
			for clientMessageChan := range broker.clients {
				clientMessageChan <- *event
			}
		}
	}

}
