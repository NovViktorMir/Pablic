package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input string
	fmt.Print("Вод : ")
	fmt.Scanln(&input)

	result := calculate(input)
	fmt.Println("Результат равен:", result) //          Немогу сформулировать логигу выполнения ответа того или инного типа. if  esle <=0,  Идёт работа над кейсом Ответ.

	num := result
	roman := decimalToRomanIterative(num)
	fmt.Printf("Римскими результат равен: %s\n", roman)

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
		fmt.Println("Паника, так как используются одновременно разные системы счисления.")
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

	num := 0           // возникновение проблем с интеграцией паники через convertToNumber.. что то делаю не так, было решенно пропустить ответ через
	for len(str) > 0 { //decimalToRomanIterative
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
		} //if num <0 {panic("Ответ не коректен")} Вновь непойму. Углубляюсь в Условные конструкции нехватает знаний
	}
	return result
}

