package main

import (
	"flag"
	"log"
)

func main() {
	var questionFontSize float64
	var answerFontSize float64
	var questionAnswerDelineator string
	var cardDelineator string
	var inputFilePath string
	var outputFilePath string
	flag.Float64Var(&questionFontSize, "qfs", 25, "Sets the font size for the question side of the flashcard.")
	flag.Float64Var(&answerFontSize, "afs", 20, "Sets the font size for the answer side of the flashcard.")
	flag.StringVar(&questionAnswerDelineator, "qd", "##", "Sets the delineator in between a question and answer. (Can cause errors)")
	flag.StringVar(&cardDelineator, "cd", "####", "Sets the delineator in between each card. (Can cause errors)")
	flag.StringVar(&inputFilePath, "i", "./cards.txt", "Sets the input filepath.")
	flag.StringVar(&outputFilePath, "o", "./docs/flashcards.pdf", "Sets the output filepath. Make sure to include the \".pdf\" extension")

	flag.Parse()
	cards := getCardsList(inputFilePath)
	m := getMaroto(cards, questionFontSize, answerFontSize)

	//Generate document
	doc, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	//Save document
	err = doc.Save(outputFilePath)
	if err != nil {
		log.Fatal(err.Error())
	}

}
