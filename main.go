package main

import (
	"KataCalculator/notation/arabian"
	"KataCalculator/notation/roman"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Operation struct {
	a        string
	b        string
	operator string
}

var ERRNOTMATH = "[ошибка] строка не является математической операцией."
var ERRNOTATION = "[ошибка] используются одновременно разные системы счисления."
var ERRFORMATSTR = "[ошибка] формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."
var ERRWRONGNUMBER = "[ошибка] неправильное число на входе. Доступные число от 1 до 10 включительно. \n Либо от I до X в римской системе"
var hellostr = `
______________________________________________________________________
Доступные числа: 1-10 
Доступные римские числа для ввода: I,II,III,IV,VI,VII,VIII,IX,X
exit - выход 
Введите выражение:
`

func main() {
	// использую сканер на случай если формат ввода у пользователя будет I + I, а не I+I
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(hellostr)
	for {
		scanner.Scan()
		str := scanner.Text()
		if str == "exit" {
			return
		}
		calc(str)
		fmt.Print("Введите выражение: ")
	}
}

// calc разбивает текущую строку на оператор и операнды и вызывает соответствующую функцию для их выполнения
func calc(str string) {
	// удаляем пробелы
	tempStr := strings.Split(str, " ")
	str = strings.Join(tempStr, "")
	operation := Operation{}
	flag := false
	for _, c := range str {
		if isOperator(string(c)) {
			if flag == true {
				log.Fatal(ERRNOTMATH)
			}
			flag = true
			operation.operator = string(c)
			continue
		}
		if !flag {
			operation.a += string(c)
		} else {
			operation.b += string(c)
		}
	}
	Validate(operation)
}

// Validate проверяет, соответствует ли входная строка формату математической операции и вызывает соответствующую функцию для ее вычисления.
func Validate(operation Operation) {
	if operation.a == "" || operation.b == "" || operation.operator == "" {
		log.Fatal(ERRNOTMATH)
	}

	if roman.IsRoman(operation.a) && roman.IsRoman(operation.b) {
		roman.CalcRoman(operation.a, operation.b, operation.operator)
		return
	}
	if arabA, err := strconv.Atoi(operation.a); err == nil {
		if arabB, err := strconv.Atoi(operation.b); err == nil {
			if arabA <= 10 && arabA >= 1 && arabB >= 1 && arabB <= 10 {
				arabian.CalcArabian(arabA, arabB, operation.operator)
			} else {
				log.Fatal(ERRWRONGNUMBER)
			}
			return
		}
	}

	if _, err := strconv.Atoi(operation.b); err == nil {
		if roman.IsRoman(operation.a) {
			log.Fatal(ERRNOTATION)
		}
	}
	if _, err := strconv.Atoi(operation.a); err == nil {
		if roman.IsRoman(operation.b) {
			log.Fatal(ERRNOTATION)
		}
	}

	if roman.RomanToArabian(operation.a) > 10 || roman.RomanToArabian(operation.b) > 10 {
		log.Fatal(ERRWRONGNUMBER)
	}
	log.Fatal(ERRFORMATSTR)
}

func isOperator(o string) bool {
	switch o {
	case "+":
		return true
	case "-":
		return true
	case "*":
		return true
	case "/":
		return true
	}
	return false
}
