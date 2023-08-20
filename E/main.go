package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int

	fmt.Fscan(in, &testCount)
	for i := 0; i < testCount; i++ {
		prev, cur, taskCount, yesOrNo := 0, 0, 0, true
		taskMap := make(map[int]int)

		fmt.Fscan(in, &taskCount)
		for j := 0; j < taskCount; j++ {
			fmt.Fscan(in, &cur)
			if cur != prev && prev != 0 {
				if taskMap[cur] != 0 {
					yesOrNo = false
					in.ReadString('\n')
					break
				}
			}
			taskMap[cur] = 1
			prev = cur
		}

		if yesOrNo {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
