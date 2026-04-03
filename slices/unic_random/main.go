package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(uniqN(10))
}

// если чисел будет очень много, то может это все повиснуть потому, что
// все числа уже будут в мапе)
func uniqN(n int) []int {
	uniqmap := make(map[int]struct{})
	result := make([]int, 0, n)

	for {
		i := rand.Int()

		if _, ok := uniqmap[i]; !ok {
			result = append(result, i)
			uniqmap[i] = struct{}{}
		}

		if len(result) == n {
			break
		}
	}

	return result
}
