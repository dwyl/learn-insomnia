package main

import (
	"encoding/json"
	"log"
	"net/http"

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
	router.GET("/", Index)
	router.POST("/", New)

	log.Fatal(http.ListenAndServe(":8080", router))
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
		log.Print("Coudl not decode msg")
	}

	t.ID = idCounter
	idCounter++
	list = append(list, t)

	json.NewEncoder(w).Encode(t)
}
