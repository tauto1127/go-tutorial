package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var p *person = &person{name: "Taro", age: 20}

	fmt.Println(p)
	convertPersonName(p)
	fmt.Println(p)
}

func convertInt(num *int) {
	*num = 10
}

func convertPersonName(p *person) {
	p.name = "Jiro"
}
