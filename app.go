package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Client struct {
	Id   int
	Name string
}

type App struct {
	Router  *mux.Router
	Clients []Client
}

func (app *App) Initialise() error {
	app.Clients = []Client{
		{Id: 1, Name: "Holding"},
		{Id: 2, Name: "Extrusion"},
		{Id: 3, Name: "Injection moulding"},
	}

	app.Router = mux.NewRouter().StrictSlash(true)
	app.handleRoutes()
	return nil
}

func (app *App) Run(address string) {
	log.Fatal(http.ListenAndServe(address, app.Router))
}

func (app *App) handleRoutes() {
	app.Router.HandleFunc("/clients", app.returnAllClients)
	app.Router.HandleFunc("/client/{id}", app.getClient)
	app.Router.HandleFunc("/", home)
}

func (app *App) returnAllClients(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint: returnAllClients")
	json.NewEncoder(w).Encode(app.Clients)
}

func (app *App) getClient(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	for _, client := range app.Clients {
		if client.Id == id {
			json.NewEncoder(w).Encode(client)
		}
	}

	w.WriteHeader(http.StatusNotFound)
}
