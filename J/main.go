package main

import "fmt"

func main() {
	a := make([]int, 0, 1)
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)

	abc(a)

	fmt.Println(a)
}

func abc(a []int) {
	a = append(a, 4)
	_ = a
}
