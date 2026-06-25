package main

import (
	"fmt"
	"math/rand"
)

var a, b, number, userInput, i int = 1, 100, rand.Intn(100) + 1, 0, 0

func getuserInput() int {
	fmt.Println("Введите число от", a, "до", b)
	fmt.Scanln(&userInput)
	switch {
	case userInput < a || userInput > b:
		fmt.Println("Вы ввели число не из диапазона! Попробуйте еще раз!")
		getuserInput()
	default:
	}
	return userInput
}

func checkuserInput() bool {
	switch {
	case userInput == number:
		return true
	case userInput < number:
		fmt.Println("Загаданное число больше!")
		a = userInput + 1
	case userInput > number:
		fmt.Println("Загаданное число меньше!")
		b = userInput - 1
	}
	return false
}

func main() {
	fmt.Println("Добро пожаловать в игру 'Угадай число'! \nПрограмма загадала число от '1' до '100'! \nВам нужно угадать его за 10 попыток! \nНачинаем!")
	for i := 0; i < 10; i++ {
		getuserInput()
		if checkuserInput() {
			fmt.Println("Поздравляю! Вы угадали число с", i+1, "попытки!")
			break
		}
		fmt.Println("У вас осталось", 9-i, "попыток!")
	}

}
