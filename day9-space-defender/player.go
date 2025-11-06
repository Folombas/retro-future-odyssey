package main

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Player struct {
	X, Y     float32
	Speed    float32
	Width    float32
	Height   float32
	lastShot time.Time
}

func NewPlayer() *Player {
	return &Player{
		X:      400,
		Y:      500,
		Speed:  5,
		Width:  30,
		Height: 20,
		lastShot: time.Now(),
	}
}

func (p *Player) Update(game *Game) {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) && p.X > 20 {
		p.X -= p.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) && p.X < 750 {
		p.X += p.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) && p.Y > 20 {
		p.Y -= p.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) && p.Y < 550 {
		p.Y += p.Speed
	}

	// Стрельба пробелом
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if time.Since(p.lastShot) > 200*time.Millisecond {
			bullet := NewBullet(p.X + p.Width/2, p.Y)
			game.AddBullet(bullet)
			p.lastShot = time.Now()
		}
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	// Корпус корабля
	vector.DrawFilledRect(screen, p.X, p.Y, p.Width, p.Height, color.RGBA{0, 255, 0, 255}, false)

	// Кабина
	vector.DrawFilledCircle(screen, p.X+p.Width/2, p.Y+5, 6, color.RGBA{100, 200, 255, 255}, false)

	// Крылья
	vector.DrawFilledRect(screen, p.X-10, p.Y+5, 10, 8, color.RGBA{0, 200, 0, 255}, false)
	vector.DrawFilledRect(screen, p.X+p.Width, p.Y+5, 10, 8, color.RGBA{0, 200, 0, 255}, false)
}

func (p *Player) GetX() float32      { return p.X }
func (p *Player) GetY() float32      { return p.Y }
func (p *Player) GetWidth() float32  { return p.Width }
func (p *Player) GetHeight() float32 { return p.Height }
