package main

import (
	"github.com/fogleman/gg"
	"github.com/h8gi/canvas"
	"golang.org/x/image/colornames"
)

// 사람 그리기
func main() {
	c := canvas.NewCanvas(&canvas.CanvasConfig{
		Width:     900,
		Height:    900,
		FrameRate: 60,
		Title:     "Basic",
	})

	c.Setup(func(ctx *canvas.Context) {
		ctx.SetColor(colornames.Black)
		ctx.Clear()

		ctx.SetColor(colornames.White)
		ctx.SetLineWidth(5)

		ctx.DrawArc(420, 675, 25, gg.Radians(30), gg.Radians(150)) // left eye
		ctx.Stroke()

		ctx.DrawArc(480, 675, 25, gg.Radians(30), gg.Radians(150)) // right eye
		ctx.Stroke()

		ctx.DrawArc(450, 665, 50, gg.Radians(-30), gg.Radians(-150)) // mouth
		ctx.Stroke()

		ctx.SetLineWidth(15)
		ctx.DrawCircle(450, 675, 100)    // head
		ctx.DrawLine(450, 570, 450, 375) // body
		ctx.DrawLine(450, 475, 300, 650) // left arm
		ctx.DrawLine(450, 475, 600, 650) // right arm
		ctx.DrawLine(450, 375, 375, 150) // left leg
		ctx.DrawLine(450, 375, 525, 150) // right leg
		ctx.Stroke()
	})

	c.Draw(func(ctx *canvas.Context) {
	})
}
