package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Bullet struct {
	X, Y   float32
	Speed  float32
	Width  float32
	Height float32
}

func NewBullet(x, y float32) *Bullet {
	return &Bullet{
		X:      x - 2,
		Y:      y,
		Speed:  7,
		Width:  4,
		Height: 12,
	}
}

func (b *Bullet) Update() {
	b.Y -= b.Speed
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, b.X, b.Y, b.Width, b.Height, color.RGBA{255, 255, 0, 255}, false)
}

func (b *Bullet) GetX() float32      { return b.X }
func (b *Bullet) GetY() float32      { return b.Y }
func (b *Bullet) GetWidth() float32  { return b.Width }
func (b *Bullet) GetHeight() float32 { return b.Height }
