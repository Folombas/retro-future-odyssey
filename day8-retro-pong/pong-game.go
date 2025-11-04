package main

import (
	"fmt"
	"time"
)

type PongGame struct {
	ballPos   int
	direction int
	score     int
}

func NewPongGame() *PongGame {
	return &PongGame{
		ballPos:   10,
		direction: 1,
		score:     0,
	}
}

func (g *PongGame) Start() {
	fmt.Println("🚀 Запускаем ретро-Pong...")
	time.Sleep(1 * time.Second)

	for i := 0; i < 20; i++ {
		g.Update()
		g.Render()
		time.Sleep(200 * time.Millisecond)
	}
}

func (g *PongGame) Update() {
	g.ballPos += g.direction

	if g.ballPos <= 0 || g.ballPos >= 20 {
		g.direction = -g.direction
		g.score++
	}
}

func (g *PongGame) Render() {
	field := make([]rune, 21)
	for i := range field {
		field[i] = ' '
	}

	field[g.ballPos] = '●'
	field[0] = '|'
	field[20] = '|'

	fmt.Printf("🎯 Счёт: %d | Поле: %s\n", g.score, string(field))
}
