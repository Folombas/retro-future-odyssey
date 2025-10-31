package main

import "fmt"

type DigitalHabit struct {
	Name        string
	Type        string // "game", "movie", "series", "social"
	TimeWasted  int    // часов потрачено
	IsAddiction bool
}

type DetoxManager struct {
	badHabits    []DigitalHabit
	goodHabits   []string
	totalSaved   int
}

func NewDetoxManager() *DetoxManager {
	return &DetoxManager{
		badHabits: []DigitalHabit{
			{Name: "GTA серия", Type: "game", TimeWasted: 2000, IsAddiction: true},
			{Name: "The Sims", Type: "game", TimeWasted: 1500, IsAddiction: true},
			{Name: "Просмотр сериалов", Type: "series", TimeWasted: 3000, IsAddiction: true},
			{Name: "Соцсети", Type: "social", TimeWasted: 1000, IsAddiction: true},
		},
		goodHabits: []string{
			"Изучение Go",
			"Чтение документации",
			"Практика программирования",
			"Создание проектов",
		},
		totalSaved: 0,
	}
}

func (dm *DetoxManager) ShowBadHabits() {
	fmt.Println("🎮 ВРЕДНЫЕ ЦИФРОВЫЕ ПРИВЫЧКИ:")
	fmt.Println("─────────────────────────────")

	totalWasted := 0
	for i, habit := range dm.badHabits {
		status := "🔴 АКТИВНА"
		if !habit.IsAddiction {
			status = "🟢 ПРЕОДОЛЕНА"
		}
		fmt.Printf("%d. %-20s %4d часов %s\n", i+1, habit.Name, habit.TimeWasted, status)
		totalWasted += habit.TimeWasted
	}

	fmt.Printf("\n💀 Всего потрачено: %d часов (≈ %.1f лет)\n",
		totalWasted, float64(totalWasted)/24/365)
	fmt.Printf("🎯 Можно было изучить Go на продвинутом уровне %d раз!\n", totalWasted/500)
}

func (dm *DetoxManager) ShowGoodHabits() {
	fmt.Println("\n🐹 ПОЛЕЗНЫЕ ПРИВЫЧКИ GO-РАЗРАБОТЧИКА:")
	fmt.Println("─────────────────────────────────")

	for i, habit := range dm.goodHabits {
		fmt.Printf("%d. %s\n", i+1, habit)
	}

	fmt.Printf("\n✅ Сэкономлено времени: %d часов\n", dm.totalSaved)
	fmt.Printf("🚀 Это ≈ %d полноценных проектов на Go!\n", dm.totalSaved/50)
}

func (dm *DetoxManager) OvercomeAddiction(index int) {
	if index >= 0 && index < len(dm.badHabits) && dm.badHabits[index].IsAddiction {
		dm.badHabits[index].IsAddiction = false
		savedTime := dm.badHabits[index].TimeWasted / 10 // символическое сохранение
		dm.totalSaved += savedTime

		fmt.Printf("\n🎉 ПОЗДРАВЛЯЮ! Ты преодолел: %s\n", dm.badHabits[index].Name)
		fmt.Printf("💪 Сэкономлено: %d часов для изучения Go!\n", savedTime)
		fmt.Printf("📈 Твой прогресс: %d/%d привычек преодолено\n",
			dm.countOvercomeHabits(), len(dm.badHabits))
	} else {
		fmt.Println("❌ Неверный индекс или привычка уже преодолена")
	}
}

func (dm *DetoxManager) AddGoodHabit(habit string) {
	dm.goodHabits = append(dm.goodHabits, habit)
	fmt.Printf("✅ Добавлена полезная привычка: %s\n", habit)
}

func (dm *DetoxManager) countOvercomeHabits() int {
	count := 0
	for _, habit := range dm.badHabits {
		if !habit.IsAddiction {
			count++
		}
	}
	return count
}

func (dm *DetoxManager) GetProgress() float64 {
	return float64(dm.countOvercomeHabits()) / float64(len(dm.badHabits)) * 100
}
