package main

import "fmt"

func welcomeMessage() {
	fmt.Println("Welcome to Go programming!")
}

func getUserName() string {
	var name string
	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)

	return name
}

func getTwoNumbers() (int, int) {
	var num1 int
	var num2 int

	fmt.Println("Enter first number: ")
	fmt.Scanln(&num1)

	fmt.Println("Enter second number: ")
	fmt.Scanln(&num2)
	return num1, num2
}

func displayInput(name string, num1 int, num2 int) {
	fmt.Println("Hello", name)
	fmt.Println("First number is:", num1)
	fmt.Println("Second number is:", num2)
}

func add(num1 int, num2 int) {
	sum := num1 + num2
	fmt.Println("Sum is:", sum)
}

func diff(num1 int, num2 int) {
	diff := num1 - num2
	fmt.Println("Difference is:", diff)
}

func multiply(num1 int, num2 int) {
	multiply := num1 * num2
	fmt.Println("Product is:", multiply)
}

func quotient(num1 int, num2 int) {
	quotient := num1 / num2
	fmt.Println("Quotient is:", quotient)
}

func remainder(num1 int, num2 int) {
	remainder := num1 % num2
	fmt.Println("Remainder is:", remainder)
}

func goodbyeMessage() {
	fmt.Println("Goodbye!")
}

func main() {
	welcomeMessage()

	// get user input
	name := getUserName()

	// get two numbers
	num1, num2 := getTwoNumbers()

	displayInput(name, num1, num2)
	add(num1, num2)
	diff(num1, num2)
	multiply(num1, num2)
	quotient(num1, num2)
	remainder(num1, num2)
	goodbyeMessage()

}
