package controllers

import (
	"html/template"
	"log"
	"net/http"
)

func ShowIndexPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("C:/Users/Bacancy/Desktop/golang_practice/Golang_GlobalSearch/templates/index.html")
	if err != nil {
		log.Fatal(err)
	}
	if err := tmpl.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
}
