// Given sorted slices A and B. Slices keep unique values. Produce a union of them into sorted slice RES without duplicates.
// A = [1, 3, 5, 7]
// B = [1, 2, 3, 7, 9]
// RES = [1, 2, 3, 5, 7, 9]

package main

import "fmt"

func main() {
	a := []int{1, 3, 5, 7}
	b := []int{1, 2, 3, 7, 9, 10, 11}
	fmt.Println(merge(a, b))
}

func merge(a, b []int) []int {
	p1 := 0
	p2 := 0

	result := []int{}

	for p1 <= len(a)-1 && p2 <= len(b)-1 {
		if a[p1] < b[p2] {
			result = append(result, a[p1])
			p1++
		} else {
			if a[p1] == b[p2] {
				result = append(result, b[p2])
				p1++
				p2++
			} else {
				result = append(result, b[p2])
				p2++
			}
		}

		fmt.Println(p1, "...", p2)

	}

	for i := p1; i <= len(a)-1; i++ {
		result = append(result, a[i])
	}

	for i := p2; i <= len(b)-1; i++ {
		result = append(result, b[i])
	}

	return result
}
