package router

import (
	"log"
	"net/http"

	"github.com/Shashwat5522/golang_gloabalsearch_mongo/controllers"
	"github.com/gorilla/mux"
)

func Run() {
	router := mux.NewRouter()
	router.HandleFunc("/objects", controllers.CreateObjects).Methods("POST")
	router.HandleFunc("/objects", controllers.SearchObjects).Methods("GET")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/",http.FileServer(http.Dir("./static/"))))
	router.HandleFunc("/index",controllers.ShowIndexPage).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
