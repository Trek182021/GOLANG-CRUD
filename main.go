package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Event struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Title  string `json:"title"`
	Host   *Host  `json:"host"` // "*" is a pointer, pointing to Director struct
}

type Host struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func getEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

func getEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range events {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range events {
		if item.ID == params["id"] {
			events = append(events[:index], events[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(events)
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var event Event
	_ = json.NewDecoder(r.Body).Decode(&event)
	event.ID = strconv.Itoa(rand.Intn(100000000))
	events = append(events, event)
	fmt.Println(event)
	json.NewEncoder(w).Encode(event)
}

func updateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range events {
		if item.ID == params["id"] {
			events = append(events[:index], events[index+1:]...)
			var event Event
			_ = json.NewDecoder(r.Body).Decode(&event)
			event.ID = params["id"]
			events = append(events, event)
			json.NewEncoder(w).Encode(event)
			break
		}
	}
}

var events []Event

func main() {
	// fmt.Printf("Hello")
	r := mux.NewRouter()

	events = append(events, Event{ID: "191231", Status: "Active", Title: "Event One", Host: &Host{Firstname: "John", Lastname: "Doe"}})
	events = append(events, Event{ID: "91231", Status: "Active", Title: "Event Two", Host: &Host{Firstname: "John", Lastname: "Doe2"}})
	r.HandleFunc("/events", getEvents).Methods("GET")
	r.HandleFunc("/events/{id}", getEvent).Methods("GET")
	r.HandleFunc("/events", createEvent).Methods("POST")
	r.HandleFunc("/events/{id}", updateEvent).Methods("PUT")
	r.HandleFunc("/events/{id}", deleteEvent).Methods("DELETE")

	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
