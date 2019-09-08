package main

import (
	"github.com/h8gi/canvas"
	"golang.org/x/image/colornames"
	"golang.org/x/mobile/event/key"
)

// WASD로 도형 이동
func main() {
	c := canvas.New(&canvas.NewCanvasOptions{
		Width:     1200,
		Height:    1200,
		FrameRate: 60,
		Title:     "WASD",
	})

	c.Setup(func(ctx *canvas.Context) {

	})

	me := struct {
		X float64
		Y float64
	}{600, 600}
	c.Draw(func(ctx *canvas.Context) {
		ctx.SetColor(colornames.White)
		ctx.Clear()
		ctx.SetColor(colornames.Red)

		if ctx.KeyPressed() {
			e := ctx.KeyEvent()
			if e.Code == key.CodeW {
				me.Y -= 10
			} else if e.Code == key.CodeA {
				me.X -= 10
			} else if e.Code == key.CodeS {
				me.Y += 10
			} else if e.Code == key.CodeD {
				me.X += 10
			}
		}

		ctx.DrawRectangle(me.X-50, me.Y-50, 100, 100)
		ctx.Stroke()
	})
}
