package window

import (
	"fmt"
	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/widget/material"
	"keltas/schedule"
	"keltas/widgets"
	ui "keltas/window/schedule"
)

func Schedule(f schedule.Ferry) func(gtx layout.Context) {
	if len(f.Schedules) != 2 {
		return Error(fmt.Errorf("no schedules found"))
	}
	shL := f.Schedules[1]
	rhR := f.Schedules[0]

	return func(gtx layout.Context) {
		th := material.NewTheme(gofont.Collection())
		title := layout.Rigid(ui.Title(th, f.Title))
		dir1 := layout.Rigid(ui.Direction(th, shL.Title))
		dir2 := layout.Rigid(ui.Direction(th, rhR.Title))

		splitLayout := layout.Flexed(100, widgets.Split{
			Left: func(g layout.Context) layout.Dimensions {
				return layout.Flex{
					Axis:      layout.Vertical,
					Alignment: layout.Start,
				}.Layout(g, dir1, layout.Rigid(ui.Times(th, shL.Table)))
			},
			Right: func(g layout.Context) layout.Dimensions {
				return layout.Flex{
					Axis:      layout.Vertical,
					Alignment: layout.Start,
				}.Layout(g, dir2, layout.Rigid(ui.Times(th, rhR.Table)))
			},
		}.Layout)

		layout.Flex{
			Axis:      layout.Vertical,
			Alignment: layout.Start,
		}.Layout(gtx, title, splitLayout)
	}
}
