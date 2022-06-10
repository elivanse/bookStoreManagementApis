package main

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm/dialects/mysql"
	"home/ivanse/dev/bookStoreManagementApis/pkg/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))

}
