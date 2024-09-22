package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var AR string
	fmt.Print("В какой системе исчисление вы хотите произвести операцию: Арабская или Риская (A|R): ")
	fmt.Scanln(&AR)

	tip := AR
	A := "Вы выбрали систему исчесления - Арабская:"
	R := "Вы выбрали систему исчесления - Римская:"

	if tip == "A" {
		fmt.Println(A)
		var input string
		fmt.Println("Вы можете записать вырожение как АРАБСКИМИ, так и РИМСКИМИ числами")
		fmt.Println("╔═══════════════════════════════════════════════════════════════════════╗\n║\t.( ͡ ಠ ʖ̯ ͡ಠ).\t\t\t\t\t\t\t   ║\n║\t\t☛\t\t\t\t\t\t\t║\n║РЕЗУЛЬТАТ БУДЕТ ОТОБРАЖЕН В ВЫБРАНОЙ СИСТИЕМЕ ИСЧИСЛЕНИЯ\t\t║")
		fmt.Println("╚═══════════════════════════════════════════════════════════════════════╝")
		fmt.Print("Ввод :")
		fmt.Scanln(&input)
		dool := len(input)
		if dool > 3 {
			fmt.Println("ВНИМАНИЕ Вы велли неверный формат вырожения, после ввода значений Число&Число :  остальное было удаленно ")
			fmt.Println("╔═══════════════════════════════════════════════════════════════════════╗\n║\t.((╬◣   ◢))\t\t\t\t\t\t\t║\n║\t      〇   ☛\t\t\t\t\t\t\t║\n║\t\t\tОТОБРАЖОН КОРРЕКТНЫЙ РЕЗУЛЬТАТ\t\t\t║")
			fmt.Println("╚═══════════════════════════════════════════════════════════════════════╝")
		}

		result := calculate(input)
		fmt.Println("Результат равен:", result)

	}

	if tip == "R" {
		fmt.Println(R)
		var input string
		fmt.Println("Вы можете записать вырожение как АРАБСКИМИ, так и РИМСКИМИ числами")
		fmt.Println("╔═══════════════════════════════════════════════════════════════════════╗\n║\t.( ͡ ಠ ʖ̯ ͡ಠ).\t\t\t\t\t\t\t   ║\n║\t\t☛\t\t\t\t\t\t\t║\n║РЕЗУЛЬТАТ БУДЕТ ОТОБРАЖЕН В ВЫБРАНОЙ СИСТИЕМЕ ИСЧИСЛЕНИЯ\t\t║")
		fmt.Println("╚═══════════════════════════════════════════════════════════════════════╝")
		fmt.Print("Ввод :")
		fmt.Scanln(&input)
		dool := len(input)
		if dool > 3 {
			fmt.Println("ВНИМАНИЕ Вы велли неверный формат вырожения, после ввода значений Число&Число :  остальное было удаленно ")
			fmt.Println("╔═══════════════════════════════════════════════════════════════════════╗\n║\t.((╬◣   ◢))\t\t\t\t\t\t\t║\n║\t      〇   ☛\t\t\t\t\t\t\t║\n║\t\t\tОТОБРАЖОН КОРРЕКТНЫЙ РЕЗУЛЬТАТ\t\t\t║")
			fmt.Println("╚═══════════════════════════════════════════════════════════════════════╝")
		}
		result := calculate(input)
		num := result
		roman := decimalToRomanIterative(num)
		fmt.Printf("Результат равен: %s\n", roman)
		if result <= 0 {
			fmt.Println("Паника, ответ в Римской системе исчисления не может быть равен 0 или отрицательному значению:")

		}

	}

}

func calculate(input string) int {
	operators := "+-*/"
	var operator string
	for _, op := range operators {
		if strings.Contains(input, string(op)) {
			operator = string(op)
			break
		}
	}

	expression := strings.Split(input, operator)
	num1, isRoman1 := convertToNumber(expression[0])
	num2, isRoman2 := convertToNumber(expression[1])

	if isRoman1 != isRoman2 {
		fmt.Println("╔════════════════════════════════════════════════════════════════════════╗\n║\t ┐(￣ヘ￣)┌\t\t\t\t\t\t\t ║\n║\t           \t\t\t\t\t\t\t ║\n║Паника! В выражении используются переменные из разных систем исчисления.║")
		fmt.Println("╚════════════════════════════════════════════════════════════════════════╝")
		return 0
	}

	var Ровно int
	switch operator {
	case "+":
		Ровно = num1 + num2
	case "-":
		Ровно = num1 - num2
	case "*":
		Ровно = num1 * num2
	case "/":
		Ровно = num1 / num2
	}

	return Ровно
}

func convertToNumber(str string) (int, bool) {
	ARARIM := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}

	_, err := strconv.Atoi(str)
	if err == nil {
		num, _ := strconv.Atoi(str)

		return num, false
	}

	num := 0
	for len(str) > 0 {
		for key, value := range ARARIM {
			if strings.HasPrefix(str, key) {
				num += value
				str = str[len(key):]

			}
		}
	}

	return num, true

}

var romanMap = []struct {
	decVal int
	symbol string
}{
	{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
	{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
	{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
}

func decimalToRomanIterative(num int) string {
	result := ""
	for _, pair := range romanMap {
		for num >= pair.decVal {
			result += pair.symbol
			num -= pair.decVal
		}
	}
	return result
}

