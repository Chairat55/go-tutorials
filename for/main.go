package main

import "fmt"

func main() {
	// Three-component loop
	sum := 0
	for i := 1; i < 5; i++ {
		sum += i
	}
	fmt.Println("Three-component loop --> ", sum) // 10 (1+2+3+4)


	// While loop
	n := 1
	for n < 5 {
		n *= 2
	}
	fmt.Println("While loop --> ", n) // 8 (1*2*2*2)


	// Infinite loop
	sumInfiniteLoop := 0
	for {
		sumInfiniteLoop++ // repeated forever
		if sumInfiniteLoop == 5 {
			break
		}
	}
	fmt.Println("sumInfiniteLoop --> ", sumInfiniteLoop)


	// For-each range loop
	strings := []string{"hello", "world"}
	for i, s := range strings {
		fmt.Println("For-each range loop --> ", i, s)
	}


	// Exit a loop
	sumExitALoop := 0
	for i := 1; i < 5; i++ {
		if i%2 != 0 { // skip odd numbers
			continue
		}
		sumExitALoop += i
	}
	fmt.Println("Exit a loop --> ", sumExitALoop) // 6 (2+4)
}