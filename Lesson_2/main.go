package main

import "fmt"

func main() {
	fmt.Println("Math")

	fmt.Println("10 + 5 = ", add(10, 5))
	fmt.Println("10 - 5 = ", subtract(10, 5))
	fmt.Println("10 * 5 = ", multiply(10, 5))
	fmt.Println(" 10 / 5 = ", divide(10, 5))
	fmt.Println("Even and odd numbers")
	fmt.Println("The number is ", even(51))
	fmt.Println("Grades")
	fmt.Println("The grade is ", grade(78))
	num(10)
	oddprint(20)
	evenprint(10)
	fizz(100)
	Fib(10)
	back(0)
	fives(100)
	multiples(10)
	star(5)
	triangle(3)
	triangle1(5)
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}
func divide(a, b int) int {
	return a / b
}
