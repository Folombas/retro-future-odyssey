package main

import (
	"fmt"
	"time"
)

type AddictionAnalyzer struct {
	Addictions     map[string]Addiction
	RecoveryStages map[int]string
}

type Addiction struct {
	Years       int
	Hours       int
	Impact      string
	RecoveryTime int
}

func NewAddictionAnalyzer() *AddictionAnalyzer {
	return &AddictionAnalyzer{
		Addictions: map[string]Addiction{
			"gaming": {
				Years:       25,
				Hours:       14000,
				Impact:      "Потеря амбиций, времени, социальных навыков",
				RecoveryTime: 12, // месяцев
			},
			"porn": {
				Years:       20,
				Hours:       6000,
				Impact:      "Снижение энергии, мотивации, уверенности",
				RecoveryTime: 6,
			},
			"social_media": {
				Years:       15,
				Hours:       3000,
				Impact:      "Клиповое мышление, снижение концентрации",
				RecoveryTime: 3,
			},
			"streaming": {
				Years:       13,
				Hours:       2000,
				Impact:      "Пассивное потребление, потеря времени",
				RecoveryTime: 2,
			},
		},
		RecoveryStages: map[int]string{
			1:  "Острое желание вернуться к старым привычкам",
			7:  "Первые признаки ясности мышления",
			14: "Появление новой энергии и мотивации",
			21: "Стабилизация нового образа жизни",
			30: "Полная трансформация идентичности",
		},
	}
}

func (a *AddictionAnalyzer) AnalyzeAddictions() {
	fmt.Println("\n🔍 ГЛУБОКИЙ АНАЛИЗ ЗАВИСИМОСТЕЙ:")
	fmt.Println("==============================")

	totalHours := 0
	for name, addiction := range a.Addictions {
		fmt.Printf("\n📊 %s:\n", getAddictionName(name))
		fmt.Printf("   📅 Лет зависимости: %d\n", addiction.Years)
		fmt.Printf("   ⏱️  Потрачено часов: %d\n", addiction.Hours)
		fmt.Printf("   💔 Воздействие: %s\n", addiction.Impact)
		fmt.Printf("   🏥 Время восстановления: %d месяцев\n", addiction.RecoveryTime)

		totalHours += addiction.Hours
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("\n💀 ОБЩАЯ СТАТИСТИКА:\n")
	fmt.Printf("   📈 Всего часов зависимости: %d\n", totalHours)
	fmt.Printf("   📅 Это %.1f лет непрерывной жизни!\n", float64(totalHours)/24/365)
}

func (a *AddictionAnalyzer) ShowRecoveryRoadmap() {
	fmt.Println("\n🛣️  ДОРОЖНАЯ КАРТА ВЫЗДОРОВЛЕНИЯ:")
	fmt.Println("================================")

	for day, stage := range a.RecoveryStages {
		fmt.Printf("   📍 День %d: %s\n", day, stage)
		time.Sleep(500 * time.Millisecond)
	}
}

func getAddictionName(key string) string {
	names := map[string]string{
		"gaming":       "🎮 Игровая зависимость",
		"porn":         "🔞 Порно-зависимость",
		"social_media": "📱 Зависимость от соцсетей",
		"streaming":    "🎬 Сериальная зависимость",
	}
	return names[key]
}

