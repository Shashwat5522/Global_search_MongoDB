package controllers

import (
	"log"
	"net/http"

	"github.com/Shashwat5522/golang_gloabalsearch_mongo/models"
)

func CreateObjects(w http.ResponseWriter, r *http.Request) {

	var obj models.Object

	err := obj.CreateObjects()
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("objects created successfully!!!!"))
}

// func SearchObjects(w http.ResponseWriter, r *http.Request) {
// 	searchword := r.URL.Query()

// 	var obj models.Object
// 	resp, err := obj.SearchObject(searchword.Get("search"))

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	response, jsonErr := json.Marshal(resp)
// 	if jsonErr != nil {
// 		log.Fatal(jsonErr)
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(response)

// }
