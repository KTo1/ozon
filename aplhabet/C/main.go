package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//3
//6
//2 1 3 1 1 4
//2
//5 5
//8
//1 4 2 5 4 2 6 3

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscan(in, &testCount)

	for i := 0; i < testCount; i++ {
		var progsCount, progLevel int
		progsArray := make([][2]int, 0)

		fmt.Fscan(in, &progsCount)
		for j := 1; j <= progsCount; j++ {
			fmt.Fscan(in, &progLevel)
			progsArray = append(progsArray, [2]int{j, progLevel})
		}

		sortProgs(progsArray, out)
	}
}

// private

func sortProgs(progsArray [][2]int, out io.Writer) {
	if len(progsArray) <= 2 {
		fmt.Fprintln(out, progsArray[0][0], progsArray[1][0])
		fmt.Fprintln(out, "")
		return
	}

	firstLevel := progsArray[0][1]
	minLevelValue := progsArray[0][1]
	minLevelDiff := 1000000000
	for i := 1; i < len(progsArray); i++ {
		if abs(progsArray[i][1]-firstLevel) < minLevelDiff {
			minLevelDiff = abs(progsArray[i][1] - firstLevel)
			minLevelValue = progsArray[i][1]
		}
	}

	for i := 1; i < len(progsArray); i++ {
		if progsArray[i][1] == minLevelValue {
			minLevelDiff = i
			break
		}
	}

	fmt.Fprintln(out, progsArray[0][0], progsArray[minLevelDiff][0])

	progsArray = remove(progsArray, minLevelDiff)
	progsArray = remove(progsArray, 0)

	sortProgs(progsArray, out)
}

func abs(x int) int {
	if x < 0 {
		x = x * -1
	}

	return x
}

func remove(slice [][2]int, s int) [][2]int {
	return append(slice[:s], slice[s+1:]...)
}
