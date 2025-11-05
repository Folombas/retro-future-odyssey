package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Enemy struct {
	X, Y float32
	Speed float32
}

func NewEnemy() *Enemy {
	return &Enemy{
		X: float32(rand.Intn(700) + 50),
		Y: -30,
		Speed: float32(rand.Intn(3) + 2),
	}
}

func (e *Enemy) Update() {
	e.Y += e.Speed
}

func (e *Enemy) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, e.X, e.Y, 15, color.RGBA{255, 0, 0, 255}, false)
}
