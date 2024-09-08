package main

import (
	 "fmt" 
	 "strconv"      
	 "strings" )    

func main() { var input string      
	fmt.Print("Введите свой пример выражения: ") 
	fmt.Scanln(&input)

result := calculate(input)
fmt.Println("Результат равен:", result)
}

func calculate(input string) int { 
	operators := "+-*/" 
	var operator string 
	for _, op := range operators { if strings.Contains(input, string(op)) { operator = string(op) 
		break } 
}

expression := strings.Split(input, operator)
num1, isRoman1 := convertToNumber(expression[0])
num2, isRoman2 := convertToNumber(expression[1])

if isRoman1 != isRoman2 {
	fmt.Println("Ошибка: нехватка мозгового модуля по преоброзованию")
	return 0
}

var result int
switch operator {
case "+":
	result = num1 + num2
case "-":
	result = num1 - num2
case "*":
	result = num1 * num2
case "/":
	result = num1 / num2
}

return result
}

func convertToNumber(str string) (int, bool) { romanToArabic := map[string]int{ "I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10, }

_, err := strconv.Atoi(str)
if err == nil {
	num, _ := strconv.Atoi(str)
	return num, false
}

num := 0
for len(str) > 0 {
	for key, value := range romanToArabic {
		if strings.HasPrefix(str, key) {
			num += value
			str = str[len(key):]
		}
	}
}

return num, true
}
