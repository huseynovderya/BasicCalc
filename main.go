package main

import (
	"fmt"
)

func main() {

	var a int
	var b int
	var operator string

	fmt.Println("Enter Your First Number")

	fmt.Scanln(&a)

	fmt.Println("Enter Your Second Number") ///// bura kimi ozum yazdim

	fmt.Scanln(&b)

	fmt.Println("Enter Operator")

	fmt.Scanln(&operator)

	switch operator { //bunu tapdim amma ardin ozum yazdim
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 { //burasin videodan baxdim
			fmt.Println("Divide by Zero situation")
		} else {
			fmt.Println(a / b)
		}
	default:
		fmt.Println("Invalid Operator")
	}

}
