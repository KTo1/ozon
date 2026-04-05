// Максимальная сумма подмассива длины K (фиксированное окно)
// [1, 2, 3, 4, 5] и k = 3
// max sum = 12
package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	k := 3

	fmt.Println(maxSumSubarray(arr, k))
}

func maxSumSubarray(arr []int, k int) int {
	firstSum := 0
	for i := 0; i < k; i++ {
		firstSum += arr[i]
	}

	maxSum := firstSum
	for i := k; i < len(arr); i++ {
		firstSum = firstSum + arr[i] - arr[i-k]
		if firstSum > maxSum {
			maxSum = firstSum
		}
	}

	return maxSum
}
