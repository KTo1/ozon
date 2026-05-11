// Задача: длина самой длинной подстроки без повторяющихся символов
// Дана строка s. Нужно найти длину самой длинной подстроки, в которой все символы различны.
//
// Примеры:
// "abcabcbb" → ответ 3 ("abc" или "bca" или "cab")
// "bbbbb" → ответ 1 ("b")
// "pwwkew" → ответ 3 ("wke" или "kew")
//
// Строка состоит из английских букв символов и пробелов
package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(str string) int {
	left := 0
	freq := map[byte]int{}
	maxLength := 0

	for right := 0; right < len(str); right++ {
		ch := str[right]
		freq[ch]++

		for freq[ch] > 1 {
			leftCh := str[left]
			freq[leftCh]--
			left++
		}

		if right-left+1 > maxLength {
			maxLength = right - left + 1
		}
	}

	return maxLength
}
