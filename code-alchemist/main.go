package main

import (
	"fmt"
	"time"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

type CodeAlchemist struct {
	Name           string
	Age            int
	GamingHistory  []GameEra
	AddictionLevel int
	Transmutation  *TransmutationEngine
	Progress       *TransformationProgress
}

type GameEra struct {
	Years    string
	Games    []string
	Hours    int
	Lessons  []string
}

func main() {
	alchemist := &CodeAlchemist{
		Name:           "Гоша",
		Age:            37,
		AddictionLevel: 85, // 0-100 scale
		Transmutation:  NewTransmutationEngine(),
		Progress:       NewProgressTracker(),
	}

	alchemist.InitGamingHistory()
	alchemist.RunTransmutationCeremony()
}

func (c *CodeAlchemist) InitGamingHistory() {
	c.GamingHistory = []GameEra{
		{
			Years: "2000-2005",
			Games: []string{"The Sims", "Sims 2", "Need for Speed", "Age of Empires"},
			Hours: 5000,
			Lessons: []string{
				"Стратегическое планирование (Sims)",
				"Управление ресурсами (Age of Empires)",
				"Тайм-менеджмент (Sims)",
				"Креативное мышление (строительство)",
			},
		},
		{
			Years: "2005-2010",
			Games: []string{"GTA Vice City", "GTA San Andreas", "The Sims 3", "Counter-Strike"},
			Hours: 8000,
			Lessons: []string{
				"Пространственное мышление (GTA карты)",
				"Быстрое принятие решений (Counter-Strike)",
				"Социальная динамика (GTA сюжеты)",
				"Многозадачность (Sims 3)",
			},
		},
		{
			Years: "2010-2020",
			Games: []string{"GTA V", "The Sims 4", "Cyberpunk 2077", "Assassin's Creed"},
			Hours: 12000,
			Lessons: []string{
				"Системное мышление (открытые миры)",
				"Анализ сложных систем (игровая экономика)",
				"Паттерны поведения (NPC AI)",
				"Архитектура больших проектов (игровые движки)",
			},
		},
	}
}

func (c *CodeAlchemist) RunTransmutationCeremony() {
	c.ShowAlchemistIntro()
	c.AnalyzeGamingLegacy()
	c.CalculateTotalInvestment()
	c.StartTransmutationProcess()
	c.ShowTransformationResults()
	c.CreateNewIdentity()
}

func (c *CodeAlchemist) ShowAlchemistIntro() {
	clearScreen()
	fmt.Println("🧪 АЛХИМИК КОДА: ТРАНСМУТАЦИЯ ЗАВИСИМОСТИ")
	fmt.Println("==========================================")
	fmt.Printf("👨‍🔬 Алхимик: %s, %d лет\n", c.Name, c.Age)
	fmt.Println("🎯 Миссия: Превратить 25000+ часов игр в мастерство Go")
	fmt.Println("🔥 Начинаем великое преобразование...")
	pressToContinue()
}

func (c *CodeAlchemist) AnalyzeGamingLegacy() {
	clearScreen()
	fmt.Println("📊 АНАЛИЗ ИГРОВОГО НАСЛЕДИЯ:")
	fmt.Println("============================")

	totalHours := 0
	for i, era := range c.GamingHistory {
		fmt.Printf("\n🎮 Эпоха %s:\n", era.Years)
		fmt.Printf("   Игры: %v\n", era.Games)
		fmt.Printf("   Часов: %d\n", era.Hours)
		fmt.Printf("   Извлеченные навыки:\n")
		for _, lesson := range era.Lessons {
			fmt.Printf("     💡 %s\n", lesson)
		}
		totalHours += era.Hours
		time.Sleep(2 * time.Second)
	}

	fmt.Printf("\n📈 ВСЕГО ИГРОВЫХ ЧАСОВ: %d\n", totalHours)
	fmt.Printf("⏰ Это %.1f лет непрерывного программирования!\n", float64(totalHours)/24/365)

	c.Progress.SetInitialMetrics(totalHours, len(c.GamingHistory))
	pressToContinue()
}

func (c *CodeAlchemist) CalculateTotalInvestment() {
	clearScreen()
	fmt.Println("💰 РАСЧЕТ ИНВЕСТИЦИЙ В ИГРЫ:")
	fmt.Println("============================")

	totalHours := c.Progress.TotalGamingHours
	hourlyRate := 500 // руб в час - примерная ставка разработчика

	potentialEarnings := totalHours * hourlyRate
	potentialSkills := totalHours / 100 // Предположим, 100 часов на навык

	fmt.Printf("🎮 Потрачено часов на игры: %d\n", totalHours)
	fmt.Printf("💵 Потенциальный заработок: %d руб\n", potentialEarnings)
	fmt.Printf("🚀 Потенциальных навыков: %d\n", potentialSkills)
	fmt.Printf("📚 Книг можно было прочитать: %d\n", totalHours/50)
	fmt.Printf("🏆 Проектов можно было завершить: %d\n", totalHours/200)

	fmt.Println("\n💡 Но это НЕ потерянное время!")
	fmt.Println("   Это инвестиции в твой уникальный опыт!")
	pressToContinue()
}

func (c *CodeAlchemist) StartTransmutationProcess() {
	clearScreen()
	fmt.Println("🔥 ЗАПУСК ПРОЦЕССА ТРАНСМУТАЦИИ:")
	fmt.Println("================================")

	// Процесс трансмутации навыков
	skillsMap := map[string]string{
		"Стратегическое планирование (Sims)": "→ Архитектура ПО",
		"Управление ресурсами (RTS)": "→ Memory Management в Go",
		"Пространственное мышление (GTA)": "→ Data Structures",
		"Быстрое принятие решений (FPS)": "→ Algorithm Optimization",
		"Социальная динамика (RPG)": "→ API Design",
		"Системное мышление (открытые миры)": "→ System Architecture",
		"Анализ сложных систем": "→ Debugging Complex Systems",
		"Паттерны поведения (NPC)": "→ Design Patterns",
		"Многозадачность (Sims)": "→ Concurrent Programming",
		"Архитектура больших проектов": "→ Microservices Design",
	}

	fmt.Println("🔄 Преобразование игровых навыков в IT-компетенции:")
	for gamingSkill, itSkill := range skillsMap {
		fmt.Printf("   %-45s %s\n", gamingSkill, itSkill)
		time.Sleep(500 * time.Millisecond)
	}

	// Запуск алхимического двигателя
	c.Transmutation.StartTransmutation(c.Progress.TotalGamingHours)
	pressToContinue()
}

func (c *CodeAlchemist) ShowTransformationResults() {
	clearScreen()
	fmt.Println("🎉 РЕЗУЛЬТАТЫ ТРАНСМУТАЦИИ:")
	fmt.Println("===========================")

	results := c.Transmutation.GetResults()

	fmt.Printf("🧪 Эффективность трансмутации: %d%%\n", results.Efficiency)
	fmt.Printf("🚀 Преобразованно часов: %d\n", results.TransmutedHours)
	fmt.Printf("💡 Получено IT-навыков: %d\n", results.AcquiredSkills)
	fmt.Printf("🎯 Уровень мастерства Go: %d%%\n", results.GoMasteryLevel)

	fmt.Println("\n📈 ПРОГРЕСС ЧЕЛЛЕНДЖА:")
	c.Progress.DisplayProgress(4, 69) // Day 4 of 69

	pressToContinue()
}

func (c *CodeAlchemist) CreateNewIdentity() {
	clearScreen()
	fmt.Println("🎭 СОЗДАНИЕ НОВОЙ ИДЕНТИЧНОСТИ:")
	fmt.Println("===============================")

	fmt.Println(`
	🧙‍♂️┌─────────────────────────────────┐
	🧙‍♂️│        КОД-АЛХИМИК              │
	🧙‍♂️│                                 │
	🧙‍♂️│  Игровая зависижность   →       │
	🧙‍♂️│       85%               →   15%  │
	🧙‍♂️│                                 │
	🧙‍♂️│  Мастерство Go         →       │
	🧙‍♂️│       15%               →   85%  │
	🧙‍♂️│                                 │
	🧙‍♂️│  Трансмутация завершена!        │
	🧙‍♂️└─────────────────────────────────┘
	`)

	fmt.Println("🎯 Новая идентичность: КОД-АЛХИМИК")
	fmt.Println("💪 Сильные стороны:")
	fmt.Println("   • Уникальный игровой опыт")
	fmt.Println("   • Глубокое понимание систем")
	fmt.Println("   • Стратегическое мышление")
	fmt.Println("   • Упорство (25 лет игр!)")

	fmt.Println("\n🚀 План развития:")
	fmt.Println("   День 5-30: Освоение баз Go + ретро-проекты")
	fmt.Println("   День 31-50: Backend разработка + фреймворки")
	fmt.Println("   День 51-69: Собственный проект + портфолио")

	c.AddictionLevel = 15 // Резкое снижение зависимости
	fmt.Printf("\n✅ Уровень игровой зависимости: %d%%\n", c.AddictionLevel)
	fmt.Printf("🎮 Преобразовано в код: %d часов\n", c.Progress.TransformedHours)
}

// Вспомогательные функции
func clearScreen() {
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func pressToContinue() {
	fmt.Println("\n↵ Нажми Enter чтобы продолжить...")
	fmt.Scanln()
}

