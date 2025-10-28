package main

import "math"

type ProgressData struct {
	Day                 int
	TotalDays           int
	OverallProgress     float64
	TransformationRate  float64
	SkillsAcquired      int
	DependencyLevel     float64
	Efficiency          float64
	TransmutedHours     int
	GoMastery           float64
}

type ProgressCalculator struct {
	BaseTransformation float64
	BaseSkills         int
	BaseDependency     float64
}

func NewProgressCalculator() *ProgressCalculator {
	return &ProgressCalculator{
		BaseTransformation: 75.0,
		BaseSkills:         187,
		BaseDependency:     85.0,
	}
}

func (p *ProgressCalculator) CalculateProgress(currentDay, totalDays int) ProgressData {
	// Основной прогресс (линейный)
	overallProgress := float64(currentDay) / float64(totalDays) * 100

	// Трансформация игрового времени (логарифмический рост - быстрый старт)
	transformationRate := p.calculateTransformationRate(currentDay, totalDays)

	// Приобретенные навыки (квадратичный рост)
	skillsAcquired := p.calculateSkillsAcquired(currentDay, totalDays)

	// Уровень зависимости (экспоненциальное снижение)
	dependencyLevel := p.calculateDependencyLevel(currentDay, totalDays)

	// Эффективность трансмутации (асимптотический рост к 100%)
	efficiency := p.calculateEfficiency(currentDay, totalDays)

	// Преобразованные часы (основано на реальных данных)
	transmutedHours := p.calculateTransmutedHours(currentDay)

	// Мастерство Go (логистический рост)
	goMastery := p.calculateGoMastery(currentDay, totalDays)

	return ProgressData{
		Day:                currentDay,
		TotalDays:          totalDays,
		OverallProgress:    overallProgress,
		TransformationRate: transformationRate,
		SkillsAcquired:     skillsAcquired,
		DependencyLevel:    dependencyLevel,
		Efficiency:         efficiency,
		TransmutedHours:    transmutedHours,
		GoMastery:          goMastery,
	}
}

func (p *ProgressCalculator) calculateTransformationRate(day, totalDays int) float64 {
	// Быстрый рост в начале, затем замедление
	baseRate := p.BaseTransformation
	maxRate := 100.0
	progress := float64(day) / float64(totalDays)

	if day <= 4 {
		return baseRate
	}

	// Логистическая функция для плавного роста
	growth := 1.0 / (1.0 + math.Exp(-5.0*(progress-0.3)))
	return baseRate + (maxRate-baseRate)*growth
}

func (p *ProgressCalculator) calculateSkillsAcquired(day, totalDays int) int {
	baseSkills := p.BaseSkills
	maxSkills := 250
	progress := float64(day) / float64(totalDays)

	if day <= 4 {
		return baseSkills
	}

	// Квадратичный рост
	skills := float64(baseSkills) + (float64(maxSkills-baseSkills) * progress * progress)
	return int(skills)
}

func (p *ProgressCalculator) calculateDependencyLevel(day, totalDays int) float64 {
	baseDependency := p.BaseDependency
	minDependency := 0.0
	progress := float64(day) / float64(totalDays)

	if day <= 4 {
		return baseDependency
	}

	// Экспоненциальное снижение
	decay := math.Exp(-3.0 * progress)
	return baseDependency*decay + minDependency*(1-decay)
}

func (p *ProgressCalculator) calculateEfficiency(day, totalDays int) float64 {
	minEfficiency := 75.0
	maxEfficiency := 100.0
	progress := float64(day) / float64(totalDays)

	// Сигмоидальный рост
	efficiency := minEfficiency + (maxEfficiency-minEfficiency)/(1+math.Exp(-8.0*(progress-0.5)))

	if efficiency > maxEfficiency {
		return maxEfficiency
	}
	return efficiency
}

func (p *ProgressCalculator) calculateTransmutedHours(day int) int {
	baseHours := 25000
	// Линейная зависимость от дня
	return baseHours * day / 69
}

func (p *ProgressCalculator) calculateGoMastery(day, totalDays int) float64 {
	minMastery := 15.0
	maxMastery := 100.0
	progress := float64(day) / float64(totalDays)

	// Логистический рост с медленным началом и быстрым ростом в середине
	mastery := minMastery + (maxMastery-minMastery)/(1+math.Exp(-6.0*(progress-0.7)))

	if mastery > maxMastery {
		return maxMastery
	}
	return mastery
}
