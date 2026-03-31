package main

import "fmt"

func main() {

	data := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	s := data[1:3]

	fmt.Println("s = ", s)
	fmt.Println("data = ", data)

	s[1] = 100
	s = append(s, 101)

	fmt.Println("s = ", s)
	fmt.Println("data = ", data)
}
