package main

import (
	"fmt"
	"time"
)

type GamingLegacy struct {
	TotalGamesPlayed int
	FavoriteGenres   []string
	Achievements     []GameAchievement
	Skills           map[string]int
}

type GameAchievement struct {
	Game      string
	Achievement string
	Date      string
	Impact    string // Как это повлияло на развитие
}

func AnalyzeGamingSkills() *GamingLegacy {
	legacy := &GamingLegacy{
		TotalGamesPlayed: 50, // Примерное количество
		FavoriteGenres:   []string{"RPG", "Simulation", "Open World", "Strategy"},
		Skills:           make(map[string]int),
	}

	// Заполняем achievements
	legacy.Achievements = []GameAchievement{
		{
			Game:       "The Sims 2",
			Achievement: "Построил успешную виртуальную семью",
			Date:       "2004",
			Impact:     "Развил стратегическое планирование и управление ресурсами",
		},
		{
			Game:       "GTA San Andreas",
			Achievement: "100% завершение игры",
			Date:       "2006",
			Impact:     "Улучшил persistence и целеустремленность",
		},
		{
			Game:       "Age of Empires II",
			Achievement: "Победил в хардкор режиме",
			Date:       "2008",
			Impact:     "Развил аналитическое мышление и многозадачность",
		},
	}

	// Инициализируем навыки
	legacy.Skills = map[string]int{
		"Стратегическое планирование": 85,
		"Быстрое принятие решений":    80,
		"Многозадачность":             75,
		"Анализ сложных систем":       70,
		"Креативное мышление":         65,
		"Управление ресурсами":        80,
		"Пространственное мышление":   75,
	}

	return legacy
}

func (g *GamingLegacy) DisplaySkillsAnalysis() {
	fmt.Println("\n🎮 АНАЛИЗ ИГРОВЫХ НАВЫКОВ:")
	fmt.Println("=========================")

	for skill, level := range g.Skills {
		visualLevel := ""
		for i := 0; i < level/10; i++ {
			visualLevel += "█"
		}
		for i := level/10; i < 10; i++ {
			visualLevel += "░"
		}
		fmt.Printf("   %-30s [%s] %d%%\n", skill, visualLevel, level)
		time.Sleep(200 * time.Millisecond)
	}
}

func (g *GamingLegacy) CalculateITPotential() {
	fmt.Println("\n🚀 ПОТЕНЦИАЛ ДЛЯ IT:")
	fmt.Println("===================")

	skillConversions := map[string]string{
		"Стратегическое планирование": "Архитектура ПО и планирование проектов",
		"Быстрое принятие решений":    "Оптимизация и дебаггинг",
		"Многозадачность":             "Concurrent programming",
		"Анализ сложных систем":       "System design и анализ требований",
		"Креативное мышление":         "UI/UX design и инновационные решения",
		"Управление ресурсами":        "Memory management и оптимизация",
		"Пространственное мышление":   "Data structures и алгоритмы",
	}

	fmt.Println("💫 Игровые навыки → IT компетенции:")
	for gameSkill, itSkill := range skillConversions {
		level := g.Skills[gameSkill]
		fmt.Printf("   🎮 %-28s → 💻 %s (уровень: %d%%)\n", gameSkill, itSkill, level)
		time.Sleep(300 * time.Millisecond)
	}
}

