// Что выведет программа
package main

import "fmt"

func main() {
	weekTmpArray := [7]int{1, 2, 3, 4, 5, 6, 7}
	worksDaySlice := weekTmpArray[0:5] // [1, 2, 3, 4, 5] len 5, cap 7
	weekendSlice := weekTmpArray[5:]   // [6, 7] len 2 cap 2
	weekTmpSlice := weekTmpArray[:]    // [1, 2, 3, 4, 5, 6, 7] len 7 cap 7
	weekTmpSlice[0] = 0                // [0, 2, 3, 4, 5, 6, 7]
	// worksDaySlice = [0, 2, 3, 4, 5]

	fmt.Println(worksDaySlice, len(worksDaySlice), cap(worksDaySlice)) // [0, 2, 3, 4, 5]  len 5, cap 7
	fmt.Println(weekendSlice, len(weekendSlice), cap(weekendSlice))    // [6, 7] len 2 cap 2
	fmt.Println(weekTmpSlice, len(weekTmpSlice), cap(weekTmpSlice))    // [0, 2, 3, 4, 5, 6, 7] len 7 cap 7
}
