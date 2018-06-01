package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	log.Println("starting server...")
	db := &database{}
	http.Handle("/developers", &handler{db: db})
	http.HandleFunc("/healthz", healthz)
	http.ListenAndServe(":8080", nil)
}

type database struct{}

func (d *database) AllDevelopers() []developer {
	return []developer{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}
}

type handler struct {
	db *database
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json;charset=utf-8")
	ds := h.db.AllDevelopers()
	json.NewEncoder(w).Encode(developerList{Developers: ds})
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json;charset=utf-8")
	w.Write([]byte(`{"healthy": true}`))
}

type developer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type developerList struct {
	Developers []developer `json:"developers"`
}
