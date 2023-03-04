package arabian

import (
	"KataCalculator/calculate"
	"fmt"
	"log"
)

func CalcArabian(a int, b int, operator string) {
	switch operator {
	case "+":
		fmt.Println(calculate.Sum(a, b))
	case "-":
		fmt.Println(calculate.Subtraction(a, b))
	case "*":
		fmt.Println(calculate.Multiplication(a, b))
	case "/":
		value, err := calculate.Division(a, b)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(value)
	}
}
