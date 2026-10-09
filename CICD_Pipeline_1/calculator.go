package main

import "fmt"

func main() {
	var a int
	var b int
	var pemdas_choice int

	fmt.Println("Please select option:")
	fmt.Println("1 - Addition")
	fmt.Println("2 - Subtract")
	fmt.Println("3 - Divide")
	fmt.Println("4 - Multiply")
	fmt.Print("Enter Choice : ")
	fmt.Scan(&pemdas_choice)
	fmt.Print("Enter first number : ")
	fmt.Scan(&a)
	fmt.Print("Enter second number : ")
	fmt.Scan(&b)
	if pemdas_choice == 1 {
    	result := Addition(a, b)
    	println(a, "added to", b, "=", result)
	}else if pemdas_choice == 2{
		result := Subtract(a, b)
    	println(a, "subtracted by", b, "=", result)
	}else if pemdas_choice == 3{
		result := Divide(a, b)
    	println(a, "divided by", b, "=", result)
	}else if pemdas_choice == 4{
		result := Multiply(a, b)
    	println(a, "multiplied by", b, "=", result)
	}
	
}

func Addition(a int, b int) int {
	return a + b
}

func Subtract(a int, b int) int {
	return a - b
}

func Divide(a int, b int) int {
	if b == 0 {
		print("you cannot divide by 0")
	} else {
		return a / b
	}
	return a / b
}

func Multiply(a int, b int) int {
	return a * b
}
