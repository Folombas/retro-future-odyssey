package main

import (
    "fmt"
    "strings"
)

func (mm *MemoryManager) FreeGameMemory() {
    freed := 0
    for i := range mm.blocks {
        if mm.blocks[i].Type == "game" && mm.blocks[i].Active {
            freed += mm.blocks[i].Size
            mm.blocks[i].Active = false
        }
    }

    if freed > 0 {
        fmt.Printf("✅ Очищено %d MB от игр!\n", freed)
        fmt.Println("🎯 Теперь у тебя есть место для изучения Go!")
    } else {
        fmt.Println("ℹ️  Игр в памяти не найдено.")
    }
}

func (mm *MemoryManager) FreeMovieMemory() {
    freed := 0
    for i := range mm.blocks {
        if mm.blocks[i].Type == "movie" && mm.blocks[i].Active {
            freed += mm.blocks[i].Size
            mm.blocks[i].Active = false
        }
    }

    if freed > 0 {
        fmt.Printf("✅ Очищено %d MB от фильмов!\n", freed)
        fmt.Println("🎬 Освобождённое время теперь можно потратить на программирование!")
    } else {
        fmt.Println("ℹ️  Фильмов в памяти не найдено.")
    }
}

func (mm *MemoryManager) AddGoProject(name string, size int) {
    mm.blocks = append(mm.blocks, MemoryBlock{
        ID:     mm.nextID,
        Name:   name,
        Size:   size,
        Type:   "go_code",
        Active: true,
    })
    mm.nextID++

    fmt.Printf("✅ Добавлен Go проект: %s (%d MB)\n", name, size)
    fmt.Println("🐹 Отличный выбор! Этот проект приблизит тебя к работе мечты!")
}

func (mm *MemoryManager) ShowMotivation() {
    fmt.Println()
    fmt.Println("💡 СОВЕТ ОТ СИСТЕМЫ:")
    fmt.Println("   Каждый освобожденный мегабайт от игр - это")
    fmt.Println("   шаг к тёплому офису и достойной зарплате!")
    fmt.Println("   Помни: время, потраченное на Go, окупится")
    fmt.Println("   в 100 раз больше, чем время за играми!")
    fmt.Println()
}

func ParseInput(input string) (string, string) {
    parts := strings.Fields(input)
    if len(parts) == 0 {
        return "", ""
    }

    command := strings.ToLower(parts[0])
    argument := ""
    if len(parts) > 1 {
        argument = strings.Join(parts[1:], " ")
    }

    return command, argument
}
