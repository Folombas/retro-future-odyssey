package main

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	player   *Player
	enemies  []*Enemy
	score    int
	lastSpawn time.Time
}

func NewGame() *Game {
	return &Game{
		player:    NewPlayer(),
		score:     0,
		lastSpawn: time.Now(),
	}
}

func (g *Game) Update() error {
	g.player.Update()

	// Спавн врагов каждые 2 секунды
	if time.Since(g.lastSpawn) > 2*time.Second && len(g.enemies) < 5 {
		g.enemies = append(g.enemies, NewEnemy())
		g.lastSpawn = time.Now()
	}

	// Обновление врагов
	for i := len(g.enemies) - 1; i >= 0; i-- {
		g.enemies[i].Update()
		if g.enemies[i].Y > 600 {
			g.enemies = append(g.enemies[:i], g.enemies[i+1:]...)
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 25, 255}) // Темно-синий космос

	// Отрисовка игрока
	g.player.Draw(screen)

	// Отрисовка врагов
	for _, enemy := range g.enemies {
		enemy.Draw(screen)
	}

	// Интерфейс
	ebitenutil.DebugPrint(screen, "🎯 Score: 0 | 🚀 Use ARROWS to move | ESC to quit")
	ebitenutil.DebugPrintAt(screen, "💪 Day 9: Replacing gaming addiction with code!", 10, 580)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}
