package main

import (
	"fmt"
	"time"
)

type EnergyAlchemist struct {
	Name          string
	Age           int
	Birthday      string
	CurrentStatus string
	NoFapStreak   int
	GoSkills      int
	EnergyLevel   int
	LifeVision    string
}

func main() {
	alchemist := &EnergyAlchemist{
		Name:          "Гоша",
		Age:           37,
		Birthday:      "30 ноября",
		CurrentStatus: "начало трансформации",
		NoFapStreak:   1,
		GoSkills:      20,
		EnergyLevel:   30,
		LifeVision:    "профессиональный Go разработчик к 38 годам",
	}

	alchemist.ShowIntro()
	alchemist.AnalyzeCurrentSituation()
	alchemist.StartEnergyTransmutation()
	alchemist.BuildNewIdentity()
	alchemist.CreateNovemberPlan()
	alchemist.ShowTransformationVision()
}

func (e *EnergyAlchemist) ShowIntro() {
	fmt.Println("⚡ АЛХИМИЯ ЭНЕРГИИ: ОТ ЗАВИСИМОСТЕЙ К GO ⚡")
	fmt.Println("=========================================")
	fmt.Printf("👨‍🔬 Алхимик: %s, %d лет\n", e.Name, e.Age)
	fmt.Printf("🎂 День рождения: %s (через 29 дней!)\n", e.Birthday)
	fmt.Printf("🚫 NoFap стрик: %d день\n", e.NoFapStreak)
	fmt.Printf("💻 Навыки Go: %d%%\n", e.GoSkills)
	fmt.Printf("⚡ Уровень энергии: %d%%\n", e.EnergyLevel)
	fmt.Println("\n💫 Миссия: Сублимация сексуальной энергии в код!")
	pressToContinue()
}

func (e *EnergyAlchemist) AnalyzeCurrentSituation() {
	fmt.Println("\n📊 АНАЛИЗ ТЕКУЩЕЙ СИТУАЦИИ:")
	fmt.Println("===========================")

	situation := []struct {
		aspect string
		status string
		impact string
	}{
		{"Семейное положение", "нет жены/детей", "свобода для интенсивного обучения"},
		{"Финансы", "нищета", "мотивация для изменений"},
		{"Время", "много свободного", "потенциал для роста"},
		{"Зависимости", "фильмы для взрослых", "источник энергии для трансформации"},
		{"Возраст", "37 → 38 лет", "последний рывок к успеху"},
	}

	for _, item := range situation {
		fmt.Printf("🔍 %s: %s\n", item.aspect, item.status)
		fmt.Printf("   💡 %s\n", item.impact)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("\n🎯 Ключевая insight:")
	fmt.Println("   Одиночество + свободное время + нереализованная энергия =")
	fmt.Println("   = МОЩНЫЙ ДВИГАТЕЛЬ ДЛЯ ТРАНСФОРМАЦИИ!")

	e.EnergyLevel = 50
	pressToContinue()
}

func (e *EnergyAlchemist) StartEnergyTransmutation() {
	fmt.Println("\n🔥 ЗАПУСК ТРАНСМУТАЦИИ ЭНЕРГИИ:")
	fmt.Println("==============================")

	transmutationProcess := []struct {
		step    string
		source  string
		target  string
		boost   int
	}{
		{"Перераспределение внимания", "фильмы для взрослых", "Go видеокурсы", 15},
		{"Сублимация либидо", "сексуальная энергия", "творческая энергия кода", 20},
		{"Трансформация времени", "пассивное потребление", "активное создание", 25},
		{"Перенаправление фокуса", "виртуальные фантазии", "реальные достижения", 20},
	}

	totalBoost := 0
	for _, step := range transmutationProcess {
		fmt.Printf("\n🔄 %s:\n", step.step)
		fmt.Printf("   📉 %s → 📈 %s\n", step.source, step.target)
		fmt.Printf("   ⚡ +%d%% к энергии\n", step.boost)

		e.EnergyLevel += step.boost
		totalBoost += step.boost
		e.GoSkills += step.boost / 2

		for i := 0; i < 3; i++ {
			fmt.Print("💫")
			time.Sleep(500 * time.Millisecond)
		}
		fmt.Println()
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("\n🎉 Суммарный прирост энергии: +%d%%\n", totalBoost)
	fmt.Printf("🚀 Текущий уровень энергии: %d%%\n", e.EnergyLevel)
	fmt.Printf("💻 Навыки Go: %d%%\n", e.GoSkills)

	e.NoFapStreak = 6
	pressToContinue()
}

func (e *EnergyAlchemist) BuildNewIdentity() {
	fmt.Println("\n🎭 СОЗДАНИЕ НОВОЙ ИДЕНТИЧНОСТИ:")
	fmt.Println("==============================")

	fmt.Println("🧬 Старая идентичность:")
	oldIdentity := []string{
		"🔴 Одинокий потребитель контента",
		"🔴 Зависимый от цифровых наркотиков",
		"🔴 Пассивный наблюдатель жизни",
		"🔴 Жертва обстоятельств",
	}

	for _, trait := range oldIdentity {
		fmt.Printf("   %s\n", trait)
		time.Sleep(700 * time.Millisecond)
	}

	fmt.Println("\n🟢 Новая идентичность:")
	newIdentity := []string{
		"🟢 Energy Alchemist - преобразователь энергии",
		"🟢 NoFap Warrior - воин чистоты сознания",
		"🟢 Go Developer - создатель цифрового будущего",
		"🟢 Architect of Destiny - архитектор своей судьбы",
	}

	for _, trait := range newIdentity {
		fmt.Printf("   %s\n", trait)
		time.Sleep(700 * time.Millisecond)
	}

	fmt.Println("\n💪 Сильные стороны новой идентичности:")
	strengths := []string{
		"• Дисквалинация 25 лет игр → Глубокая концентрация",
		"• Одиночество → Свобода для интенсивного обучения",
		"• Сексуальная энергия → Творческая сила кода",
		"• Возраст 37 лет → Зрелость и мудрость решений",
	}

	for _, strength := range strengths {
		fmt.Printf("   %s\n", strength)
		time.Sleep(800 * time.Millisecond)
	}

	e.CurrentStatus = "active transformation"
	pressToContinue()
}

func (e *EnergyAlchemist) CreateNovemberPlan() {
	fmt.Println("\n📅 ПЛАН НА НОЯБРЬ (НОЯБРЬ-НЕДРОЧАБРЬ):")
	fmt.Println("===================================")

	novemberWeeks := []struct {
		week    string
		focus   string
		targets []string
	}{
		{
			"Неделя 1 (1-7 ноября)",
			"Детокс и установка ритма",
			[]string{
				"Полный отказ от фильмов для взрослых",
				"Ежедневные занятия Go по 4+ часов",
				"Утренние ритуалы энергии (медитация, спорт)",
				"Первые 5 программ на Go",
			},
		},
		{
			"Неделя 2 (8-14 ноября)",
			"Глубокое погружение в Go",
			[]string{
				"Изучение concurrent programming",
				"Создание первого веб-сервиса",
				"Участие в Go комьюнити",
				"Навыки Go: 20% → 50%",
			},
		},
		{
			"Неделя 3 (15-21 ноября)",
			"Строительство проектов",
			[]string{
				"Портфолио из 3 проектов",
				"Изучение фреймворков (Gin, Echo)",
				"Подготовка резюме",
				"Навыки Go: 50% → 70%",
			},
		},
		{
			"Неделя 4 (22-30 ноября)",
			"Финальный рывок к дню рождения",
			[]string{
				"Собственный полноценный проект",
				"Подготовка к собеседованиям",
				"Навыки Go: 70% → 85%",
				"День рождения как точка нового старта!",
			},
		},
	}

	for _, week := range novemberWeeks {
		fmt.Printf("\n📅 %s\n", week.week)
		fmt.Printf("🎯 Фокус: %s\n", week.focus)
		for _, target := range week.targets {
			fmt.Printf("   ✅ %s\n", target)
			time.Sleep(500 * time.Millisecond)
		}
		time.Sleep(2 * time.Second)
	}

	fmt.Printf("\n🎂 Цель на день рождения (30 ноября):\n")
	fmt.Printf("   🚀 Навыки Go: 85%%+\n")
	fmt.Printf("   💼 Готовность к работе\n")
	fmt.Printf("   🔥 NoFap стрик: 30 дней!\n")
	fmt.Printf("   💪 Новая идентичность: утверждена!\n")

	pressToContinue()
}

func (e *EnergyAlchemist) ShowTransformationVision() {
	fmt.Println("\n🔮 ВИДЕНИЕ БУДУЩЕГО ПОСЛЕ ТРАНСФОРМАЦИИ:")
	fmt.Println("========================================")

	fmt.Println(`
	🌟┌─────────────────────────────────┐
	🌟│        ENERGY ALCHEMIST         │
	🌟│                                 │
	🌟│  Возраст:       37 → 38 лет     │
	🌟│  Статус:        одинок → востребован │
	🌟│  Навыки:        20%  → 85%      │
	🌟│  Энергия:       30%  → 95%      │
	🌟│  Доход:         нищета → 100к+  │
	🌟└─────────────────────────────────┘
	`)

	futureScenarios := []struct {
		aspect string
		before string
		after  string
	}{
		{"Социальный статус", "невидимый одиночка", "профессионал в IT комьюнити"},
		{"Финансы", "выживание", "создание накоплений и инвестиций"},
		{"Личная жизнь", "одиночество + фильмы", "реальные знакомства и отношения"},
		{"Самооценка", "неудачник", "специалист, создающий ценность"},
		{"Энергия", "рассеянная и подавленная", "сфокусированная и созидательная"},
	}

	for _, scenario := range futureScenarios {
		fmt.Printf("🎯 %s:\n", scenario.aspect)
		fmt.Printf("   📉 БЫЛО: %s\n", scenario.before)
		fmt.Printf("   📈 СТАЛО: %s\n", scenario.after)
		fmt.Println()
		time.Sleep(1 * time.Second)
	}

	fmt.Println("\n💑 ОТНОШЕНИЯ ПОСЛЕ ТРАНСФОРМАЦИИ:")
	relationshipInsights := []string{
		"💡 Привлекательность = Уверенность + Целеустремленность + Финансовая стабильность",
		"💡 Женщин привлекают мужчины, которые создают ценность и имеют видение",
		"💡 Go разработчик с зарплатой 150k+ = высокая привлекательность на рынке отношений",
		"💡 Энергия NoFap + профессиональные успехи = магнит для качественных партнерш",
		"💡 Вместо виртуальных фантазий - реальные отношения с реальными женщинами",
	}

	for _, insight := range relationshipInsights {
		fmt.Printf("   %s\n", insight)
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("\n🎉 Всего до дня рождения: 29 дней трансформации!\n")
	fmt.Println("🚀 От одиночества и зависимостей - к востребованности и реальным отношениям!")
}

func pressToContinue() {
	fmt.Print("\n↵ Нажми Enter чтобы продолжить...")
	fmt.Scanln()
}
