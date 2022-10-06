package widgets

import (
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"image"
	"image/color"
)

type Label struct {
	Font       text.Font
	Color      color.NRGBA
	Background color.NRGBA
	Alignment  text.Alignment
	Text       string
	TextSize   unit.Sp
	shaper     text.Shaper
}

func (l Label) Layout(gtx layout.Context) (dims layout.Dimensions) {
	paint.FillShape(gtx.Ops, l.Background, clip.Rect{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: gtx.Constraints.Max.X, Y: gtx.Sp(l.TextSize + 4)},
	}.Op())

	paint.ColorOp{Color: l.Color}.Add(gtx.Ops)
	tl := widget.Label{Alignment: l.Alignment, MaxLines: 1}
	dims = tl.Layout(gtx, l.shaper, l.Font, l.TextSize, l.Text)

	return
}

func NewLabel(th *material.Theme, size unit.Sp, txt string) Label {
	return Label{
		Text:       txt,
		Color:      th.Palette.Fg,
		Background: th.Palette.Bg,
		TextSize:   size,
		shaper:     th.Shaper,
	}
}
