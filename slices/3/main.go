package main

import "fmt"

func main() {

	data := []int{10, 20, 30, 40, 50}
	s := data[1:3]

	fmt.Println("s = ", s)
	fmt.Println("data = ", data)
}
