package main

import "fmt"

type MemoryBlock struct {
    ID      int
    Name    string
    Size    int
    Type    string // "game", "movie", "go_code", "system"
    Active  bool
}

type MemoryManager struct {
    blocks []MemoryBlock
    nextID int
}

func NewMemoryManager() *MemoryManager {
    return &MemoryManager{
        blocks: []MemoryBlock{
            {ID: 1, Name: "GTA Vice City", Size: 1500, Type: "game", Active: true},
            {ID: 2, Name: "The Sims 3", Size: 2000, Type: "game", Active: true},
            {ID: 3, Name: "Movies Collection", Size: 5000, Type: "movie", Active: true},
            {ID: 4, Name: "Go Compiler", Size: 300, Type: "go_code", Active: true},
        },
        nextID: 5,
    }
}

func (mm *MemoryManager) ShowMemory() {
    fmt.Println("╔══════════════════════════════════════════════╗")
    fmt.Println("║           RETRO MEMORY MANAGER v1.0         ║")
    fmt.Println("║                 (c) 2025 RFO               ║")
    fmt.Println("╚══════════════════════════════════════════════╝")
    fmt.Println()

    totalMemory := 0
    usedByGames := 0
    usedByGo := 0

    for _, block := range mm.blocks {
        if block.Active {
            status := "🟢 АКТИВЕН"
            if block.Type == "game" {
                status = "🎮 ИГРА"
                usedByGames += block.Size
            } else if block.Type == "movie" {
                status = "🎬 ФИЛЬМ"
                usedByGames += block.Size
            } else if block.Type == "go_code" {
                status = "🐹 GO КОД"
                usedByGo += block.Size
            } else {
                status = "⚙️  СИСТЕМА"
            }

            fmt.Printf("▓ %-25s %6d MB   %s\n", block.Name, block.Size, status)
            totalMemory += block.Size
        }
    }

    fmt.Println()
    fmt.Printf("📊 СТАТИСТИКА ПАМЯТИ:\n")
    fmt.Printf("   Всего занято: %d MB\n", totalMemory)
    fmt.Printf("   Игры/фильмы: %d MB\n", usedByGames)
    fmt.Printf("   Go программы: %d MB\n", usedByGo)
    fmt.Printf("   Свободно для Go: %d MB\n", usedByGames)
    fmt.Println()
}
