package main

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	"github.com/johnfercher/maroto/v2/pkg/consts/breakline"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

var colStyle = &props.Cell{
	BorderType:      border.Full,
	BorderColor:     &props.BlackColor,
	LineStyle:       linestyle.Solid,
	BorderThickness: 0.5,
}

var questionTextStyle = props.Text{
	Size:              16,
	Align:             align.Center,
	Top:               10,
	BreakLineStrategy: breakline.EmptyLineStrategy,
}
