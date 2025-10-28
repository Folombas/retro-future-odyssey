package main

import (
	"fmt"
	"time"
)

type TransmutationEngine struct {
	IsActive      bool
	StartTime     time.Time
	TotalHours    int
	Efficiency    int
}

type TransmutationResults struct {
	Efficiency       int
	TransmutedHours  int
	AcquiredSkills   int
	GoMasteryLevel   int
}

type TransformationProgress struct {
	TotalGamingHours int
	ErasAnalyzed     int
	TransformedHours int
	SkillsAcquired   int
	CurrentAddiction int
}

func NewTransmutationEngine() *TransmutationEngine {
	return &TransmutationEngine{
		IsActive:   false,
		Efficiency: 75, // базовая эффективность трансмутации
	}
}

func NewProgressTracker() *TransformationProgress {
	return &TransformationProgress{
		CurrentAddiction: 85,
	}
}

func (t *TransmutationEngine) StartTransmutation(totalHours int) {
	t.IsActive = true
	t.StartTime = time.Now()
	t.TotalHours = totalHours

	fmt.Println("🔥 АКТИВАЦИЯ АЛХИМИЧЕСКОГО ДВИГАТЕЛЯ...")

	// Имитация процесса трансмутации
	stages := []string{
		"Анализ игровых паттернов...",
		"Извлечение когнитивных навыков...",
		"Конвертация в алгоритмическое мышление...",
		"Оптимизация для Go...",
		"Синтез новых нейронных связей...",
	}

	for _, stage := range stages {
		fmt.Printf("   %s", stage)
		// Анимация загрузки
		for j := 0; j < 3; j++ {
			fmt.Printf(".")
			time.Sleep(300 * time.Millisecond)
		}
		fmt.Printf(" ✅\n")

		// Увеличиваем эффективность на каждом этапе
		t.Efficiency += 5
		time.Sleep(1 * time.Second)
	}

	fmt.Println("🎊 ТРАНСМУТАЦИЯ ЗАВЕРШЕНА!")
}

func (t *TransmutationEngine) GetResults() TransmutationResults {
	transmutedHours := t.TotalHours * t.Efficiency / 100
	acquiredSkills := transmutedHours / 100 // 100 часов на навык
	goMastery := acquiredSkills * 2         // упрощенная формула

	if goMastery > 100 {
		goMastery = 100
	}

	return TransmutationResults{
		Efficiency:      t.Efficiency,
		TransmutedHours: transmutedHours,
		AcquiredSkills:  acquiredSkills,
		GoMasteryLevel:  goMastery,
	}
}

func (p *TransformationProgress) SetInitialMetrics(totalHours int, eras int) {
	p.TotalGamingHours = totalHours
	p.ErasAnalyzed = eras
	p.TransformedHours = totalHours * 75 / 100 // начальная трансформация
	p.SkillsAcquired = p.TransformedHours / 100
	p.CurrentAddiction = 85
}

func (p *TransformationProgress) DisplayProgress(currentDay int, totalDays int) {
	progressPercent := (currentDay * 100) / totalDays
	transformationRate := (p.TransformedHours * 100) / p.TotalGamingHours

	fmt.Printf("   📅 День %d из %d\n", currentDay, totalDays)
	fmt.Printf("   📊 Прогресс челленджа: %d%%\n", progressPercent)
	fmt.Printf("   🔄 Трансформация игрового времени: %d%%\n", transformationRate)
	fmt.Printf("   💡 Приобретено навыков: %d\n", p.SkillsAcquired)
	fmt.Printf("   🎮 Текущая зависимость: %d%%\n", p.CurrentAddiction)

	// Визуальный прогресс-бар
	bar := "["
	for i := 0; i < 20; i++ {
		if i < progressPercent/5 {
			bar += "█"
		} else {
			bar += "░"
		}
	}
	bar += "]"
	fmt.Printf("   %s\n", bar)
}

func (p *TransformationProgress) UpdateTransformation(hoursTransformed int, skillsGained int) {
	p.TransformedHours += hoursTransformed
	p.SkillsAcquired += skillsGained
	p.CurrentAddiction -= 10 // уменьшаем зависимость

	if p.CurrentAddiction < 0 {
		p.CurrentAddiction = 0
	}
}
