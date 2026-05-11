// Посчитать и вывести топ к самых часто встречающихся элементов
// k = 2
// [1, 1, 1, 2, 2, 3] = [1, 2]
// [2, 3, 1, 1, 2, 3, 3, 2] = [2, 3]
// [1] = [1]

// [_,_,_,_,_]

package main

import "fmt"

func main() {
	nums := []int{2, 3, 1, 1, 2, 3, 3, 2}
	k := 2
	fmt.Println(topKFrequent(nums, k))
}

func topKFrequent(nums []int, k int) []int {
	frec := make(map[int]int)
	for _, num := range nums {
		frec[num]++
	}

	tmp := make([][]int, len(nums)+1, len(nums)+1)
	for key, val := range frec {
		tmp[val] = append(tmp[val], key)
	}

	result := []int{}
	count := 0
	for i := len(tmp) - 1; i >= 0; i-- {
		if len(tmp[i]) == 0 {
			continue
		}

		for j := 0; j < len(tmp[i]); j++ {
			if count >= k {
				break
			}

			result = append(result, tmp[i][j])
			count++
		}
	}

	return result
}
