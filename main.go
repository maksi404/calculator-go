package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("___Калькулятор___")
	fmt.Println("Поддерживает ввод примера только через пробелы, н/п '1 + 10'")
	fmt.Print("Введите пример: ")

	reader := bufio.NewReader(os.Stdin)
	expression, _ := reader.ReadString('\n')
	expression = strings.TrimSpace(expression)

	parts := strings.Split(expression, " ")

	number1, _ := strconv.ParseFloat(parts[0], 64)
	action := parts[1]
	number2, _ := strconv.ParseFloat(parts[2], 64)

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
