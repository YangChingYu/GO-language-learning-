package main

import "fmt"

func main() {
	n1, l1 := FullName("Zaphod", "Bebox")
	fmt.Printf("%v, %v", n1, l1)
}

//naming covention here do not need to add get like we do in java
func FullName(f, l string) (string, int) {
	full := f + " " + l
	length := len(full)
	return full, length
}
