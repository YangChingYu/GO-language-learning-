package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "An implicity typed string"

	fmt.Println(strings.ToUpper(str1))
	fmt.Println(strings.Title(str1))

	lValue := "hello"
	uValue := "HELLO"
	fmt.Print("Equal Non-Case-Sensitive?", strings.EqualFold(lValue, uValue))

}
