package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func startWebServer() {
	mux := chi.NewMux()

	mux.Get("/", handleIndex)
	mux.Post("/pdf", handlePostPDF)

	fmt.Println("starting server on port :3000")
	log.Fatal(http.ListenAndServe(":3000", mux))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	file, err := os.ReadFile("index.html")
	if err != nil {
		fmt.Println(err.Error())
	}
	w.Header().Set("Content-type", "text/html")
	w.Write(file)
}

func handlePostPDF(w http.ResponseWriter, r *http.Request) {
	data := r.FormValue("data")
	if data == "" {
		w.Write([]byte("Error"))
		return
	}

	cards := getCardsList(data)
	m := getMaroto(cards, 20.0, 12.0)

	docs, err := m.Generate()
	if err != nil {
		fmt.Println(err.Error())
		w.Write([]byte("Error"))
	}

	bytes := docs.GetBytes()
	w.Write(bytes)
}
