package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type allEvents []event

var events = allEvents{
	{
		ID:          "1",
		Title:       "Introduction to Golang",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
	},
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	err = json.Unmarshal(reqBody, &newEvent)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Error de Unmarshall")
	}

	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(newEvent)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Error de NewEncoder")
	}
}
func homeLink(w http.ResponseWriter, _ *http.Request) {

	_, _ = fmt.Fprintf(w, "This new Site comming soon!")
}
func getAllEvents(w http.ResponseWriter, _ *http.Request) {
	var _ = json.NewEncoder(w).Encode(events)
}
func main() {
	port := os.Getenv("PORT")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/kitchen", homeLink)
	router.HandleFunc("/event", createEvent).Methods("POST")
	router.HandleFunc("/event2", createEvent).Methods("POST")
	router.HandleFunc("/event3", createEvent).Methods("POST")
	router.HandleFunc("/events", getAllEvents).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
