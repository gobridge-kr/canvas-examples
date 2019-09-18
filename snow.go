package main

import (
	"math/rand"

	"github.com/h8gi/canvas"
	"golang.org/x/image/colornames"
)

const width = 1200
const height = 1200
const n = 100

// 눈 내리는 효과
func main() {
	c := canvas.NewCanvas(&canvas.CanvasConfig{
		Width:     width,
		Height:    height,
		FrameRate: 60,
		Title:     "Snow",
	})

	particles := [n]SnowParticle{}
	c.Setup(func(ctx *canvas.Context) {
		for i := 0; i < n; i += 1 {
			particles[i] = randomSnowParticle()
		}
	})

	c.Draw(func(ctx *canvas.Context) {
		ctx.SetColor(colornames.Aqua)
		ctx.Clear()
		ctx.SetLineWidth(10)
		ctx.SetColor(colornames.White)
		for i, particle := range particles {
			ctx.DrawCircle(particle.X, particle.Y, 10)
			ctx.Fill()
			newY := particle.Y + particle.Velocity
			if newY > height {
				particles[i] = randomSnowParticle()
				continue
			}
			particles[i] = SnowParticle{
				X:        particle.X,
				Y:        newY,
				Velocity: particle.Velocity * (1 + rand.Float64()*0.01),
			}
		}
		ctx.Stroke()
	})
}

func randomSnowParticle() SnowParticle {
	return SnowParticle{
		X:        rand.Float64() * width,
		Y:        0,
		Velocity: rand.Float64() * height * 0.01,
	}
}

type SnowParticle struct {
	X        float64
	Y        float64
	Velocity float64
}
