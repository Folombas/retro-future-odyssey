package main

import "fmt"

func main() {
    fmt.Println("🎮 Day 8: Retro Terminal Pong")
    fmt.Println("==============================")

    game := NewPongGame()
    game.Start()

    fmt.Println("\n🏆 Ты отлично провёл время за кодом вместо игр!")
}
