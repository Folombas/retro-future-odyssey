package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func clearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func main() {
	clearScreen()

	fmt.Println("╔═══════════════════════════════════════╗")
	fmt.Println("║    RETRO TERMINAL COMMANDER v1.0     ║")
	fmt.Println("║           (c) 2025 RFO Team          ║")
	fmt.Println("║      Написано на Go • День 3         ║")
	fmt.Println("╚═══════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("Добро пожаловать в ретро-файловый менеджер!")
	fmt.Println("Стиль MS-DOS • Мощь Go • Будущее уже здесь!")
	fmt.Println()

	fm := &FileManager{}
	fm.currentDir, _ = os.Getwd()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("\n[%s]\n> ", fm.currentDir)

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		parts := strings.Fields(input)

		if len(parts) == 0 {
			continue
		}

		command := strings.ToLower(parts[0])
		args := parts[1:]

		switch command {
		case "dir", "ls":
			fm.ListFiles()
		case "cd":
			if len(args) > 0 {
				fm.ChangeDirectory(args[0])
			} else {
				fmt.Println("Использование: cd <папка>")
			}
		case "info":
			if len(args) > 0 {
				fm.ShowFileInfo(args[0])
			} else {
				fmt.Println("Использование: info <файл>")
			}
		case "mkdir":
			if len(args) > 0 {
				fm.CreateDirectory(args[0])
			} else {
				fmt.Println("Использование: mkdir <папка>")
			}
		case "touch":
			if len(args) > 0 {
				fm.CreateFile(args[0])
			} else {
				fmt.Println("Использование: touch <файл>")
			}
		case "exit", "quit":
			fmt.Println("До свидания! Возвращайтесь в будущее!")
			return
		case "help", "?":
			showHelp()
		default:
			fmt.Printf("Неизвестная команда: %s\nВведите 'help' для справки\n", command)
		}
	}
}

func showHelp() {
	fmt.Println(`
Доступные команды:
dir, ls    - список файлов
cd <папка> - сменить папку
info <файл>- информация о файле
mkdir <папка> - создать папку
touch <файл> - создать файл
help, ?    - справка
exit, quit - выход

Ретро-стиль • Go мощь • Будущее здесь!`)
}
