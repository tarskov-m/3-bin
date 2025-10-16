package main

import (
	"fmt"
	"strings"

	"myapp/bins"
	"myapp/storage"
)

func main() {
	name, private := inputUser()

	bin := bins.NewBin(name, private)
	storage.SaveBin(bin)

	fmt.Printf("\nBin создан успешно!\n")
	fmt.Printf("ID: %s\n", bin.ID)
	fmt.Printf("Имя: %s\n", bin.Name)
	fmt.Printf("Приватный: %t\n", bin.Private)
	fmt.Printf("Создан: %s\n", bin.CreatedAt.Format("2006-01-02 15:04:05"))
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
