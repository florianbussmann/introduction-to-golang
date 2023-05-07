package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
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
	key := strings.Split(r.URL.Path, "/")[2]

	id, err := strconv.Atoi(key)
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
	http.HandleFunc("/clients", returnAllClients)
	http.HandleFunc("/client/", getClient)
	http.HandleFunc("/", home)
	http.ListenAndServe("localhost:8000", nil)
}

func main() {
	Clients = []Client{
		{Id: 1, Name: "Holding"},
		{Id: 2, Name: "Extrusion"},
		{Id: 3, Name: "Injection moulding"},
	}
	handleRequests()
}
