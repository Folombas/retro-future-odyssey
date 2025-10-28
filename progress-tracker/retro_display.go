package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

type RetroDisplay struct {
	Width int
}

func NewRetroDisplay() *RetroDisplay {
	return &RetroDisplay{
		Width: 50,
	}
}

func (r *RetroDisplay) ShowWelcome() {
	r.clearScreen()
	fmt.Println("╔" + strings.Repeat("═", r.Width) + "╗")
	fmt.Println("║" + r.centerText("РЕТРО-ТРЕКЕР ПРОГРЕССА") + "║")
	fmt.Println("║" + r.centerText("Digital Transformation Odyssey") + "║")
	fmt.Println("║" + strings.Repeat(" ", r.Width) + "║")
	fmt.Println("║" + r.centerText("69-дневный челлендж") + "║")
	fmt.Println("║" + r.centerText("От MS-DOS к Go мастерству!") + "║")
	fmt.Println("╚" + strings.Repeat("═", r.Width) + "╝")

	time.Sleep(1 * time.Second)
	fmt.Println("\n🎮 Загрузка ретро-интерфейса...")
	time.Sleep(500 * time.Millisecond)

	// Анимация загрузки
	for i := 0; i <= 100; i += 10 {
		fmt.Printf("▓▓▓ %d%%\r", i)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("\n✅ Система готова!")
}

func (r *RetroDisplay) ShowProgress(data ProgressData, currentDay, totalDays int) {
	r.clearScreen()

	fmt.Println("╔" + strings.Repeat("═", r.Width) + "╗")
	fmt.Println("║" + r.centerText(fmt.Sprintf("ДЕНЬ %d ИЗ %d", currentDay, totalDays)) + "║")
	fmt.Println("╚" + strings.Repeat("═", r.Width) + "╝")

	fmt.Println("\n🎯 РЕЗУЛЬТАТЫ ТРАНСМУТАЦИИ:")
	fmt.Println(strings.Repeat("─", r.Width))

	// Основные метрики
	r.showMetric("✔ Эффективность трансмутации", data.Efficiency, "%")
	r.showMetric("✔ Преобразовано часов", float64(data.TransmutedHours), "")
	r.showMetric("✔ Получено IT-навыков", float64(data.SkillsAcquired), "")
	r.showMetric("✔ Уровень мастерства Go", data.GoMastery, "%")

	fmt.Println("\n📊 ПРОГРЕСС ЧЕЛЛЕНДЖА:")
	fmt.Println(strings.Repeat("─", r.Width))

	fmt.Printf("   День %d из %d\n", currentDay, totalDays)
	r.showProgressBar("Общий прогресс", data.OverallProgress)
	r.showProgressBar("Трансформация времени", data.TransformationRate)
	fmt.Printf("   Приобретено навыков: %d\n", data.SkillsAcquired)
	r.showProgressBar("Уровень зависимости", data.DependencyLevel)

	// Дополнительная информация
	fmt.Println("\n💡 СТАТУС:")
	fmt.Println(strings.Repeat("─", r.Width))

	daysRemaining := totalDays - currentDay
	fmt.Printf("   Осталось дней: %d\n", daysRemaining)
	fmt.Printf("   Текущая фаза: %s\n", r.getCurrentPhase(currentDay))
	fmt.Printf("   Следующий рубеж: %s\n", r.getNextMilestone(currentDay))

	fmt.Println("\n" + strings.Repeat("═", r.Width))
}

func (r *RetroDisplay) showMetric(name string, value float64, unit string) {
	valueStr := fmt.Sprintf("%.0f%s", value, unit)
	if value == float64(int(value)) {
		valueStr = fmt.Sprintf("%d%s", int(value), unit)
	}
	fmt.Printf("   %s: %s\n", name, valueStr)
}

func (r *RetroDisplay) showProgressBar(label string, percent float64) {
	barWidth := 20
	filled := int(math.Round(percent * float64(barWidth) / 100.0))

	bar := "["
	for i := 0; i < barWidth; i++ {
		if i < filled {
			bar += "█"
		} else {
			bar += "░"
		}
	}
	bar += fmt.Sprintf("] %.0f%%", percent)

	fmt.Printf("   %-25s %s\n", label+":", bar)
}

func (r *RetroDisplay) getCurrentPhase(day int) string {
	switch {
	case day <= 25:
		return "🏛️  Retro Resurrection"
	case day <= 50:
		return "⚡ Binary Evolution"
	case day <= 69:
		return "🚀 Future Legacy"
	default:
		return "Завершено"
	}
}

func (r *RetroDisplay) getNextMilestone(day int) string {
	milestones := map[int]string{
		5:  "День 5: Первый ретро-проект",
		10: "День 10: 10 программ на Go",
		25: "День 25: Конец Retro фазы",
		30: "День 30: Собственный язык",
		50: "День 50: Конец Binary фазы",
		69: "День 69: ФИНИШ!",
	}

	for milestoneDay := day + 1; milestoneDay <= 69; milestoneDay++ {
		if name, exists := milestones[milestoneDay]; exists {
			return name
		}
	}
	return "Финальный день!"
}

func (r *RetroDisplay) ShowError(message string) {
	fmt.Printf("\n❌ ОШИБКА: %s\n", message)
	time.Sleep(2 * time.Second)
}

func (r *RetroDisplay) ShowFarewell() {
	fmt.Println("\n" + strings.Repeat("═", r.Width))
	fmt.Println("🎉 Спасибо за использование трекера прогресса!")
	fmt.Println("🚀 Продолжайте трансформацию!")
	fmt.Println(strings.Repeat("═", r.Width))
}

func (r *RetroDisplay) centerText(text string) string {
	padding := (r.Width - len(text)) / 2
	if padding < 0 {
		padding = 0
	}
	return strings.Repeat(" ", padding) + text + strings.Repeat(" ", r.Width-len(text)-padding)
}

func (r *RetroDisplay) clearScreen() {
	// Простая очистка экрана через вывод новых строк
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
}

