package main

import (
	"github.com/h8gi/canvas"
	"golang.org/x/image/colornames"
)

// 마우스로 선 그리기
func main() {
	c := canvas.New(&canvas.NewCanvasOptions{
		Width:     1200,
		Height:    1200,
		FrameRate: 60,
		Title:     "Canvas",
	})

	c.Setup(func(ctx *canvas.Context) {
		ctx.SetColor(colornames.White)
		ctx.Clear()
		ctx.SetColor(colornames.Black)
		ctx.SetLineWidth(10)
	})

	c.Draw(func(ctx *canvas.Context) {
		if ctx.MouseDragged() {
			ctx.DrawLine(
				ctx.MouseX(),
				ctx.MouseY(),
				ctx.PreviousMouseX(),
				ctx.PreviousMouseY(),
			)
			ctx.Stroke()
		}

		if ctx.KeyPressed() {
			ctx.SetColor(colornames.White)
			ctx.Clear()
			ctx.SetColor(colornames.Black)
		}
	})
}
