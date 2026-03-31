package main

import "fmt"

type Person struct {
	name string
	asd  map[string]interface{}
}

func changeName(person **Person) {
	*person = &Person{name: "Ivan"}
}

func main() {
	person := &Person{name: "Vasya"}
	changeName(&person)
	fmt.Println(person.name)

	var a, b int
	a = 1
	switch {
	case a == 1:
		b := 1
		_ = b
	}

	if person.asd == nil {
		person.asd = make(map[string]interface{})
	}

	person.asd["asd"] = 123

	fmt.Println(b)
}
