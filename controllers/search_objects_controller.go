package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Shashwat5522/golang_gloabalsearch_mongo/models"
)

func SearchObjects(w http.ResponseWriter, r *http.Request) {
	searchword := r.URL.Query()

	var obj models.Object
	if searchword.Get("search")==""{
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))
		return
	}
	resp, err := obj.SearchObject(searchword.Get("search"))

	if err != nil {
		log.Fatal(err)
	}
	response, jsonErr := json.Marshal(resp)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}
