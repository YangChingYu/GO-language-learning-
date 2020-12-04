package main

import "fmt"

func main() {
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "White"
	fmt.Println(colors)
	fmt.Println(colors[1])

	var numbers = [5]int{5, 4, 2, 1, 5}
	fmt.Println(numbers)

	fmt.Println("Number of colors:", len(colors))
	fmt.Println("Number of numbers:", len(numbers))
}
