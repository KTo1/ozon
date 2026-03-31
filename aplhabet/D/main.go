package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Matrix struct {
	col    int
	stroke []int
}

type MatrixSort []Matrix

func (m MatrixSort) Len() int {
	return len(m)
}

func (m MatrixSort) Less(i, j int) bool {
	return m[j].stroke[m[j].col] > m[i].stroke[m[i].col]
}

func (m MatrixSort) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount, clickCount, n, m, c int
	fmt.Fscan(in, &testCount)

	for k := 0; k < testCount; k++ {
		fmt.Fscan(in, "")
		fmt.Fscan(in, &n, &m)

		data := []Matrix{}

		for i := 0; i < n; i++ {
			stroke := make([]int, m)
			for j := 0; j < m; j++ {
				fmt.Fscan(in, &c)
				stroke[j] = c
			}
			data = append(data, Matrix{0, stroke})
		}

		fmt.Fscan(in, &clickCount)
		for i := 0; i < clickCount; i++ {
			fmt.Fscan(in, &c)
			sortMatrixByCol(data, c)
		}

		for _, v := range data {
			for i := 0; i < m; i++ {
				fmt.Fprint(out, v.stroke[i], " ")
			}
			fmt.Fprintln(out, "")
		}
		fmt.Fprintln(out, "")
	}
}

// private
func sortMatrixByCol(data []Matrix, col int) {
	for k, _ := range data {
		data[k].col = col - 1
	}
	sort.Stable(MatrixSort(data))
}
