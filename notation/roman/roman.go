package roman

import (
	"KataCalculator/calculate"
	"errors"
	"fmt"
	"log"
)

type Roman struct {
	arabic int
	value  string
}

var ERRNOTNULLROMAN = "[Ошибка] В римской системе нет чисел меньше I"

func Division(a string, b string) (string, error) {
	result, err := calculate.Division(RomanToArabian(a), RomanToArabian(b))
	if err != nil {
		return "", err
	}
	if result <= 0 {
		return "", errors.New(ERRNOTNULLROMAN)
	}
	return ArabianToRoman(result), nil
}

func Sum(a string, b string) string {
	result := calculate.Sum(RomanToArabian(a), RomanToArabian(b))
	return ArabianToRoman(result)
}

func Multiplication(a string, b string) string {
	result := calculate.Multiplication(RomanToArabian(a), RomanToArabian(b))
	return ArabianToRoman(result)
}

func Subtraction(a string, b string) (string, error) {
	result := calculate.Subtraction(RomanToArabian(a), RomanToArabian(b))
	if result <= 0 {
		return "", errors.New(ERRNOTNULLROMAN)
	}
	return ArabianToRoman(result), nil
}

func IsRoman(s string) bool {

	switch s {
	case "I":
		return true
	case "II":
		return true
	case "III":
		return true
	case "IV":
		return true
	case "V":
		return true
	case "VI":
		return true
	case "VII":
		return true
	case "VIII":
		return true
	case "IX":
		return true
	case "X":
		return true
	default:
		return false
	}
}

func newRomans() []Roman {
	return []Roman{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
}

func ArabianToRoman(num int) string {
	roman := newRomans()
	result := ""
	for _, numeral := range roman {
		for num >= numeral.arabic {
			result += numeral.value
			num -= numeral.arabic
		}
	}
	return result
}

func RomanToArabian(romanNum string) int {
	romanNums := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
	}
	result := 0
	for i := 0; i < len(romanNum); i++ {
		if i > 0 && romanNums[romanNum[i]] > romanNums[romanNum[i-1]] {
			result += romanNums[romanNum[i]] - 2*romanNums[romanNum[i-1]]
		} else {
			result += romanNums[romanNum[i]]
		}
	}
	return result
}

func CalcRoman(a string, b string, operator string) {
	switch operator {
	case "+":
		fmt.Println(Sum(a, b))
	case "-":
		value, err := Subtraction(a, b)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(value)
	case "*":
		fmt.Println(Multiplication(a, b))
	case "/":
		value, err := Division(a, b)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(value)
	}
}
