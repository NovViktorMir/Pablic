package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	hello()

	var sempai string
	fmt.Print("Результат операции:  ")
	fmt.Scanln(&sempai)
	tip := sempai
	sow := tip

	A := "Вы выбрали систему счесления - Арабская:"
	R := "Вы выбрали систему счесления - Римская:"
	G := "Вы выбрали систему счесления - Греческая:"
	IA := "Вы выбрали систему счесления - Индо-арабская:"
	EI := "Вы выбрали систему счесления - Иврит:"
	ALL := "Вы выбрали все РЕЗУЛЬТАТЫ:"
	AVTO := "Автоматически определить результат по вводу:"
	switch sow {
	case "A", "R", "G", "IA", "EI", "ALL", "AVTO":
		fmt.Println("Команда введена корректно:")

	default:
		fmt.Println("Команда введена не корректно:Список команд представлен выше:")
		fmt.Println("Внимательней читайте условие выполнение программы!")
		panic("Mistake: you entered a conditionally incorrect expression : \n The session is over")

	}

	switch tip {
	case "A":
		if tip == "A" {
			fmt.Println(A)
			var input string
			fmt.Print("Ввод :")
			fmt.Scanln(&input)
			dool := len(input)

			if dool > 5 {
				err()
				defer panic("ERROR: The calculator only works with two variables :\n Attention it is allowed to enter two-digit numbers:\n The session is over")
			} else if dool < 2 {
				eer()
				defer panic("ERROR: The calculator only works with two variables : \n The session is over")
			}

			result := calculate(input)
			fmt.Println("Результат равен:", result)

		}
	case "R":
		if tip == "R" {
			fmt.Println(R)
			var input string
			fmt.Print("Ввод :")
			fmt.Scanln(&input)
			dool := len(input)
			if dool < 2 {
				eer()
			}
			result := calculate(input)
			num := result
			roman := decimalToRomanIterative(num)
			fmt.Printf("Римский результат равен: %s\n", roman)
			if result <= 0 {
				panic("Римской системе счисления не может быть Результат: 0 или отрицательное значение:\n Mistake: you entered a conditionally incorrect expression : \n The session is over")
			}

		}
	case "G":
		if tip == "G" {
			fmt.Println(G)
			var input string
			fmt.Print("Ввод :")
			fmt.Scanln(&input)
			dool := len(input)
			if dool < 2 {
				eer()
			}
			result := calculate(input)
			num := result
			gresh := decimalToGreshIterative(num)
			fmt.Printf("Греческий результат равен: %s\n", gresh)
			if result <= 0 {
				panic("Греческой системе счисления не может быть Результат: 0 или отрицательное значение:\n Mistake: you entered a conditionally incorrect expression : \n The session is over")
			}

		}
	case "IA":
		if tip == "IA" {
			fmt.Println(IA)
			var input string
			fmt.Print("Ввод :")
			fmt.Scanln(&input)
			dool := len(input)
			if dool < 2 {
				eer()
			}
			result := calculate(input)
			num := result
			Iarab := decimalToIarabIterative(num)
			fmt.Printf("Индо-арабский результат равен: %s\n", Iarab)
			if result <= 0 {
				panic("Индо-арабской системе счисления не может быть Результат: 0 или отрицательное значение:\n Mistake: you entered a conditionally incorrect expression : \n The session is over")

			}

		}
	case "EI":
		if tip == "EI" {
			fmt.Println(EI)
			var input string
			fmt.Print("Ввод :")
			fmt.Scanln(&input)
			dool := len(input)
			if dool < 2 {
				eer()
			}
			result := calculate(input)
			num := result
			eivr := decimalToEivrIterative(num)
			fmt.Printf("Результат на Иврите равен: %s\n", eivr)
			if result <= 0 {
				panic("Еврейской системе счисления не может быть Результат: 0 или отрицательное значение:\n Mistake: you entered a conditionally incorrect expression : \n The session is over")

			}

		}

	case "AVTO":

		if tip == "AVTO" {
			fmt.Println(AVTO)
			for i := 0; i < 10; i++ {
				if i == 9 {
					defer fmt.Println("Goodbye: we hope you enjoyed it!")
				}

				var input string
				fmt.Print("Ввод :")
				fmt.Scanln(&input)
				dool := len(input)
				ven := input[0]
				sol := ven
				switch sol {
				case 48, 49, 50, 51, 52, 53, 54, 55, 56, 57:
					if dool > 5 {
						result := calculate(input)
						fmt.Println("Результат равен:", result)
						err()
						panic("ERROR: The calculator only works with two variables :\n Attention it is allowed to enter two-digit numbers: \n The session is over")
					} else if dool < 2 {
						eer()
						panic("ERROR: The calculator only works with two variables : \n The session is over")
					}
					result := calculate(input)
					fmt.Println("Результат равен:", result)
				case 73, 88, 86:
					result := calculate(input)
					num := result
					roman := decimalToRomanIterative(num)
					fmt.Printf("Результат равен: %s\n", roman)
					if result <= 0 {
						panic("Римской системе счисления не может быть Результат: 0 или отрицательное значение:\n Mistake: you entered a conditionally incorrect expression : \n The session is over")

					}

				}

			}
		}

	case "ALL":
		if tip == "ALL" {
			fmt.Println(ALL)
			var input string
			fmt.Print("Ввод :")
			fmt.Scanln(&input)
			dool := len(input)
			if dool < 1 {
				eer()
			}
			result := calculate(input)
			num := result
			roman := decimalToRomanIterative(num)
			gresh := decimalToGreshIterative(num)
			Iarab := decimalToIarabIterative(num)
			eivr := decimalToEivrIterative(num)
			fmt.Println("Арабский результат равен:", result)
			fmt.Printf("Римский результат равен: %s\n", roman)
			fmt.Printf("Греческий результат равен: %s\n", gresh)
			fmt.Printf("Индо-арабский результат равен: %s\n", Iarab)
			fmt.Printf("Еврейский результат равен: %s\n", eivr)

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

	nouoperators := "@#$^&=!№;:?`%|"
	var nouoperator string
	for _, ti := range nouoperators {
		if strings.Contains(input, string(ti)) {
			nouoperator = string(ti)
			break
		}
	}
	switch nouoperator {
	case "@", "#", "$", "^", "&", "=", "!", "№", ";", ":", "?", "`", "%", "|":
		panic("Вырожение веденно не верно: используйте операторы +|-|*|/ \n The expression is not executed...\n The session is over.")
	}

	expression := strings.Split(input, operator)
	num1, isRoman1 := convertToNumber(expression[0])
	num2, isRoman2 := convertToNumber(expression[1])

	if isRoman1 != isRoman2 {
		ups()
		panic("The expression is not executed...\n The session is over.")
	}

	var Равно int
	switch operator {
	case "+":
		Равно = num1 + num2
	case "-":
		Равно = num1 - num2
	case "*":
		Равно = num1 * num2
	case "/":

		if num2 != 0 {
			ZERO := num1 / num2
			return ZERO
		} else {
			eer()
			defer panic("На ноль делить нельзя! \n ERROR:Division by zero!: \n The session is over")

		}
		Равно = num1 / num2

	}

	return Равно

}

func convertToNumber(str string) (int, bool) {
	ARARIM := map[string]int{"0": 0, "I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
	GRESH := map[string]int{"0": 0, "α": 1, "β": 2, "γ": 3, "δ": 4, "ε": 5, "ϝ": 6, "ζ": 7, "η": 8, "θ": 9, "ι": 10}
	IARAB := map[string]int{"0": 0, "١": 1, "٢": 2, "٣": 3, "٤": 4, "٥": 5, "٦": 6, "٧": 7, "٨": 8, "٩": 9, "١٠": 10}
	EIVR := map[string]int{"0": 0, "א": 1, "ב": 2, "ג": 3, "ד": 4, "ה": 5, "ו": 6, "ז": 7, "ח": 8, "ט": 9, "י": 10}

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
	for len(str) > 0 {
		for key, value := range GRESH {
			if strings.HasPrefix(str, key) {
				num += value
				str = str[len(key):]

			}
		}
	}
	for len(str) > 0 {
		for key, value := range IARAB {
			if strings.HasPrefix(str, key) {
				num += value
				str = str[len(key):]

			}
		}
	}
	for len(str) > 0 {
		for key, value := range EIVR {
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
	{100, "C"}, {90, "XC"}, {80, "LXXX"}, {70, "LXX"}, {60, "LX"}, {50, "L"}, {40, "XL"},
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

var greshMap = []struct {
	decVal int
	symbol string
}{
	{100, "ρ"}, {90, "ϟ"}, {50, "ν"}, {40, "μ"}, {30, "λ٠"}, {20, "κ"},
	{10, "ι"}, {9, "θ"}, {5, "ε"}, {4, "δ"}, {1, "α"},
}

func decimalToGreshIterative(num int) string {
	result := ""
	for _, pair := range greshMap {
		for num >= pair.decVal {
			result += pair.symbol
			num -= pair.decVal
		}
	}
	return result
}

var IarabhMap = []struct {
	decVal int
	symbol string
}{
	{100, "١٠٠"}, {90, "٩٠"}, {50, "٥٠"}, {40, "٤٠"}, {30, "٣٠"}, {20, "٢٠"},
	{10, "١٠"}, {9, "٩"}, {5, "٥"}, {4, "٤"}, {1, "١"},
}

func decimalToIarabIterative(num int) string {
	result := ""
	for _, pair := range IarabhMap {
		for num >= pair.decVal {
			result += pair.symbol
			num -= pair.decVal
		}
	}
	return result
}

var EivrhMap = []struct {
	decVal int
	symbol string
}{
	{100, "ק"}, {90, "צ"}, {50, "נ"}, {40, "מ"}, {30, "ל"}, {20, "כ"},
	{10, "י"}, {9, "ט"}, {5, "ה"}, {4, "ד"}, {1, "א"},
}

func decimalToEivrIterative(num int) string {
	result := ""
	for _, pair := range EivrhMap {
		for num >= pair.decVal {
			result += pair.symbol
			num -= pair.decVal
		}
	}
	return result
}

func err() {
	fmt.Println("╔═══════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║.((╬X   X))                                                            ║")
	fmt.Println("║      A                                                                ║")
	fmt.Println("║                    ОТОБРАЖЕН КОРРЕКТНЫЙ РЕЗУЛЬТАТ                     ║")
	fmt.Println("╚═══════════════════════════════════════════════════════════════════════╝")
}

func eer() {
	fmt.Println("╔═══════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║.((╬X   X))                                                            ║")
	fmt.Println("║     ,-,                                                               ║")
	fmt.Println("║                    ВВЕЛИ НЕ КОРРЕКТНОЕ ВЫРАЖЕНИE                      ║")
	fmt.Println("╚═══════════════════════════════════════════════════════════════════════╝")
}

func ups() {
	fmt.Println("╔═══════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║.((╬0   0))                                                            ║")
	fmt.Println("║      ^                                                                ║")
	fmt.Println("║Паника! В выражении используются переменные из разных систем счисления.║")
	fmt.Println("╚═══════════════════════════════════════════════════════════════════════╝")

}

func hello() {
	fmt.Println("╔════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                                         ДОБРО ПОЖАЛОВАТЬ В КАЛЬКУЛЯТОР!                                                                   ( ^_^ )   version : 0.0041               ║")
	fmt.Println("║Данный калькулятор принимает символы из разных систем счисления:                                                                                                                    ║")
	fmt.Println("║Такие как: Арабские, Римские, Греческие, Индо-арабские, Иврит.                                                                                                                      ║")
	fmt.Println("║Вы можете использовать выражение в любой системе счисления и получить результат в другой системе счисления.                                                                         ║")
	fmt.Println("║Для этого введите команду в какой системе вы хотите получить результат:                                                                                                             ║")
	fmt.Println("║Результат:(введите одно из значений) A|R|G|IA|EI                                                                                                                                    ║")
	fmt.Println("║Также реализованы команды Результат: ALL||AVTO (вывести результат во всех представленных системах счисления или автоматически определить систему счисления по введенному выражению) ║")
	fmt.Println("║ВНИМАНИЕ: режим AVTO: version : 0.0041 : работает только с Римскими и Арабскими выражениями.                                                                                        ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝")
}

