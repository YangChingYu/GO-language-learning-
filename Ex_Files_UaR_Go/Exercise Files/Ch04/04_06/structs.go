package main

import (
	"fmt"
)

type Dog struct {
	Breed  string
	Weight int
}

func main() {
	poddle := Dog{"Poddle", 34}
	fmt.Print(poddle)
	fmt.Printf("%+v\n", poddle)
	fmt.Printf("Breed: %v\nweight: %v", poddle.Breed, poddle.Weight)

}
