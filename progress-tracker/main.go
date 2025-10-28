package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ProgressTracker struct {
	TotalDays    int
	CurrentDay   int
	Calculator   *ProgressCalculator
	Display      *RetroDisplay
}

func main() {
	tracker := &ProgressTracker{
		TotalDays:  69,
		Calculator: NewProgressCalculator(),
		Display:    NewRetroDisplay(),
	}

	tracker.Run()
}

func (p *ProgressTracker) Run() {
	p.Display.ShowWelcome()

	for {
		day := p.getUserInput()
		if day == 0 {
			break
		}

		if day < 1 || day > p.TotalDays {
			p.Display.ShowError("День должен быть от 1 до 69!")
			continue
		}

		p.CurrentDay = day
		progress := p.Calculator.CalculateProgress(day, p.TotalDays)
		p.Display.ShowProgress(progress, day, p.TotalDays)

		fmt.Print("\nПроверить другой день? (y/n): ")
		reader := bufio.NewReader(os.Stdin)
		response, _ := reader.ReadString('\n')
		if strings.TrimSpace(strings.ToLower(response)) != "y" {
			break
		}
	}

	p.Display.ShowFarewell()
}

func (p *ProgressTracker) getUserInput() int {
	fmt.Print("\n🎯 Введите номер дня челленджа (1-69) или 0 для выхода: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "0" {
		return 0
	}

	day, err := strconv.Atoi(input)
	if err != nil {
		p.Display.ShowError("Неверный формат! Введите число.")
		return -1
	}

	return day
}
