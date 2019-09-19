package main

import (
	"image/color"
	"math/rand"

	"github.com/h8gi/canvas"
	"golang.org/x/image/colornames"
)

// 마우스 지나간 자국
func main() {
	c := canvas.NewCanvas(&canvas.CanvasConfig{
		Width:     900,
		Height:    900,
		FrameRate: 60,
		Title:     "Trail",
	})

	c.Setup(func(ctx *canvas.Context) {
		ctx.SetColor(colornames.Black)
		ctx.Clear()
	})

	c.Draw(func(ctx *canvas.Context) {
		if ctx.Mouse.X == ctx.PMouse.X && ctx.Mouse.Y == ctx.PMouse.Y {
			return
		}
		ctx.SetColor(color.RGBA{
			R: uint8(rand.Intn(256)),
			G: uint8(rand.Intn(256)),
			B: uint8(rand.Intn(256)),
			A: 255,
		})
		ctx.DrawCircle(
			ctx.PMouse.X,
			ctx.PMouse.Y,
			15,
		)
		ctx.Fill()
		ctx.Stroke()
	})
}
