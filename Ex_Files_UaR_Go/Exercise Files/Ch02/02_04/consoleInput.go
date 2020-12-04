package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//var s string
	//fmt.Scanln(&s)
	//fmt.Println(s)
	//unable to retrun the sepreate string

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text:")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)

	fmt.Print("Enter number:")
	str, _ = reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	//the space will disappeare
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("number", f)
	}

}
