package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("🔄 DIGITAL DETOX MANAGER v1.0")
	fmt.Println("══════════════════════════════")
	fmt.Println()
	fmt.Println("💡 Добро пожаловать, Гоша!")
	fmt.Println("🎯 Твоя миссия: превратить 25 лет игрового опыта")
	fmt.Println("   в 25 лет успешной карьеры Go-разработчика!")
	fmt.Println()

	detoxManager := NewDetoxManager()
	progressTracker := NewProgressTracker()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n" + strings.Repeat("═", 50))
		fmt.Println("🏠 ГЛАВНОЕ МЕНЮ:")
		fmt.Println("   1 - Показать вредные привычки")
		fmt.Println("   2 - Показать полезные привычки")
		fmt.Println("   3 - Преодолеть зависимость")
		fmt.Println("   4 - Добавить полезную привычку")
		fmt.Println("   5 - Записать время изучения Go")
		fmt.Println("   6 - Добавить завершённый проект")
		fmt.Println("   7 - Показать прогресс")
		fmt.Println("   8 - Получить мотивацию")
		fmt.Println("   9 - Выход")
		fmt.Print("🎯 Выберите действие: ")

		if !scanner.Scan() {
			break
		}

		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			detoxManager.ShowBadHabits()
		case "2":
			detoxManager.ShowGoodHabits()
		case "3":
			fmt.Print("Введите номер привычки для преодоления: ")
			scanner.Scan()
			if index, err := strconv.Atoi(scanner.Text()); err == nil {
				detoxManager.OvercomeAddiction(index - 1)
			}
		case "4":
			fmt.Print("Введите новую полезную привычку: ")
			scanner.Scan()
			detoxManager.AddGoodHabit(scanner.Text())
		case "5":
			fmt.Print("Сколько часов изучал Go? ")
			scanner.Scan()
			if hours, err := strconv.Atoi(scanner.Text()); err == nil {
				progressTracker.AddStudyTime(hours)
			}
		case "6":
			progressTracker.AddProject()
		case "7":
			progressTracker.ShowProgress()
			progressTracker.CalculateFuture()
		case "8":
			fmt.Printf("\n💫 МОТИВАЦИЯ: %s\n", progressTracker.GetRandomMotivation())
		case "9":
			fmt.Println("\n👋 До встречи, Гоша!")
			fmt.Println("💪 Помни: каждый день без игр - это день к успешной карьере!")
			return
		default:
			fmt.Println("❌ Неверный выбор. Попробуйте снова.")
		}

		// Показываем прогресс детокса
		progress := detoxManager.GetProgress()
		fmt.Printf("\n📈 Общий прогресс детокса: %.1f%%\n", progress)

		if progress > 50 {
			fmt.Println("🎉 Отличные результаты! Ты на полпути к трансформации!")
		}
	}
}
