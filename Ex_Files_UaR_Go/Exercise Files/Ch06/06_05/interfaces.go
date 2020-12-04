package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "wooof"
}

func main() {
	poodle := Animal(Dog{})
	fmt.Println(poodle)
}
