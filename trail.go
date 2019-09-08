package main

import (
	"image/color"
	"math/rand"

	"github.com/h8gi/canvas"
	"golang.org/x/image/colornames"
)

// 마우스 지나간 자국
func main() {
	c := canvas.New(&canvas.NewCanvasOptions{
		Width:     1200,
		Height:    1200,
		FrameRate: 60,
		Title:     "Trail",
	})

	c.Setup(func(ctx *canvas.Context) {
		ctx.SetColor(colornames.Black)
		ctx.Clear()
	})

	c.Draw(func(ctx *canvas.Context) {
		if ctx.MouseX() == ctx.PreviousMouseX() && ctx.MouseY() == ctx.PreviousMouseY() {
			return
		}
		ctx.SetColor(color.RGBA{
			R: uint8(rand.Intn(256)),
			G: uint8(rand.Intn(256)),
			B: uint8(rand.Intn(256)),
			A: 255,
		})
		ctx.DrawCircle(
			ctx.PreviousMouseX(),
			ctx.PreviousMouseY(),
			15,
		)
		ctx.Fill()
		ctx.Stroke()
	})
}
