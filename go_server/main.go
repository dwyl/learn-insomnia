package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Todo for our list
type Todo struct {
	ID   int    `json:"id"`
	Item string `json:"item"`
}

// Todos is a Todo list
type Todos []Todo

var idCounter = 3

var list = Todos{
	Todo{
		ID:   1,
		Item: "Star dwyl on GitHub",
	},
	Todo{
		ID:   2,
		Item: "Contribute to open source",
	},
}

func main() {
	router := httprouter.New()
	router.GET("/todo/:id", View)
	router.GET("/admin", Admin)
	router.GET("/", Index)
	router.POST("/", New)

	log.Fatal(http.ListenAndServe(":4000", router))
}

// Index lists our shop items
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	json.NewEncoder(w).Encode(list)
}

// New creates a new item
func New(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var t Todo
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		log.Print("Could not decode msg")
	}

	t.ID = idCounter
	idCounter++
	list = append(list, t)

	json.NewEncoder(w).Encode(t)
}

// View gets one item
func View(w http.ResponseWriter, r *http.Request, hp httprouter.Params) {
	var todo *Todo
	id, _ := strconv.Atoi(hp.ByName("id"))

	for _, item := range list {
		if item.ID == id {
			todo = &item
			break
		}
	}

	if todo == nil {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Item not found")
	} else {
		json.NewEncoder(w).Encode(*todo)
	}
}

// Admin allows us to test protecting routes
func Admin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Authorization") == "Bearer hunter2" {
		fmt.Fprintf(w, "Successfully Authenticated!")
	} else {
		fmt.Fprint(w, "Authorization failed.")
	}
}
