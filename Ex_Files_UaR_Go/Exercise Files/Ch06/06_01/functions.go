package main

import "fmt"

func main() {
	doSomething()

	sum := addValues(23, 55)
	fmt.Println("sum:", sum)

	sum = addAllvalue(12, 54, 456789)
	fmt.Println("new sum:", sum)
}

func doSomething() {
	fmt.Println("Doing something")
}

func addValues(value1 int, value2 int) int {
	return value1 + value2
}

func addAllvalue(values ...int) int {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	fmt.Printf("%T\n", values)
	return sum
}
