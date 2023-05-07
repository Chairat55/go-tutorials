package main

import "fmt"

type First func(int) int
type Second func(int) First

func plus(a int, b int) int {

    return a + b
}

func plusPlus(a, b, c int) int {
    return a + b + c
}

func sum(x, y int) int {
	return x + y
}
func partialSum(x int) func(int) int {
	return func(y int) int {
		return sum(x, y)
	}
}

func squareSum(x int) func(int) func(int) int {
	return func(y int) func(int) int {
		return func(z int) int {
			return x*x + y*y + z*z
		}
	}
}

func squareSumDefineFunctionType(x int) Second {
	return func(y int) First {
		return func(z int) int {
			return x*x + y*y + z*z
		}
	}
}


func main() {

    res := plus(1, 2)
    fmt.Println("plus: 1+2 =", res)

    res = plusPlus(1, 2, 3)
    fmt.Println("plusPlus: 1+2+3 =", res)

	partial := partialSum(3)
	fmt.Println("partialSum:", partial(7))

	fmt.Println("squareSum:", squareSum(5)(6)(7))

	fmt.Println("squareSumDefineFunctionType:", squareSumDefineFunctionType(5)(6)(8))
}