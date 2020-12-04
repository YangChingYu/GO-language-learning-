package main

import "fmt"

type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

func (d Dog) speak() {
	fmt.Println(d.Sound)
}

func (d Dog) S3() {
	d.Sound = fmt.Sprintf("%v! %v! %v!\n", d.Sound, d.Sound, d.Sound)
	fmt.Print(d.Sound)
}

func main() {
	poddle := Dog{"Poddle", 37, "Woof"}

	fmt.Println(poddle)
	poddle.speak()

	poddle.Sound = "Arf"
	poddle.speak()
	poddle.S3()

}
