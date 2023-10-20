package main

import (
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	"github.com/johnfercher/maroto/v2/pkg/consts/breakline"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

// Styles

var colStyle = &props.Cell{
	BorderType:      border.Full,
	BorderColor:     &props.BlackColor,
	LineStyle:       linestyle.Solid,
	BorderThickness: 0.5,
}

func getMaroto(cards []Card, questionFontSize, answerFontSize float64) core.Maroto {

	var questionTextStyle = props.Text{
		Size:              questionFontSize,
		Align:             align.Center,
		Top:               10,
		BreakLineStrategy: breakline.EmptyLineStrategy,
	}

	var answerTextStyle = props.Text{
		Size:  answerFontSize,
		Align: align.Center,
		Top:   10,
	}

	cfg := config.NewBuilder().Build()
	m := maroto.New(cfg)

	rows := []core.Row{}

	//Create list of card rows to add to the
	for _, card := range cards {
		rows = append(rows,
			row.New(80).Add(
				text.NewCol(6, card.question, questionTextStyle).WithStyle(colStyle),
				text.NewCol(6, card.answer, answerTextStyle).WithStyle(colStyle)),
			row.New(5),
		)
	}

	m.AddRows(rows...)

	return m
}
