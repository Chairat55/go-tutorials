package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero = errors.New("divide by zero")

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

func DivideFmt(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("fmt.Errorf --> can't divide '%d' by zero", a)
	}
	return a / b, nil
}

func main() {
	a, b := 10, 0

	// errors
	result, err := Divide(a, b)

	if err != nil {
		switch {
		case errors.Is(err, ErrDivideByZero):
			fmt.Println("errors --> divide by zero error")
		default:
			fmt.Printf("errors --> unexpected division error: %s\n", err)
		}
	}

	fmt.Printf("errors --> %d / %d = %d\n", a, b, result)

	// fmt.Errorf
	resultFmt, errFmt := DivideFmt(a, b)
	if errFmt != nil {
		fmt.Println(errFmt)
	}

	fmt.Printf("fmt.Errorf --> %d / %d = %d\n", a, b, resultFmt)
}
