package window

import (
	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/widget/material"
	"image/color"
)

func Error(err error) func(gtx layout.Context) {
	return func(gtx layout.Context) {
		th := material.NewTheme(gofont.Collection())
		title := material.H3(th, "Error:")
		maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
		title.Color = maroon
		title.Alignment = text.Middle
		title.Layout(gtx)

		body := material.Body1(th, err.Error())
		body.Alignment = text.Start
		body.Layout(gtx)
	}
}
