package main

import (
	"github.com/fogleman/gg"
	"github.com/h8gi/canvas"
	"golang.org/x/image/colornames"
)

// 사람 그리기
func main() {
	c := canvas.New(&canvas.NewCanvasOptions{
		Width:     1200,
		Height:    1200,
		FrameRate: 60,
		Title:     "Basic",
	})

	c.Setup(func(ctx *canvas.Context) {
		ctx.SetColor(colornames.Black)
		ctx.Clear()

		ctx.SetColor(colornames.White)
		ctx.SetLineWidth(5)

		ctx.DrawArc(560, 300, 25, gg.Radians(-30), gg.Radians(-150)) // left eye
		ctx.Stroke()

		ctx.DrawArc(640, 300, 25, gg.Radians(-30), gg.Radians(-150)) // right eye
		ctx.Stroke()

		ctx.DrawArc(600, 325, 50, gg.Radians(30), gg.Radians(150)) // mouth
		ctx.Stroke()

		ctx.SetLineWidth(15)
		ctx.DrawCircle(600, 300, 100)     // head
		ctx.DrawLine(600, 400, 600, 700)  // body
		ctx.DrawLine(600, 500, 400, 300)  // left arm
		ctx.DrawLine(600, 500, 800, 300)  // right arm
		ctx.DrawLine(600, 700, 500, 1000) // left leg
		ctx.DrawLine(600, 700, 700, 1000) // right leg
		ctx.Stroke()
	})

	c.Draw(func(ctx *canvas.Context) {

	})
}
