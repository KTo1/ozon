package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount, timeCount int
	var times string

	fmt.Fscan(in, &testCount)

	for i := 0; i < testCount; i++ {
		fmt.Fscan(in, &timeCount)
		for j := 0; j < timeCount; j++ {
			fmt.Fscan(in, &times)
			times = strings.Replace(times, "-", " ", 1)
			fmt.Fprintln(out, strings.Fields(times))
		}
	}
}
