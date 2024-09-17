package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Простой калькулятор")

	var num1Str, num2Str string
	var operation string

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&num1Str)
	num1, err := strconv.ParseFloat(num1Str, 64)
	if err != nil {
		fmt.Println("Ошибка: введите корректное число")
		os.Exit(1)
	}

	fmt.Print("Введите операцию (+, -, *, /): ")
	fmt.Scanln(&operation)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&num2Str)
	num2, err := strconv.ParseFloat(num2Str, 64)
	if err != nil {
		fmt.Println("Ошибка: введите корректное число")
		os.Exit(1)
	}

	var result float64
	switch operation {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Ошибка: деление на ноль невозможно")
			os.Exit(1)
		}
		result = num1 / num2
	default:
		fmt.Printf("Ошибка: неизвестная операция '%s'\n", operation)
		os.Exit(1)
	}

	fmt.Printf("Результат: %.2f\n", result)
}
