package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Player struct {
	X, Y float32
	Speed float32
}

func NewPlayer() *Player {
	return &Player{
		X: 400,
		Y: 500,
		Speed: 5,
	}
}

func (p *Player) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) && p.X > 20 {
		p.X -= p.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) && p.X < 780 {
		p.X += p.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) && p.Y > 20 {
		p.Y -= p.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) && p.Y < 580 {
		p.Y += p.Speed
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, p.X-15, p.Y-10, 30, 20, color.RGBA{0, 255, 0, 255}, false)
}
