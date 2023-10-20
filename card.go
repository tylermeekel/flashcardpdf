package main

import (
	"log"
	"os"
	"strings"
)

type Card struct {
	question string
	answer   string
}

func getCardsList() []Card {
	cards := []Card{}
	rawData, err := os.ReadFile("cards.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	//split each card into a slice of strings
	cardsStrings := strings.Split(string(rawData), "####")
	for _, cardString := range cardsStrings {
		var card Card

		//split each card string into a slice containing question and answer strings
		questionAnswerSlice := strings.Split(cardString, "##")

		//indicates end of array, so break to avoid out of bounds index error
		if len(questionAnswerSlice) == 1 {
			break
		}

		//Add data to struct and append it to cards slice
		card.question = questionAnswerSlice[0]
		card.answer = questionAnswerSlice[1]
		cards = append(cards, card)
	}

	return cards
}
