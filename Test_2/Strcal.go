package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const maxLength = 40

func isString(str1 string) bool {
	return strings.HasPrefix(str1, "\"") && strings.HasSuffix(str1, "\"")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение:")

	// Считывание строки от пользователя
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	// Разделяем строку на части (аргументы)
	parts := parseInput(input)
	if len(parts) != 3 {
		fmt.Println("Ошибка: недостаточно аргументов или некорректный формат")
		return
	}

	firstStr := strings.Trim(parts[0], "\"")
	operator := parts[1]
	secondArg := strings.Trim(parts[2], "\"")

	if !isString(parts[0]) {
		fmt.Println("Ошибка: первый аргумент должен быть строкой в кавычках ")
		return
	} else if len(firstStr) > 10 {
		fmt.Println("Ошибка: строка не может быть длиннее 10 символов")
		return
	}

	var result string

	switch operator {
	case "+":

		if len(secondArg) > 10 {
			fmt.Println("Ошибка: строка не может быть длиннее 10 символов")
			return
		}
		result = firstStr + secondArg
	case "-":

		result = strings.Replace(firstStr, secondArg, " ", 1)
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

	fmt.Println(result)
}
func parseInput(input string) []string {
	parts := []string{}
	current := ""
	inQuotes := false

	for _, symbol := range input {
		switch symbol {
		case ' ':
			if !inQuotes {
				if current != "" {
					parts = append(parts, current)
					current = ""
				} else {
					current += string(symbol)
				}

			}
		case '"':
			inQuotes = !inQuotes
			current += string(symbol)
		default:
			current += string(symbol)
		}
	}
	if current != "" {
		parts = append(parts, current)
	}
	return parts
}
