package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    fmt.Println("🖥️  ЗАГРУЗКА RETRO MEMORY MANAGER...")
    fmt.Println("=====================================")
    fmt.Println()
    fmt.Println("💾 Добро пожаловать, Гоша!")
    fmt.Println("📟 Твоя миссия: очистить память от цифровых наркотиков")
    fmt.Println("🐹 И освободить место для Go программирования!")
    fmt.Println()

    mm := NewMemoryManager()
    scanner := bufio.NewScanner(os.Stdin)

    for {
        mm.ShowMemory()

        fmt.Println()
        fmt.Println("⚙️  КОМАНДЫ:")
        fmt.Println("   free games    - очистить память от игр")
        fmt.Println("   free movies   - очистить память от фильмов")
        fmt.Println("   add <проект> <размер> - добавить Go проект")
        fmt.Println("   motivation    - показать мотивацию")
        fmt.Println("   exit          - выход")
        fmt.Println()
        fmt.Print("🚀 Введите команду: ")

        if !scanner.Scan() {
            break
        }

        input := strings.TrimSpace(scanner.Text())
        command, argument := ParseInput(input)

        switch command {
        case "free":
            if argument == "games" {
                mm.FreeGameMemory()
            } else if argument == "movies" {
                mm.FreeMovieMemory()
            }
        case "add":
            parts := strings.Fields(argument)
            if len(parts) >= 2 {
                size, err := strconv.Atoi(parts[len(parts)-1])
                if err == nil {
                    name := strings.Join(parts[:len(parts)-1], " ")
                    mm.AddGoProject(name, size)
                }
            }
        case "motivation":
            mm.ShowMotivation()
        case "exit", "quit":
            fmt.Println()
            fmt.Println("👋 До свидания, Гоша!")
            fmt.Println("💪 Продолжай освобождать память для своего будущего!")
            return
        default:
            fmt.Println("❌ Неизвестная команда. Попробуйте снова.")
        }

        fmt.Println()
        fmt.Print("⏎ Нажмите Enter для продолжения...")
        scanner.Scan()
    }
}
