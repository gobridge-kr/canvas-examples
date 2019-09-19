package main

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/h8gi/canvas"
	"golang.org/x/image/colornames"
)

// WASD로 도형 이동
func main() {
	c := canvas.NewCanvas(&canvas.CanvasConfig{
		Width:     900,
		Height:    900,
		FrameRate: 60,
		Title:     "WASD",
	})

	c.Setup(func(ctx *canvas.Context) {

	})

	me := struct {
		X float64
		Y float64
	}{450, 450}
	c.Draw(func(ctx *canvas.Context) {
		ctx.SetColor(colornames.White)
		ctx.Clear()
		ctx.SetColor(colornames.Red)

		if ctx.IsKeyPressed(pixelgl.KeyW) {
			me.Y += 10
		} else if ctx.IsKeyPressed(pixelgl.KeyA) {
			me.X -= 10
		} else if ctx.IsKeyPressed(pixelgl.KeyS) {
			me.Y -= 10
		} else if ctx.IsKeyPressed(pixelgl.KeyD) {
			me.X += 10
		}

		ctx.DrawRectangle(me.X-50, me.Y-50, 100, 100)
		ctx.Stroke()
	})
}
