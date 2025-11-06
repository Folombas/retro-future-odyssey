package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Menu struct {
	options      []string
	selected     int
	animationTime float64
}

func NewMenu() *Menu {
	return &Menu{
		options: []string{
			"🚀 НАЧАТЬ ИГРУ",
			"⭐ ЛУЧШИЙ РЕЗУЛЬТАТ",
			"⚙️  НАСТРОЙКИ",
			"❌ ВЫХОД",
		},
		selected: 0,
	}
}

func (m *Menu) Update(game *Game) {
	m.animationTime += 1.0 / 60.0 // Предполагаем 60 FPS

	// Навигация по меню
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		m.selected = (m.selected - 1 + len(m.options)) % len(m.options)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		m.selected = (m.selected + 1) % len(m.options)
	}

	// Выбор опции
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		switch m.selected {
		case 0: // Начать игру
			game.state = StatePlaying
			game.Reset()
		case 1: // Лучший результат
			// Можно добавить просмотр рекордов
		case 2: // Настройки
			// Можно добавить настройки
		case 3: // Выход
			// Выход из игры
		}
	}
}

func (m *Menu) Draw(screen *ebiten.Image) {
	// Анимированный космический фон
	screen.Fill(color.RGBA{0, 0, 25, 255})
	m.drawAnimatedStars(screen)

	// Заголовок игры
	titleY := 120.0
	ebitenutil.DebugPrintAt(screen, "🚀 SPACE DEFENDER PRO", 250, int(titleY))
	ebitenutil.DebugPrintAt(screen, "Цифровая трансформация 2025", 280, int(titleY)+30)

	// Анимация подчеркивания
	lineWidth := 200 + math.Sin(m.animationTime)*20
	vector.DrawFilledRect(screen, 300, float32(titleY)+50, float32(lineWidth), 2, color.RGBA{0, 255, 255, 255}, false)

	// Опции меню
	optionY := 220.0
	for i, option := range m.options {
		y := optionY + float64(i)*60

		// Выделение выбранной опции
		if i == m.selected {
			// Анимация стрелки (пульсирующее смещение)
			arrowOffset := math.Sin(m.animationTime*5) * 3
			ebitenutil.DebugPrintAt(screen, "▶ "+option, 300 + int(arrowOffset), int(y))

			// Подсветка с пульсирующей прозрачностью
			alpha := uint8(30 + math.Sin(m.animationTime*5)*15)
			vector.DrawFilledRect(screen, 290, float32(y)-5, 400, 30, color.RGBA{255, 255, 255, alpha}, false)
		} else {
			ebitenutil.DebugPrintAt(screen, "  "+option, 300, int(y))
		}
	}

	// Нижняя информация
	ebitenutil.DebugPrintAt(screen, "↑↓: Выбрать • ENTER: Подтвердить • ESC: Выход", 250, 500)
	ebitenutil.DebugPrintAt(screen, "💪 Day 9: Создаём игры вместо того, чтобы в них играть!", 200, 550)
}

func (m *Menu) drawAnimatedStars(screen *ebiten.Image) {
	for i := 0; i < 100; i++ {
		x := float32(math.Sin(m.animationTime*0.5+float64(i))*400 + 400)
		y := float32(i * 6)
		size := float32(math.Sin(m.animationTime+float64(i)*0.1)*0.5 + 1.0)
		vector.DrawFilledCircle(screen, x, y, size, color.RGBA{255, 255, 255, 255}, false)
	}
}
