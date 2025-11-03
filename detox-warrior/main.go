package main

import (
	"fmt"
	"time"
)

type DetoxWarrior struct {
	Name           string
	Age            int
	WastedYears    int
	RecoveredTime  int
	CurrentStreak  int
	GoProgress     int
	DigitalDemons  []string
	LifeVision     string
}

func main() {
	warrior := &DetoxWarrior{
		Name:          "Гоша",
		Age:           37,
		WastedYears:   25,
		RecoveredTime: 0,
		CurrentStreak: 7, // 7 дней челленджа
		GoProgress:    35,
		DigitalDemons: []string{
			"🎮 Игровая зависимость (1999-2025)",
			"📺 Порно-зависимость (2005-2025)",
			"📱 Социальные сети (2010-2025)",
			"🎬 Сериальная зависимость (2012-2025)",
		},
		LifeVision: "Профессиональный Go разработчик к 42 годам",
	}

	warrior.ShowIntro()
	warrior.ConfrontDigitalDemons()
	warrior.StartTimeReclamation()
	warrior.OptimizeBrainForGo()
	warrior.CreateRecoveryPlan()
	warrior.ShowTransformation()
}

func (d *DetoxWarrior) ShowIntro() {
	fmt.Println("🛡️  DIGITAL DETOX WARRIOR: ВОЗВРАЩЕНИЕ 25 ЛЕТ 🛡️")
	fmt.Println("==============================================")
	fmt.Printf("⚔️  Воин: %s, %d лет\n", d.Name, d.Age)
	fmt.Printf("💀 Потеряно лет: %d\n", d.WastedYears)
	fmt.Printf("🚫 Текущий стрик: %d дней\n", d.CurrentStreak)
	fmt.Printf("💻 Прогресс Go: %d%%\n", d.GoProgress)
	fmt.Println("\n🎯 Миссия: Вернуть 25 лет через код!")
	pressToContinue()
}

func (d *DetoxWarrior) ConfrontDigitalDemons() {
	fmt.Println("\n👹 КОНФРОНТАЦИЯ С ЦИФРОВЫМИ ДЕМОНАМИ:")
	fmt.Println("==================================")

	fmt.Println("📜 Хроника цифрового рабства:")
	for i, demon := range d.DigitalDemons {
		fmt.Printf("%d. %s\n", i+1, demon)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("\n💀 Самые разрушительные привычки:")
	destructiveHabits := []struct {
		habit     string
		hours     int
		impact    string
	}{
		{"GTA серии", 5000, "убийство амбиций"},
		{"The Sims", 4000, "виртуальная жизнь вместо реальной"},
		{"Порно-контент", 3000, "истощение энергии и мотивации"},
		{"Ютуб/ТикТок", 2000, "клиповое мышление"},
	}

	totalWastedHours := 0
	for _, habit := range destructiveHabits {
		fmt.Printf("   🔥 %s: %d часов → %s\n", habit.habit, habit.hours, habit.impact)
		totalWastedHours += habit.hours
		time.Sleep(800 * time.Millisecond)
	}

	fmt.Printf("\n📊 ВСЕГО ПОТРАЧЕНО: %d часов (%d лет!)\n",
		totalWastedHours, totalWastedHours/24/365)

	d.GoProgress = 45
	pressToContinue()
}

func (d *DetoxWarrior) StartTimeReclamation() {
	fmt.Println("\n⏰ ЗАПУСК ПРОЦЕССА ВОЗВРАЩЕНИЯ ВРЕМЕНИ:")
	fmt.Println("=====================================")

	reclamationProcess := []struct {
		action   string
		hours    int
		benefit  string
	}{
		{"Отказ от игр", 20, "+20 часов/неделю на Go"},
		{"Отказ от порно", 15, "+15 часов/неделю + энергия"},
		{"Отказ от соцсетей", 10, "+10 часов/неделю + фокус"},
		{"Отказ от сериалов", 12, "+12 часов/неделю + ясность"},
	}

	totalRecovered := 0
	for _, process := range reclamationProcess {
		fmt.Printf("\n🔄 %s\n", process.action)
		fmt.Printf("   ⏱️  +%d часов/неделю\n", process.hours)
		fmt.Printf("   💡 %s\n", process.benefit)

		d.RecoveredTime += process.hours
		totalRecovered += process.hours
		d.GoProgress += 5

		// Анимация восстановления
		for i := 0; i < 3; i++ {
			fmt.Print("✨")
			time.Sleep(500 * time.Millisecond)
		}
		fmt.Println()
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("\n🎉 Еженедельное возвращение времени: %d часов!\n", totalRecovered)
	fmt.Printf("🚀 Прогресс Go: %d%%\n", d.GoProgress)

	d.CurrentStreak = 7
	pressToContinue()
}

func (d *DetoxWarrior) OptimizeBrainForGo() {
	fmt.Println("\n🧠 ОПТИМИЗАЦИЯ МОЗГА ДЛЯ GO:")
	fmt.Println("==========================")

	fmt.Println("🔧 Ремонт нейронных путей:")
	brainRepairs := []struct {
		problem  string
		solution string
		result   string
	}{
		{"Клиповое мышление", "Глубокое изучение Go", "Умение решать сложные задачи"},
		{"Низкая концентрация", "Практика фокуса по Pomodoro", "2+ часа глубокой работы"},
		{"Дофаминовая зависимость", "Замена на кайф от кода", "Мотивация изнутри"},
		{"Прокрастинация", "Четкие дедлайны и цели", "Регулярный прогресс"},
	}

	for _, repair := range brainRepairs {
		fmt.Printf("\n⚡ Проблема: %s\n", repair.problem)
		fmt.Printf("   💊 Решение: %s\n", repair.solution)
		fmt.Printf("   🎯 Результат: %s\n", repair.result)
		time.Sleep(1 * time.Second)
	}

	// Особенности для нейроразнообразных
	fmt.Println("\n🌟 ОСОБЕННОСТИ ДЛЯ НЕЙРОРАЗНООБРАЗНЫХ:")
	neuroFeatures := []string{
		"СДВГ: Короткие сессии + частые перерывы = максимальная эффективность",
		"ОКР: Структурированный код + четкие правила = комфорт разработки",
		"Аутизм: Логичная архитектура + предсказуемость = идеальная среда",
	}

	for _, feature := range neuroFeatures {
		fmt.Printf("   %s\n", feature)
		time.Sleep(800 * time.Millisecond)
	}

	d.GoProgress = 65
	pressToContinue()
}

func (d *DetoxWarrior) CreateRecoveryPlan() {
	fmt.Println("\n📋 ПЛАН ВОССТАНОВЛЕНИЯ НА 62 ДНЯ:")
	fmt.Println("===============================")

	recoveryPhases := []struct {
		period  string
		focus   string
		actions []string
		target  string
	}{
		{
			"Дни 8-21: Детокс и база",
			"Очищение + основы Go",
			[]string{
				"Полный цифровой детокс",
				"Изучение синтаксиса Go",
				"Создание 10+ маленьких программ",
				"Ежедневные 4+ часа кодинга",
			},
			"Навыки Go: 65% → 80%",
		},
		{
			"Дни 22-42: Погружение",
			"Concurrency + Web",
			[]string{
				"Горутины и каналы",
				"Веб-фреймворки (Gin, Echo)",
				"Базы данных и API",
				"Первый полноценный проект",
			},
			"Навыки Go: 80% → 90%",
		},
		{
			"Дни 43-62: Профессионализация",
			"Архитектура + карьера",
			[]string{
				"Паттерны проектирования",
				"Микросервисная архитектура",
				"Подготовка портфолио",
				"Поиск работы",
			},
			"Трудоустройство!",
		},
	}

	for _, phase := range recoveryPhases {
		fmt.Printf("\n📅 %s\n", phase.period)
		fmt.Printf("🎯 Фокус: %s\n", phase.focus)
		for _, action := range phase.actions {
			fmt.Printf("   ✅ %s\n", action)
			time.Sleep(500 * time.Millisecond)
		}
		fmt.Printf("   🎯 Цель: %s\n", phase.target)
		time.Sleep(2 * time.Second)
	}

	fmt.Printf("\n💪 Всего до цели: 62 дня трансформации!\n")
	pressToContinue()
}

func (d *DetoxWarrior) ShowTransformation() {
	fmt.Println("\n🔮 ВИДЕНИЕ ПОСЛЕ ТРАНСФОРМАЦИИ:")
	fmt.Println("=============================")

	fmt.Println(`
	🦸┌─────────────────────────────────┐
	🦸│     DIGITAL DETOX WARRIOR       │
	🦸│                                 │
	🦸│  Потерянные годы: 25 → 0        │
	🦸│  Возвращено времени: 57+ часов  │
	🦸│  Навыки Go:      35% → 65%      │
	🦸│  Цифровые демоны: 4  → 0        │
	🦸└─────────────────────────────────┘
	`)

	transformations := []struct {
		aspect string
		before string
		after  string
	}{
		{"Время", "потраченное впустую", "инвестированное в будущее"},
		{"Энергия", "рассеянная по 1000 каналов", "сфокусированная на Go"},
		{"Мышление", "потребительское", "созидательное"},
		{"Самооценка", "неудачник-зависимый", "воин-созидатель"},
		{"Будущее", "безысходность и нищета", "профессия и перспективы"},
	}

	for _, transformation := range transformations {
		fmt.Printf("🎯 %s:\n", transformation.aspect)
		fmt.Printf("   📉 БЫЛО: %s\n", transformation.before)
		fmt.Printf("   📈 СТАЛО: %s\n", transformation.after)
		fmt.Println()
		time.Sleep(1 * time.Second)
	}

	fmt.Println("\n💫 КЛЮЧЕВОЙ ИНСАЙТ:")
	fmt.Println("   Каждый час, который я НЕ трачу на цифровые наркотики,")
	fmt.Println("   становится инвестицией в моё профессиональное будущее!")

	fmt.Printf("\n🎉 Неделя 1 завершена! 62 дня до полной трансформации!\n")
}

func pressToContinue() {
	fmt.Print("\n↵ Нажми Enter чтобы продолжить...")
	fmt.Scanln()
}

