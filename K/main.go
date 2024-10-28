package main

import "fmt"

func main() {
	var ar = [...]int{1, 2, 3, 4}
	sl := ar[1:3]

	fmt.Println(sl)
	ar[1] = 5
	fmt.Println(sl)
}
