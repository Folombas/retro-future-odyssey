package main

import (
	"fmt"
	"time"
)

type NeuroOptimizer struct {
	BrainState     string
	FocusLevel     int
	LearningSpeed  int
	Motivation     int
	NeuroType      string
}

func NewNeuroOptimizer() *NeuroOptimizer {
	return &NeuroOptimizer{
		BrainState:    "digital_overload",
		FocusLevel:    25,
		LearningSpeed: 30,
		Motivation:    40,
		NeuroType:     "СДВГ+ОКР",
	}
}

func (n *NeuroOptimizer) OptimizeForGo() {
	fmt.Println("\n🧠 ОПТИМИЗАЦИЯ МОЗГА ДЛЯ GO:")
	fmt.Println("==========================")

	optimizations := []struct {
		problem  string
		solution string
		result   int
	}{
		{"Рассеянное внимание", "Техника Pomodoro (25/5)", 35},
		{"Низкая мотивация", "Маленькие победы и награды", 25},
		{"Медленное обучение", "Структурированный подход", 20},
		{"Прокрастинация", "Четкие дедлайны и цели", 30},
	}

	for _, opt := range optimizations {
		fmt.Printf("\n⚡ Проблема: %s\n", opt.problem)
		fmt.Printf("   💡 Решение: %s\n", opt.solution)
		fmt.Printf("   📈 Прирост: +%d%%\n", opt.result)

		n.FocusLevel += opt.result / 4
		n.LearningSpeed += opt.result / 4
		n.Motivation += opt.result / 4

		time.Sleep(1 * time.Second)
	}

	n.BrainState = "optimized_for_go"
}

func (n *NeuroOptimizer) ShowBrainMetrics() {
	fmt.Println("\n📊 МЕТРИКИ ОПТИМИЗАЦИИ МОЗГА:")
	fmt.Println("===========================")

	metrics := map[string]int{
		"Уровень фокуса":     n.FocusLevel,
		"Скорость обучения":  n.LearningSpeed,
		"Мотивация":          n.Motivation,
		"Устойчивость":       65,
		"Креативность":       70,
	}

	for metric, value := range metrics {
		visual := ""
		for i := 0; i < value/10; i++ {
			visual += "█"
		}
		for i := value/10; i < 10; i++ {
			visual += "░"
		}
		fmt.Printf("   %-20s [%s] %d%%\n", metric, visual, value)
		time.Sleep(300 * time.Millisecond)
	}
}

func (n *NeuroOptimizer) GetNeuroSpecificTips() map[string][]string {
	return map[string][]string{
		"СДВГ": {
			"Короткие сессии кодинга (25 минут)",
			"Частые перерывы для движения",
			"Одна задача за раз",
			"Визуальный трекер прогресса",
		},
		"ОКР": {
			"Четкая структура проекта",
			"Стайл-гайд и линтеры",
			"Предсказуемый workflow",
			"Детальное планирование",
		},
		"Аутизм": {
			"Минималистичный интерфейс",
			"Логичная архитектура",
			"Четкая документация",
			"Предсказуемое поведение кода",
		},
	}
}

