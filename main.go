package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func main() {

	port := 8000

	//region Handlers

	indexHandler := func(responseWriter http.ResponseWriter, request *http.Request) {
		parsedFiles := template.Must(template.ParseFiles("index.html"))

		parsedFiles.Execute(responseWriter, "")
	}

	decayHandler := func(w http.ResponseWriter, r *http.Request) {
		dose := r.PostFormValue("dose")
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "decayed-dose", dose)
	}

	//endregion

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/calc-decay/", decayHandler)

	fmt.Println("Listening on port " + strconv.Itoa(port) + "...")
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}
