package main

import (
	"fmt"
	"log"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type Card struct {
	question string
	answer   string
}

func main() {
	cfg := config.NewBuilder().Build()
	mrt := maroto.New(cfg)

	answerTextStyle := props.Text{
		Size:  12,
		Align: align.Center,
		Top:   10,
	}

	rows := []core.Row{}

	for i := 0; i < 10; i++ {
		rows = append(rows,
			row.New(80).Add(
				text.NewCol(6, fmt.Sprintf("row %d col1", i), questionTextStyle).WithStyle(colStyle),
				text.NewCol(6, fmt.Sprintf("row %d col2", i), answerTextStyle).WithStyle(colStyle)),
			row.New(5),
		)
	}

	mrt.AddRows(rows...)

	doc, err := mrt.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = doc.Save("docs/test.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

}
