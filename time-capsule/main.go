// time-capsule/main.go
// 25.10.25 - DAY 1: TIME CAPSULE INITIATION
// Первая программа 69-дневного челленджа - Цифровая капсула времени

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Capsule представляет цифровую капсулу времени
type Capsule struct {
	Metadata    CapsuleMetadata `json:"metadata"`
	Challenge   ChallengeInfo   `json:"challenge"`
	TechStack   TechEvolution   `json:"tech_stack"`
	PersonalMsg PersonalMessage `json:"personal_message"`
	Timestamp   time.Time       `json:"timestamp"`
}

type CapsuleMetadata struct {
	Version     string    `json:"version"`
	CapsuleID   string    `json:"capsule_id"`
	SealedUntil time.Time `json:"sealed_until"`
	CreatedAt   time.Time `json:"created_at"`
}

type ChallengeInfo struct {
	Name       string   `json:"name"`
	StartDate  string   `json:"start_date"`
	Duration   int      `json:"duration"`
	Phases     []string `json:"phases"`
	Target     int      `json:"target_programs"`
	LegacyGoal string   `json:"legacy_goal"`
}

type TechEvolution struct {
	PastTech    []string `json:"past_technologies"`
	CurrentTech []string `json:"current_technologies"`
	FutureTech  []string `json:"future_technologies"`
}

type PersonalMessage struct {
	ToFutureSelf string `json:"to_future_self"`
	Motivation   string `json:"motivation"`
	Promise      string `json:"promise"`
}

func main() {
	ShowBootScreen()
	DisplayTimeline()
	capsule := CreateTimeCapsule()
	SaveTimeCapsule(capsule)
	ReadAndDisplayCapsule()
	PlayRetroPacman()
	ShowCompletionMessage()
}

func ShowBootScreen() {
	fmt.Println("╔═══════════════════════════════════════╗")
	fmt.Println("║         TIME CAPSULE BOOTER          ║")
	fmt.Println("║         25.10.1999 - 25.10.2025      ║")
	fmt.Println("║                                       ║")
	fmt.Println("║  MS-DOS 6.22 → Go 1.25              ║")
	fmt.Println("║  Pacman → AI Neural Networks         ║")
	fmt.Println("║  25 years → 69 days transformation   ║")
	fmt.Println("║                                       ║")
	fmt.Println("║  INITIALIZING TIME CAPSULE...        ║")
	fmt.Println("╚═══════════════════════════════════════╝")

	// Анимация загрузки с тематикой времени
	loadingStages := []string{
		"Connecting to 1999...",
		"Syncing with 2025...",
		"Calibrating time bridge...",
		"Initializing legacy systems...",
		"Preparing future interfaces...",
	}

	for _, stage := range loadingStages {
		fmt.Printf("⏳ %s\r", stage)
		time.Sleep(400 * time.Millisecond)
	}
	fmt.Println("\n✅ Time capsule systems ready!")
}

func DisplayTimeline() {
	fmt.Println("\n📅 25-YEAR DIGITAL TIMELINE:")
	fmt.Println("┌─────────────────────────────────────────────────────┐")

	milestones := []struct {
		year  int
		event string
		icon  string
	}{
		{1999, "First PC with MS-DOS", "💾"},
		{2005, "First 'Hello World' in Pascal", "🐪"},
		{2010, "Web development era begins", "🌐"},
		{2015, "Mobile apps revolution", "📱"},
		{2020, "AI and Cloud dominance", "🤖"},
		{2025, "Dual Legacy 69 Challenge START", "🚀"},
	}

	for _, milestone := range milestones {
		fmt.Printf("│ %s %-4d │ %-35s │\n", milestone.icon, milestone.year, milestone.event)
	}

	fmt.Println("└─────────────────────────────────────────────────────┘")
}

func CreateTimeCapsule() Capsule {
	fmt.Println("\n🕰️  CREATING DIGITAL TIME CAPSULE...")

	// Создаем капсулу с информацией о челлендже
	capsule := Capsule{
		Metadata: CapsuleMetadata{
			Version:     "1.0",
			CapsuleID:   generateCapsuleID(),
			SealedUntil: time.Date(2026, 1, 2, 0, 0, 0, 0, time.UTC), // Дата завершения челленджа
			CreatedAt:   time.Now(),
		},
		Challenge: ChallengeInfo{
			Name:       "Dual Legacy 69 Challenge",
			StartDate:  "25.10.2025",
			Duration:   69,
			Phases:     []string{"Retro Resurrection", "Binary Evolution", "Future Legacy"},
			Target:     69,
			LegacyGoal: "Bridge between computing eras through code",
		},
		TechStack: TechEvolution{
			PastTech:    []string{"MS-DOS", "Pascal", "Assembly", "Windows 98", "Dial-up Modem"},
			CurrentTech: []string{"Go", "Docker", "Kubernetes", "AI/ML", "Cloud Computing"},
			FutureTech:  []string{"Quantum Computing", "AGI", "Neural Interfaces", "Digital Immortality"},
		},
		PersonalMsg: PersonalMessage{
			ToFutureSelf: "Remember the sound of the 56k modem and the magic of first 'Hello World'",
			Motivation:   "To preserve 25 years of computing history while building the future",
			Promise:      "I will complete all 69 days and create something meaningful for the next generation",
		},
		Timestamp: time.Now(),
	}

	fmt.Printf("🔮 Capsule ID: %s\n", capsule.Metadata.CapsuleID)
	fmt.Printf("📦 Sealed until: %s\n", capsule.Metadata.SealedUntil.Format("02.01.2006"))
	fmt.Println("✅ Time capsule created successfully!")

	return capsule
}

func SaveTimeCapsule(capsule Capsule) {
	fmt.Println("\n💾 SAVING TIME CAPSULE TO DIGITAL STORAGE...")

	// Создаем папку для капсул, если ее нет
	capsuleDir := "time-capsules"
	if err := os.MkdirAll(capsuleDir, 0755); err != nil {
		fmt.Printf("❌ Error creating capsule directory: %v\n", err)
		return
	}

	// Сохраняем капсулу в JSON файл
	filename := filepath.Join(capsuleDir, fmt.Sprintf("capsule-%s.json", capsule.Metadata.CapsuleID))

	data, err := json.MarshalIndent(capsule, "", "  ")
	if err != nil {
		fmt.Printf("❌ Error encoding capsule: %v\n", err)
		return
	}

	if err := os.WriteFile(filename, data, 0644); err != nil {
		fmt.Printf("❌ Error saving capsule: %v\n", err)
		return
	}

	fmt.Printf("✅ Capsule saved to: %s\n", filename)

	// Показываем прогресс сохранения
	for i := 0; i <= 100; i += 20 {
		fmt.Printf("📊 Preservation progress: %d%%\r", i)
		time.Sleep(150 * time.Millisecond)
	}
	fmt.Println("\n🎉 Digital preservation complete!")
}

func ReadAndDisplayCapsule() {
	fmt.Println("\n🔍 READING TIME CAPSULE CONTENTS...")

	// Имитация чтения капсулы
	time.Sleep(500 * time.Millisecond)

	fmt.Println("┌─────────────────────────────────────────────────────┐")
	fmt.Println("│                  CAPSULE CONTENTS                   │")
	fmt.Println("├─────────────────────────────────────────────────────┤")
	fmt.Println("│ 📅 Challenge: Dual Legacy 69                        │")
	fmt.Println("│ 🎯 Target: 69 programs in Go                        │")
	fmt.Println("│ ⏱️  Duration: 69 days of transformation            │")
	fmt.Println("│ 🌉 Mission: Bridge computing eras                   │")
	fmt.Println("│ 💫 Message: From Pacman to AI neural networks       │")
	fmt.Println("└─────────────────────────────────────────────────────┘")

	fmt.Println("\n📜 PERSONAL COMMITMENT:")
	fmt.Println("   \"I commit to this 69-day journey to honor")
	fmt.Println("    the past while building the digital future.\"")
}

func PlayRetroPacman() {
	fmt.Println("\n👻 MINI RETRO PACMAN TRIBUTE")
	fmt.Println("Loading simplified Go version...")

	// Улучшенная ASCII-анимация Pacman
	fmt.Println("┌─────────────────────┐")
	fmt.Println("│                     │")

	pacmanFrames := []string{
		"│      C>   ···       │",
		"│      C>  ···        │",
		"│      C> ···         │",
		"│      C>···          │",
		"│      C>··●          │",
		"│      C>·●           │",
		"│      C>●            │",
	}

	for i := 0; i < 3; i++ {
		for _, frame := range pacmanFrames {
			fmt.Printf("\r%s", frame)
			time.Sleep(200 * time.Millisecond)
		}
	}

	fmt.Println("\n│                     │")
	fmt.Println("└─────────────────────┘")
	fmt.Println("🎮 Retro tribute completed!")
}

func ShowCompletionMessage() {
	fmt.Println("\n" + strings.Repeat("═", 50))
	fmt.Println("⭐ DAY 1 COMPLETE: TIME CAPSULE ACTIVATED")
	fmt.Println(strings.Repeat("═", 50))

	message := `
🎯 WHAT WE ACCOMPLISHED TODAY:
• ✅ Booted legacy systems with modern Go
• ✅ Created digital time capsule for posterity
• ✅ Preserved challenge metadata and goals
• ✅ Established technical evolution timeline
• ✅ Paid tribute to computing heritage

🚀 TOMORROW'S MISSION:
Day 2: "MS-DOS Emulator in Go" - We begin the
       retro resurrection phase by recreating
       classic command line interfaces!

📅 PROGRESS: 1/69 days (1.4% complete)
⏳ NEXT MILESTONE: Day 25 - End of Retro Phase

💫 Remember: Every great journey begins with
   a single step. Today, we've sealed our
   commitment in digital stone!`

	fmt.Println(message)
	fmt.Println(strings.Repeat("═", 50))
}

// Вспомогательные функции
func generateCapsuleID() string {
	return fmt.Sprintf("DL69-%d-%s",
		time.Now().Year(),
		time.Now().Format("0201150405"))
}

func init() {
	fmt.Println(strings.Repeat("✨", 25))
	fmt.Println("   TIME CAPSULE MODULE INITIALIZED")
	fmt.Println("   69-DAY TRANSFORMATION BEGINS")
	fmt.Println(strings.Repeat("✨", 25))
	time.Sleep(1 * time.Second)
}
