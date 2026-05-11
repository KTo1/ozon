// Даны два отсортированных массив А и В, надо их слить в третий массив С
// A = [1, 3, 5, 7]
// B = [1, 2, 3, 7, 9]
// C = [1, 2, 3, 5, 7, 9]

package main

import "fmt"

func main() {
	a := []int{1, 3, 5, 7}
	b := []int{1, 2, 3, 7, 9}

	fmt.Println(merge(a, b))
}

func merge(a, b []int) []int {
	if len(a) == 0 {
		return b
	}

	if len(b) == 0 {
		return a
	}

	p1 := len(a) - 1
	p2 := len(b) - 1

	uniq := make(map[int]struct{})

	result := []int{}
	for p1 > 0 || p2 > 0 {
		if a[p1] > b[p2] {
			if _, ok := uniq[a[p1]]; !ok {
				result = append(result, a[p1])
				uniq[a[p1]] = struct{}{}
			}

			p1--
		} else {
			if _, ok := uniq[b[p2]]; !ok {
				result = append(result, b[p2])
				uniq[b[p2]] = struct{}{}
			}
			p2--
		}
	}

	if p1 >= 0 {
		for i := p1; i >= 0; i-- {
			if _, ok := uniq[a[i]]; !ok {
				result = append(result, a[i])
				uniq[a[i]] = struct{}{}
			}
		}
	}

	if p2 >= 0 {
		for i := p2; i >= 0; i-- {
			if _, ok := uniq[b[i]]; !ok {
				result = append(result, b[i])
				uniq[b[i]] = struct{}{}
			}
		}
	}

	return result
}
