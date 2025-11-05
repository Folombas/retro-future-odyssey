package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Enemy struct {
	X, Y   float32
	Speed  float32
	Width  float32
	Height float32
}

func NewEnemy() *Enemy {
	return &Enemy{
		X:      float32(rand.Intn(700) + 50),
		Y:      -30,
		Speed:  float32(rand.Intn(3) + 2),
		Width:  24,
		Height: 24,
	}
}

func (e *Enemy) Update() {
	e.Y += e.Speed
}

func (e *Enemy) Draw(screen *ebiten.Image) {
	// Основной корпус врага
	vector.DrawFilledCircle(screen, e.X+e.Width/2, e.Y+e.Height/2, e.Width/2, color.RGBA{255, 50, 50, 255}, false)

	// Глаза
	vector.DrawFilledCircle(screen, e.X+8, e.Y+8, 3, color.RGBA{0, 0, 0, 255}, false)
	vector.DrawFilledCircle(screen, e.X+16, e.Y+8, 3, color.RGBA{0, 0, 0, 255}, false)

	// Рот
	vector.DrawFilledRect(screen, e.X+8, e.Y+16, 8, 3, color.RGBA{0, 0, 0, 255}, false)
}

func (e *Enemy) GetX() float32      { return e.X }
func (e *Enemy) GetY() float32      { return e.Y }
func (e *Enemy) GetWidth() float32  { return e.Width }
func (e *Enemy) GetHeight() float32 { return e.Height }
