package calculate

import "errors"

func Sum(a int, b int) int {
	return a + b
}
func Division(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("деление на 0")
	} else {
		return a / b, nil
	}
}
func Multiplication(a int, b int) int {
	return a * b
}
func Subtraction(a int, b int) int {
	return a - b
}
