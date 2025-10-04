package main

import "fmt"

func getNumbers(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	diff := num1 - num2
	return sum, diff
}

func main() {
	sum, diff := getNumbers(5, 3)

	fmt.Println("Sum:", sum)
	fmt.Println("diff:", diff)

}
