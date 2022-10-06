package schedule

import (
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/widget/material"
	"keltas/widgets"
)

func Direction(th *material.Theme, value string) layout.Widget {
	label := widgets.NewLabel(th, 14, value)
	label.Color = colorDark
	label.Background = colorLight
	label.Alignment = text.Middle
	label.Font.Weight = text.UltraBlack
	return label.Layout
}
