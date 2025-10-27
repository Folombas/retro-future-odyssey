package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type FileManager struct {
	currentDir string
}

func (fm *FileManager) ListFiles() {
	files, err := os.ReadDir(fm.currentDir)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	fmt.Println("Имя                   Тип      Размер")
	fmt.Println("──────────────────────────────────────")

	for _, file := range files {
		info, _ := file.Info()
		name := file.Name()
		if len(name) > 20 {
			name = name[:17] + "..."
		}

		var fileType string
		var size string

		if file.IsDir() {
			fileType = "Папка"
			size = "<DIR>"
		} else {
			fileType = "Файл"
			size = fmt.Sprintf("%d байт", info.Size())
		}

		fmt.Printf("%-20s %-8s %s\n", name, fileType, size)
	}
}

func (fm *FileManager) ChangeDirectory(path string) {
	if path == ".." {
		fm.currentDir = filepath.Dir(fm.currentDir)
		return
	}

	newPath := filepath.Join(fm.currentDir, path)

	info, err := os.Stat(newPath)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	if !info.IsDir() {
		fmt.Printf("Ошибка: %s не является папкой\n", path)
		return
	}

	fm.currentDir = newPath
	fmt.Printf("Переход в: %s\n", fm.currentDir)
}

func (fm *FileManager) ShowFileInfo(filename string) {
	filePath := filepath.Join(fm.currentDir, filename)
	info, err := os.Stat(filePath)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	fmt.Printf("Имя: %s\n", info.Name())
	fmt.Printf("Размер: %d байт\n", info.Size())
	fmt.Printf("Тип: ")
	if info.IsDir() {
		fmt.Printf("Папка\n")
	} else {
		fmt.Printf("Файл\n")
	}
	fmt.Printf("Изменен: %s\n", info.ModTime().Format("02.01.2006 15:04"))
}

func (fm *FileManager) CreateDirectory(dirname string) {
	dirPath := filepath.Join(fm.currentDir, dirname)
	err := os.Mkdir(dirPath, 0755)
	if err != nil {
		fmt.Printf("Ошибка создания папки: %v\n", err)
	} else {
		fmt.Printf("Папка '%s' создана\n", dirname)
	}
}

func (fm *FileManager) CreateFile(filename string) {
	filePath := filepath.Join(fm.currentDir, filename)
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Ошибка создания файла: %v\n", err)
		return
	}
	file.Close()
	fmt.Printf("Файл '%s' создан\n", filename)
}
