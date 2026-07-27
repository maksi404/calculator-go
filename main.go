package main

import (
	"fmt"
)

func main() {
	var number1 float64
	var number2 float64
	var action string

	fmt.Println("___Калькулятор___")
	fmt.Print("Введите первое число: ")
	fmt.Scan(&number1)
	fmt.Print("Введите действие: ")
	fmt.Scan(&action)
	fmt.Print("Введите второе число:  ")
	fmt.Scan(&number2)
	IMT := calcurateIMT(number1, number2, action)
	outputResult(IMT)
}

func outputResult(imt float64) {
	result := fmt.Sprint("Ответ: ", imt)
	fmt.Print(result)
}

func calcurateIMT(number1 float64, number2 float64, action string) float64 {
	if action == "+" {
		IMT := number1 + number2
		return IMT
	}
	if action == "-" {
		IMT := number1 - number2
		return IMT
	}
	if action == "/" {
		if number2 == 0 {
			fmt.Println("Ошибка: делить на ноль нельзя!")
			return 0
		}
		IMT := number1 / number2
		return IMT
	}
	if action == "*" {
		IMT := number1 * number2
		return IMT
	}
	return 0
}
