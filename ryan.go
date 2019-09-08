package main

import (
	"github.com/h8gi/canvas"
	"golang.org/x/image/colornames"
)

// 라이언 도장
func main() {
	c := canvas.New(&canvas.NewCanvasOptions{
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
		if ctx.MouseDragged() {
			ctx.SetColor(colornames.Darkorange)
			ctx.DrawCircle(ctx.MouseX(), ctx.MouseY(), 50)
			ctx.DrawCircle(ctx.MouseX()-33, ctx.MouseY()-40, 10)
			ctx.DrawCircle(ctx.MouseX()+33, ctx.MouseY()-40, 10)
			ctx.Fill() // 얼굴

			ctx.SetColor(colornames.Black)
			ctx.SetLineWidth(4)
			ctx.DrawLine(ctx.MouseX()-31, ctx.MouseY()-17, ctx.MouseX()-15, ctx.MouseY()-17)
			ctx.DrawLine(ctx.MouseX()+15, ctx.MouseY()-17, ctx.MouseX()+31, ctx.MouseY()-17)
			ctx.Stroke() // 눈썹

			ctx.DrawCircle(ctx.MouseX()-23, ctx.MouseY()-7, 3)
			ctx.DrawCircle(ctx.MouseX()+23, ctx.MouseY()-7, 3)
			ctx.Fill() // 눈

			ctx.SetColor(colornames.White)
			ctx.DrawCircle(ctx.MouseX()-5, ctx.MouseY()+13, 7)
			ctx.DrawCircle(ctx.MouseX()+5, ctx.MouseY()+13, 7)
			ctx.Fill() // 코1

			ctx.SetColor(colornames.Black)
			ctx.DrawCircle(ctx.MouseX(), ctx.MouseY()+7, 3)
			ctx.Fill() // 코2
		}

		if ctx.KeyPressed() { // 키보드 이벤트(지우기)
			ctx.SetColor(colornames.White)
			ctx.Clear()
		}
	})
}
