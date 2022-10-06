package window

import (
	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/widget/material"
	"image/color"
)

func Loading() func(gtx layout.Context) {
	return func(gtx layout.Context) {
		th := material.NewTheme(gofont.Collection())
		title := material.H3(th, "Loading...")
		maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
		title.Color = maroon
		title.Alignment = text.Middle
		title.Layout(gtx)
	}
}
