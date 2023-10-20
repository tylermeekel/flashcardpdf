package main

import (
	"log"
)

func main() {
	cards := getCardsList()
	m := getMaroto(cards)

	//Generate document
	doc, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	//Save document
	err = doc.Save("docs/flashcards.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

}
