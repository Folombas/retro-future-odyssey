package main

import (
	"fmt"
	"time"
	"os"
	"os/exec"
	"runtime"
)

type DigitalPhoenix struct {
	Name          string
	Age           int
	PastRegrets   []string
	FutureGoals   []string
	CurrentStatus string
	Progress      int
}

func main() {
	phoenix := &DigitalPhoenix{
		Name:          "Гоша",
		Age:           37,
		CurrentStatus: "Начало трансформации",
		Progress:      0,
	}

	phoenix.ShowIntro()
	phoenix.ShowPastRegrets()
	phoenix.ShowTransformation()
	phoenix.SetupGoEnvironment()
	phoenix.ShowFutureVision()
	phoenix.MotivationalFinale()
}

func (d *DigitalPhoenix) ShowIntro() {
	clearScreen()
	fmt.Println("🔥 ЦИФРОВОЙ ФЕНИКС: ВОЗРОЖДЕНИЕ 🔥")
	fmt.Println("===================================")
	fmt.Printf("👋 Привет, я %s, мне %d лет\n", d.Name, d.Age)
	fmt.Println("После 25 лет в цифровом забвении...")
	fmt.Println("Пришло время возрождения!")
	pressToContinue()
}

func (d *DigitalPhoenix) ShowPastRegrets() {
	clearScreen()
	fmt.Println("📻 РЕТРО-ХРОНИКА ПОТЕРЯННЫХ ЛЕТ:")
	fmt.Println("=================================")

	regrets := []string{
		"🎮 2000-2005: Sims 2 - строил виртуальные жизни",
		"🚗 2005-2010: GTA Vice City - гонял по виртуальному городу",
		"🏙️  2010-2015: GTA San Andreas - жил чужой жизнью",
		"📺 2015-2020: Бесконечные сериалы и ютуб",
		"😔 2020-2025: Осознание потраченного времени...",
	}

	for i, regret := range regrets {
		fmt.Printf("%d. %s\n", i+1, regret)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("\n💡 Но сегодня всё меняется!")
	pressToContinue()
}

func (d *DigitalPhoenix) ShowTransformation() {
	clearScreen()
	fmt.Println("🔄 ПРОЦЕСС ТРАНСФОРМАЦИИ:")
	fmt.Println("=========================")

	transformations := []string{
		"🧠 Переключаю мышление: от потребителя к создателю",
		"💻 Меняю видеокарты на алгоритмы",
		"🎮 Превращаю игровую зависимость в код-зависимость",
		"📚 Заменяю геймплей на изучение Go",
		"🚀 Готовлюсь к взлёту цифрового феникса!",
	}

	for i, step := range transformations {
		animation := getLoadingAnimation(i)
		fmt.Printf("%s %s\n", animation, step)
		time.Sleep(1 * time.Second)
	}

	d.Progress = 25
	fmt.Printf("\n✅ Прогресс трансформации: %d%%\n", d.Progress)
	pressToContinue()
}

func (d *DigitalPhoenix) SetupGoEnvironment() {
	clearScreen()
	fmt.Println("⚙️  НАСТРОЙКА GO-ОКРУЖЕНИЯ:")
	fmt.Println("==========================")

	// Проверяем установку Go
	fmt.Println("🔍 Проверяем установку Go...")
	cmd := exec.Command("go", "version")
	output, err := cmd.Output()

	if err != nil {
		fmt.Println("❌ Go не установлен. Установите Go с https://golang.org")
	} else {
		fmt.Printf("✅ %s", string(output))
	}

	// Создаем структуру проекта
	fmt.Println("\n📁 Создаем структуру проекта...")
	createProjectStructure()

	// Показываем настройки
	fmt.Println("\n🎯 Настройки окружения:")
	showGoEnvironment()

	d.Progress = 50
	fmt.Printf("\n✅ Прогресс настройки: %d%%\n", d.Progress)
	pressToContinue()
}

func (d *DigitalPhoenix) ShowFutureVision() {
	clearScreen()
	fmt.Println("🔮 ВИДЕНИЕ БУДУЩЕГО:")
	fmt.Println("====================")

	visions := []string{
		"🌟 Через 69 дней: Свободное владение Go",
		"💼 Через 6 месяцев: Первая работа Go-разработчиком",
		"🚀 Через 1 год: Senior Go Developer",
		"🌍 Через 2 года: Создаю продукты, меняющие мир",
		"🎯 Через 5 лет: Технический директор собственной компании",
	}

	for i, vision := range visions {
		fmt.Printf("%d. %s\n", i+1, vision)
		time.Sleep(1 * time.Second)
	}

	d.Progress = 75
	pressToContinue()
}

func (d *DigitalPhoenix) MotivationalFinale() {
	clearScreen()
	fmt.Println("🎉 ЦИФРОВОЙ ФЕНИКС ВОССТАЛ!")
	fmt.Println("===========================")

	fmt.Println(`
	🔥┌─────────────────────┐🔥
	🔥│     ФЕНИКС         │🔥
	🔥│        \\\\//        │🔥
	🔥│         \\/         │🔥
	🔥│    /\\_/\\_/\\_/\\     │🔥
	🔥│   / GO LANG  \\     │🔥
	🔥│  /___________ \\    │🔥
	🔥└─────────────────────┘🔥
	`)

	fmt.Println("🎯 День 1 завершён! Ты сделал самый важный шаг!")
	fmt.Println("💪 Помни: каждый день кода приближает тебя к цели")
	fmt.Println("🚀 68 дней до полного преобразования!")

	d.Progress = 100
	fmt.Printf("\n🎊 ОБЩИЙ ПРОГРЕСС: %d%%\n", d.Progress)
	fmt.Println("\n📅 Следующая остановка: День 2 - MS-DOS эмулятор на Go!")
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

func getLoadingAnimation(step int) string {
	animations := []string{"🔄", "⚡", "💫", "🌟", "🚀"}
	return animations[step%len(animations)]
}

func createProjectStructure() {
	dirs := []string{
		"cmd/digital-phoenix",
		"internal/transformation",
		"pkg/legacy",
		"api/v1",
		"configs",
		"scripts",
		"docs",
	}

	for _, dir := range dirs {
		os.MkdirAll(dir, 0755)
		fmt.Printf("   📁 Создана папка: %s\n", dir)
		time.Sleep(200 * time.Millisecond)
	}
}

func showGoEnvironment() {
	envVars := []string{"GOPATH", "GOROOT", "GOOS", "GOARCH"}

	for _, env := range envVars {
		value := os.Getenv(env)
		if value == "" {
			value = "не установлено"
		}
		fmt.Printf("   %s: %s\n", env, value)
	}
}

