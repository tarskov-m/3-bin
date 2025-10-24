package main

import (
	"fmt"
	"strings"

	"myapp/bins"
	"myapp/file"
	"myapp/storage"
)

func main() {
	// Демонстрация использования интерфейсов для DI
	// Создание экземпляров через интерфейсы
	var storageService = storage.NewJSONStorage()
	var fileReader = file.NewOSFileReader()
	var binService = bins.NewBinService()

	name, private := inputUser()

	// Использование BinService через интерфейс
	bin := binService.CreateBin(name, private)

	// Использование Storage через интерфейс
	err := storageService.SaveBin(bin)
	if err != nil {
		fmt.Printf("Ошибка при сохранении: %v\n", err)
		return
	}

	fmt.Printf("\nBin создан успешно!\n")
	fmt.Printf("ID: %s\n", bin.ID)
	fmt.Printf("Имя: %s\n", bin.Name)
	fmt.Printf("Приватный: %t\n", bin.Private)
	fmt.Printf("Создан: %s\n", bin.CreatedAt.Format("2006-01-02 15:04:05"))

	// Демонстрация использования FileReader через интерфейс
	if fileReader.IsJSONFile("bins.json") {
		fmt.Println("\nФайл bins.json является JSON файлом")
	}

	// Загрузка всех bins через интерфейс Storage
	allBins, err := storageService.LoadBins()
	if err != nil {
		fmt.Printf("Ошибка при загрузке bins: %v\n", err)
		return
	}

	fmt.Printf("\nВсего bins в хранилище: %d\n", len(allBins))

	// Демонстрация обратной совместимости со старым API
	fmt.Println("\n--- Демонстрация обратной совместимости ---")
	legacyBin := bins.NewBin("legacy-bin", false)
	fmt.Printf("Legacy bin создан: %s\n", legacyBin.Name)

	// Демонстрация гибкости интерфейсов
	fmt.Println("\n--- Демонстрация гибкости интерфейсов ---")
	fmt.Println("Все сервисы работают через интерфейсы:")
	fmt.Printf("- BinService: %T\n", binService)
	fmt.Printf("- Storage: %T\n", storageService)
	fmt.Printf("- FileReader: %T\n", fileReader)
	fmt.Println("Это позволяет легко заменять реализации!")
}

func inputUser() (string, bool) {
	var name string
	for {
		fmt.Print("Введите bin имя: ")
		fmt.Scanln(&name)
		name = strings.TrimSpace(name)
		if name == "" {
			fmt.Println("Имя не должно быть пустым")
			continue
		}
		break
	}
	var private bool
privateLoop:
	for {
		fmt.Print("Приватный bin? (y/n): ")
		var privateStr string
		fmt.Scanln(&privateStr)
		privateStr = strings.ToLower(strings.TrimSpace(privateStr))
		switch privateStr {
		case "y":
			private = true
			break privateLoop
		case "n":
			private = false
			break privateLoop
		default:
			fmt.Println("Неверный ввод")
			continue
		}
	}
	return name, private
}
