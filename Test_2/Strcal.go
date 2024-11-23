package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const maxLength = 40

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение:")

	// Считывание строки от пользователя
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input) // Убираем лишние пробелы и переносы строк

	// Разделяем строку на части (аргументы)
	parts := strings.Fields(input)
	if len(parts) != 3 {
		fmt.Println("Ошибка: недостаточно аргументов или некорректный формат")
		return
	}

	firstStr := strings.Trim(parts[0], "\"")
	operator := parts[1]
	secondArg := parts[2]

	if len(firstStr) > 10 {
		fmt.Println("Ошибка: строка не может быть длиннее 10 символов")
		return
	}

	var result string

	switch operator {
	case "+":
		secondStr := strings.Trim(secondArg, "\"")
		if len(secondStr) > 10 {
			fmt.Println("Ошибка: строка не может быть длиннее 10 символов")
			return
		}
		result = firstStr + secondStr
	case "-":
		secondStr := strings.Trim(secondArg, "\"")
		result = strings.Replace(firstStr, secondStr, "", 1)
	case "*":
		number, err := strconv.Atoi(secondArg)
		if err != nil || number < 1 || number > 10 {
			fmt.Println("Ошибка: ожидается число от 1 до 10")
			return
		}
		result = strings.Repeat(firstStr, number)
	case "/":
		number, err := strconv.Atoi(secondArg)
		if err != nil || number < 1 || number > 10 {
			fmt.Println("Ошибка: ожидается число от 1 до 10")
			return
		}
		if len(firstStr)/number == 0 {
			fmt.Println("Ошибка: строка слишком короткая для деления")
			return
		}
		result = firstStr[:len(firstStr)/number]
	default:
		fmt.Println("Ошибка: неизвестная операция")
		return
	}

	// Обрезаем строку, если она длиннее 40 символов
	if len(result) > maxLength {
		result = result[:maxLength] + "..."
	}

	// Вывод результата
	fmt.Println(result)
}
