package main

import "fmt"

func main() {

	var x float64 = 0
	var result string

	if x < 0 {
		result = "Less thn zero"
	} else if x == 0 {
		result = "equal thn zero"
	} else {
		result = "Greater than or equal to zero"
	}

	fmt.Println("Result:", result)

}
