package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	modify(s)
	fmt.Println(s) // 2 2 2 3
}

func modify(s []int) {
	for i, n := range s {
		s[i] = n * 2
		if i%2 == 0 {
			s = append(s, i*2)
		}
	}
	fmt.Println("in func:", s) // 2 4 6 0 4
}
