package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Client struct {
	Id   int
	Name string
}

var Clients []Client

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func returnAllClients(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint: returnAllClients")
	json.NewEncoder(w).Encode(Clients)
}

func getClient(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	for _, client := range Clients {
		if client.Id == id {
			json.NewEncoder(w).Encode(client)
		}
	}
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/clients", returnAllClients)
	router.HandleFunc("/client/{id}", getClient)
	router.HandleFunc("/", home)
	http.ListenAndServe("localhost:8000", router)
}

func main() {
	Clients = []Client{
		{Id: 1, Name: "Holding"},
		{Id: 2, Name: "Extrusion"},
		{Id: 3, Name: "Injection moulding"},
	}
	handleRequests()
}
