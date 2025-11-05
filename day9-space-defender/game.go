package main

import (
	"image/color"
	"time"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	player    *Player
	enemies   []*Enemy
	bullets   []*Bullet
	score     int
	lastSpawn time.Time
	gameOver  bool
}

func NewGame() *Game {
	return &Game{
		player:    NewPlayer(),
		score:     0,
		lastSpawn: time.Now(),
		gameOver:  false,
	}
}

func (g *Game) Update() error {
	if g.gameOver {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			// Рестарт игры
			g.player = NewPlayer()
			g.enemies = nil
			g.bullets = nil
			g.score = 0
			g.gameOver = false
			g.lastSpawn = time.Now()
		}
		return nil
	}

	g.player.Update(g)

	// Спавн врагов каждые 1.5 секунды
	if time.Since(g.lastSpawn) > 1500*time.Millisecond && len(g.enemies) < 8 {
		g.enemies = append(g.enemies, NewEnemy())
		g.lastSpawn = time.Now()
	}

	// Обновление врагов
	for i := len(g.enemies) - 1; i >= 0; i-- {
		g.enemies[i].Update()

		// Проверка столкновения с игроком
		if g.CheckCollision(g.player, g.enemies[i]) {
			g.gameOver = true
			return nil
		}

		// Удаление врагов за экраном
		if g.enemies[i].Y > 600 {
			g.enemies = append(g.enemies[:i], g.enemies[i+1:]...)
			g.score += 5 // Очки за уклонение
		}
	}

	// Обновление пуль
	for i := len(g.bullets) - 1; i >= 0; i-- {
		g.bullets[i].Update()

		// Удаление пуль за экраном
		if g.bullets[i].Y < -10 {
			g.bullets = append(g.bullets[:i], g.bullets[i+1:]...)
			continue
		}

		// Проверка столкновения пуль с врагами
		for j := len(g.enemies) - 1; j >= 0; j-- {
			if g.CheckCollision(g.bullets[i], g.enemies[j]) {
				g.score += 20 // Очки за уничтожение
				g.enemies = append(g.enemies[:j], g.enemies[j+1:]...)
				g.bullets = append(g.bullets[:i], g.bullets[i+1:]...)
				break
			}
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Космический фон
	screen.Fill(color.RGBA{0, 0, 25, 255})

	// Рисуем звёзды
	for i := 0; i < 50; i++ {
		x := float32(i * 16)
		y := float32((i * 7) % 600)
		vector.DrawFilledCircle(screen, x, y, 1, color.RGBA{255, 255, 255, 255}, false)
	}

	// Отрисовка игрока
	g.player.Draw(screen)

	// Отрисовка врагов
	for _, enemy := range g.enemies {
		enemy.Draw(screen)
	}

	// Отрисовка пуль
	for _, bullet := range g.bullets {
		bullet.Draw(screen)
	}

	// Интерфейс
	scoreText := fmt.Sprintf("🎯 Score: %d", g.score)
	ebitenutil.DebugPrint(screen, scoreText)
	ebitenutil.DebugPrintAt(screen, "🚀 ARROWS: Move | SPACE: Shoot | ESC: Quit", 10, 580)

	if g.gameOver {
		ebitenutil.DebugPrintAt(screen, "💀 GAME OVER! Press SPACE to restart", 250, 280)
		finalScore := fmt.Sprintf("🎯 Final Score: %d", g.score)
		ebitenutil.DebugPrintAt(screen, finalScore, 320, 300)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}

func (g *Game) CheckCollision(a, b Collider) bool {
	return a.GetX() < b.GetX()+b.GetWidth() &&
		a.GetX()+a.GetWidth() > b.GetX() &&
		a.GetY() < b.GetY()+b.GetHeight() &&
		a.GetY()+a.GetHeight() > b.GetY()
}

func (g *Game) AddBullet(bullet *Bullet) {
	g.bullets = append(g.bullets, bullet)
}
