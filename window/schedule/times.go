package schedule

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"image/color"
	"keltas/schedule"
	"keltas/widgets"
)

func Times(th *material.Theme, times []schedule.Time) layout.Widget {
	labels := make([]layout.FlexChild, 0, len(times))
	for _, t := range times {
		hour := widgets.NewLabel(th, 20, fmt.Sprintf(" %s: ", t.Hour))
		hour.Color = colorDark
		hour.Font.Weight = text.Bold
		hour.Alignment = text.Start

		minutes := widgets.NewLabel(th, 16, t.Minutes)
		minutes.Color = colorDark
		minutes.Font.Weight = text.Bold
		minutes.Alignment = text.Start

		labels = append(labels, layout.Rigid(
			func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{
					Axis:      layout.Horizontal,
					Alignment: layout.Start,
				}.Layout(gtx, layout.Rigid(hour.Layout), layout.Rigid(minutes.Layout))
			},
		))
	}

	return func(gtx layout.Context) (dims layout.Dimensions) {
		flexWidget := func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{
				Axis:      layout.Vertical,
				Alignment: layout.Start,
			}.Layout(gtx, labels...)
		}

		widget.Border{
			Color:        color.NRGBA{A: 0xFF},
			CornerRadius: 0,
			Width:        1,
		}.Layout(gtx, flexWidget)

		return
	}
}
