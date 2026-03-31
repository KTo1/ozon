package main

import (
	"bufio"
	"fmt"
	"os"
)

//1
//12
//2 2 2 2 2 2 2 3 3 3 3 3

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscan(in, &testCount)

	for i := 0; i < testCount; i++ {
		var goodsCount, price int
		priceMap := make(map[int]int)
		fmt.Fscan(in, &goodsCount)
		for j := 0; j < goodsCount; j++ {
			fmt.Fscan(in, &price)
			if priceMap[price] == 0 {
				priceMap[price] = 0
			}
			priceMap[price] = priceMap[price] + 1
		}
		sum := 0
		for k, v := range priceMap {
			sum = sum + (v-v/3)*k
		}
		fmt.Fprintln(out, sum)
	}
}
