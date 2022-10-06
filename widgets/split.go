package widgets

import (
	"gioui.org/layout"
	"gioui.org/op"
	"image"
)

type Split struct {
	Left, Right layout.Widget
}

func (w Split) Layout(gtx layout.Context) (dims layout.Dimensions) {
	leftSize := gtx.Constraints.Min.X / 2
	rightSize := gtx.Constraints.Min.X - leftSize

	{
		gtx := gtx
		gtx.Constraints = layout.Exact(image.Pt(leftSize, gtx.Constraints.Max.Y))
		w.Left(gtx)
	}

	{
		gtx := gtx
		gtx.Constraints = layout.Exact(image.Pt(rightSize, gtx.Constraints.Max.Y))
		trans := op.Offset(image.Pt(leftSize, 0)).Push(gtx.Ops)
		w.Right(gtx)
		trans.Pop()
	}

	dims = layout.Dimensions{Size: gtx.Constraints.Max}

	return
}

func NewSplit(left, right layout.Widget) Split {
	return Split{
		Left:  left,
		Right: right,
	}
}
