package main

import "fmt"

type person struct {
	name string
	age  int
}

func main2() {
	fmt.Println("Hello World")
	p1 := person{
		name: "John",
		age:  30,
	}
	fmt.Println(p1)
	fmt.Println(p1.name)
}
