package main

import (
	"fmt"
	"time"
)

type ProgressTracker struct {
	startDate    time.Time
	goStudyHours int
	projectsDone int
	motivations  []string
}

func NewProgressTracker() *ProgressTracker {
	return &ProgressTracker{
		startDate: time.Now(),
		motivations: []string{
			"Каждый час кода приближает тебя к тёплому офису!",
			"Игры дали тебе реакцию, Go даст тебе карьеру!",
			"25 лет потребления → 25 лет создания!",
			"Ты не терял время - ты накапливал опыт!",
			"С каждым днём ты ближе к работе мечты!",
		},
	}
}

func (pt *ProgressTracker) AddStudyTime(hours int) {
	pt.goStudyHours += hours
	fmt.Printf("✅ Записано %d часов изучения Go\n", hours)
	fmt.Printf("📊 Всего: %d часов\n", pt.goStudyHours)
}

func (pt *ProgressTracker) AddProject() {
	pt.projectsDone++
	fmt.Printf("🎉 Завершён проект №%d!\n", pt.projectsDone)
}

func (pt *ProgressTracker) ShowProgress() {
	daysSinceStart := int(time.Since(pt.startDate).Hours() / 24)

	fmt.Println("\n📊 ПРОГРЕСС ТРАНСФОРМАЦИИ:")
	fmt.Println("─────────────────────────")
	fmt.Printf("🕐 С начала трансформации: %d дней\n", daysSinceStart)
	fmt.Printf("🐹 Часов изучения Go: %d\n", pt.goStudyHours)
	fmt.Printf("🚀 Завершено проектов: %d\n", pt.projectsDone)
	fmt.Printf("📈 Среднее в день: %.1f часов\n", float64(pt.goStudyHours)/float64(daysSinceStart+1))

	if pt.goStudyHours > 0 {
		fmt.Printf("💡 Эффективность: 1 проект на %d часов изучения\n", pt.goStudyHours/pt.projectsDone)
	}
}

func (pt *ProgressTracker) GetRandomMotivation() string {
	if len(pt.motivations) == 0 {
		return "Ты на правильном пути!"
	}
	return pt.motivations[time.Now().Unix()%int64(len(pt.motivations))]
}

func (pt *ProgressTracker) CalculateFuture() {
	monthsToJob := 6 - (pt.goStudyHours / 50) // упрощенная формула
	if monthsToJob < 0 {
		monthsToJob = 0
	}

	fmt.Printf("\n🔮 ПРОГНОЗ:\n")
	fmt.Printf("   До первой работы Go-разработчиком: %d месяцев\n", monthsToJob)
	fmt.Printf("   При текущем темпе: до %s\n",
		time.Now().AddDate(0, monthsToJob, 0).Format("02.01.2006"))

	if monthsToJob == 0 {
		fmt.Println("   🎉 Ты готов к поиску работы!")
	}
}
