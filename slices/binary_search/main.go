// реализовать бинарный поиск
package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{3, 6, 1, 7, 4, 11, 9, 14, 2}
	sort.Ints(arr)

	target := 1
	index := binarySearch3(arr, target)
	if index != -1 {
		fmt.Printf("element fount in index=%d\n", index)
	} else {
		fmt.Printf("element not found %d\n", target)
	}
}

// Что бы бинарный поиск работал, надо что бы масси был отсортирован
// идея такая, что надо делить массив пополам, и если мы находим что число больше чем то что нашли,
// то продолжаем поиск в правой половине иначе в левой и т.д. пока не найдем
// будем искать 14
// 1, 2, 3, 4, 6, 7, 9, 11, 14
// 0  1  2  3  4  5  6  7   8
// 8 / 2 = 4, первый элемент с индексом 4
func binarySearch(arr []int, target int) int {
	var leftIdx, rightIdx, midIdx int

	leftIdx = 0
	rightIdx = len(arr) - 1

	midIdx = (rightIdx - leftIdx) / 2

	if arr[midIdx] == target {
		return midIdx
	}

	for (rightIdx - leftIdx) > 0 {
		if arr[midIdx] > target {
			rightIdx = midIdx - 1
			midIdx = (rightIdx - leftIdx) / 2
		}

		if arr[midIdx] < target {
			leftIdx = midIdx + 1
			midIdx = leftIdx + (rightIdx-leftIdx)/2
		}

		if arr[midIdx] == target {
			return midIdx
		}

	}

	return -1
}

// Решение ИИ
func binarySearch3(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2 // защита от переполнения
		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
