package main

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/h8gi/canvas"
	"golang.org/x/image/colornames"
)

// 라이언 도장
func main() {
	c := canvas.NewCanvas(&canvas.CanvasConfig{
		Width:     900,
		Height:    700,
		FrameRate: 10,
		Title:     "Ryan",
	})

	c.Setup(func(ctx *canvas.Context) {
		ctx.SetColor(colornames.White)
		ctx.Clear()
	})

	c.Draw(func(ctx *canvas.Context) {
	ctx.Push()
		if ctx.IsMouseDragged {
			ctx.SetColor(colornames.Darkorange)
			ctx.DrawCircle(ctx.Mouse.X, ctx.Mouse.Y, 50)
			ctx.DrawCircle(ctx.Mouse.X-33, ctx.Mouse.Y+40, 10)
			ctx.DrawCircle(ctx.Mouse.X+33, ctx.Mouse.Y+40, 10)
			ctx.Fill() // 얼굴

			ctx.SetColor(colornames.Black)
			ctx.SetLineWidth(4)
			ctx.DrawLine(ctx.Mouse.X-31, ctx.Mouse.Y+17, ctx.Mouse.X-15, ctx.Mouse.Y+17)
			ctx.DrawLine(ctx.Mouse.X+15, ctx.Mouse.Y+17, ctx.Mouse.X+31, ctx.Mouse.Y+17)
			ctx.Stroke() // 눈썹

			ctx.DrawCircle(ctx.Mouse.X-23, ctx.Mouse.Y+7, 3)
			ctx.DrawCircle(ctx.Mouse.X+23, ctx.Mouse.Y+7, 3)
			ctx.Fill() // 눈

			ctx.SetColor(colornames.White)
			ctx.DrawCircle(ctx.Mouse.X-5, ctx.Mouse.Y-13, 7)
			ctx.DrawCircle(ctx.Mouse.X+5, ctx.Mouse.Y-13, 7)
			ctx.Fill() // 코1

			ctx.SetColor(colornames.Black)
			ctx.DrawCircle(ctx.Mouse.X, ctx.Mouse.Y-7, 3)
			ctx.Fill() // 코2
		}
		ctx.Pop()

		if ctx.IsKeyPressed(pixelgl.KeyUp) { // 키보드 이벤트(지우기)
			ctx.Push()
			ctx.SetColor(colornames.White)
			ctx.Clear()
			ctx.Pop()
		}
	})
}
