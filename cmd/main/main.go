package main

import (
	"github.com/AkhilSharma90/Golang-MySQL-CRUD-Bookstore-Management-API/tree/master/pkg/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStore(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))

}
