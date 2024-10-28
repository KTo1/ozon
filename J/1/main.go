package main

import "fmt"

func main() {
	a := make([]int, 0, 3)
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)
	// b := append(a, 4)
	//
	// a[2] = 10

	abc(a)

	fmt.Println(a) // 1 2 3
	//	fmt.Println(b)
}

func abc(a []int) {
	a = append(a, 4)
	fmt.Println(a) // 1 2 3 4
	a[1] = 3
	// _ = a
}
