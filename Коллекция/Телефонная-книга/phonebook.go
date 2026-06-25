package main

import "fmt"

var contact = map[string]string{
	"Иванов":  "+7(999)999-99-99",
	"Петров":  "+7(888)888-88-88",
	"Сидоров": "+7(777)777-77-77",
}

func addcontact() {
	var userInputname string
	var userInputnumber string
	fmt.Println("Введите фамилию контакта!")
	fmt.Scanln(&userInputname)
	fmt.Println("Введите номер контакта!")
	fmt.Scanln(&userInputnumber)
	contact[userInputname] = userInputnumber
	fmt.Println("Контакт добавлен!")
}
func findcontact() {
	var userInputname string
	fmt.Println("Введите фамилию контакта!")
	fmt.Scanln(&userInputname)
	number, ok := contact[userInputname]
	if !ok {
		fmt.Println("Контакт не найден!")
	} else {
		fmt.Println("Номер контакта:", number)
	}
}

func showallcontacts() {
	fmt.Println("Вот все контакты: ")
	for name, number := range contact {
		fmt.Println("Фамилия:", name, "Номер:", number)
	}
}
func deletecontact() {
	var userInputname string
	fmt.Println("Введите фамилию контакта!")
	fmt.Scanln(&userInputname)
	_, ok := contact[userInputname]
	if !ok {
		fmt.Println("Контакт не найден!")
	} else {
		delete(contact, userInputname)
		fmt.Println("Контакт удален!")
	}

}

func main() {
	fmt.Println(`    1 — Добавить контакт
    2 — Найти контакт
    3 — Показать все контакты
    4 — Удалить контакт
    0 — Выйти`)
	for {
		var userInput int
		fmt.Scanln(&userInput)
		if userInput == 0 {
			fmt.Println("Выход из программы!")
			break
		}
		switch userInput {
		case 1:
			addcontact()
		case 2:
			findcontact()
		case 3:
			showallcontacts()
		case 4:
			deletecontact()
		default:
			fmt.Println("Неверный пункт меню!")
		}
	}
}
