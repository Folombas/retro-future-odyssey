package main

import (
	"fmt"
	"time"
)

type TimeReclaimer struct {
	TotalWastedHours int
	DailyRecovery    int
	WeeklyRecovery   int
	RecoveryRate     float64
}

func NewTimeReclaimer() *TimeReclaimer {
	return &TimeReclaimer{
		TotalWastedHours: 25000, // 25 лет примерно
		DailyRecovery:    8,     // часов в день
		RecoveryRate:     0.85,  // 85% эффективность
	}
}

func (t *TimeReclaimer) CalculateTimeRecovery(days int) {
	fmt.Println("\n⏳ РАСЧЕТ ВОЗВРАЩЕНИЯ ВРЕМЕНИ:")
	fmt.Println("============================")

	weeklyHours := t.DailyRecovery * 7
	totalRecovered := weeklyHours * (days / 7)
	effectiveRecovery := float64(totalRecovered) * t.RecoveryRate

	fmt.Printf("📅 Период восстановления: %d дней\n", days)
	fmt.Printf("⏱️  Ежедневное возвращение: %d часов\n", t.DailyRecovery)
	fmt.Printf("📊 Еженедельное возвращение: %d часов\n", weeklyHours)
	fmt.Printf("⚡ Эффективность восстановления: %.0f%%\n", t.RecoveryRate*100)
	fmt.Printf("🎯 Всего возвращено: %.0f часов\n", effectiveRecovery)

	// Расчет до полного восстановления
	remainingWasted := float64(t.TotalWastedHours) - effectiveRecovery
	weeksToFullRecovery := remainingWasted / (float64(weeklyHours) * t.RecoveryRate)

	if weeksToFullRecovery > 0 {
		fmt.Printf("📈 До полного восстановления: ~%.0f недель\n", weeksToFullRecovery)
	} else {
		fmt.Printf("🎉 Полное восстановление достигнуто!\n")
	}
}

func (t *TimeReclaimer) ShowRecoverySchedule() {
	fmt.Println("\n🕒 РАСПИСАНИЕ ВОССТАНОВЛЕНИЯ:")
	fmt.Println("============================")

	schedule := []struct {
		time    string
		old     string
		new     string
		impact  string
	}{
		{"19:00-21:00", "Игры/Порно", "Изучение Go", "Основной прогресс"},
		{"21:00-23:00", "Соцсети/Ютуб", "Практика Go", "Закрепление навыков"},
		{"23:00-01:00", "Сериалы/Фильмы", "Проекты на Go", "Портфолио"},
		{"Уикенды", "Бессмысленный отдых", "Интенсивные воркшопы", "Прорывной рост"},
	}

	for _, slot := range schedule {
		fmt.Printf("   🕐 %s:\n", slot.time)
		fmt.Printf("      📉 БЫЛО: %s\n", slot.old)
		fmt.Printf("      📈 СТАЛО: %s\n", slot.new)
		fmt.Printf("      💪 Эффект: %s\n", slot.impact)
		time.Sleep(800 * time.Millisecond)
	}
}

func (t *TimeReclaimer) GetProductivityGains() map[string]int {
	return map[string]int{
		"Концентрация":     60,
		"Креативность":     45,
		"Скорость обучения": 75,
		"Качество кода":    55,
		"Мотивация":        80,
	}
}
