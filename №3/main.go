package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, t int
	_, _ = fmt.Fscan(in, &n, &t)

	employees := make([]int, n)
	for i := 0; i < len(employees); i++ {
		_, _ = fmt.Fscan(in, &employees[i])
	}
	var index int
	_, _ = fmt.Fscan(in, &index)
	index -= 1
	minutes := 0
	left, right := 0, 0
	for i := index; i > 1; i-- {
		left += employees[i] - employees[i-1]
	}
	for i := index; i < len(employees)-1; i++ {
		right += employees[i+1] - employees[i]
	}
	fmt.Printf("%v %v \n", left, right)
	switch {
	case employees[index] == employees[0] || left <= index:
		for i := 0; i < len(employees)-1; i++ {
			minutes += employees[i+1] - employees[i]
		}
	case employees[index] == employees[len(employees)-1] || right <= index:
		for i := len(employees) - 1; i > 1; i-- {
			minutes += employees[i] - employees[i-1]
		}
	default:
		if left > right {
			minutes = right + left + (employees[len(employees)-1] - employees[index-1])
		} else {
			minutes = left + right + (employees[index+1] - employees[0])
		}
	}
	fmt.Fprintln(out, minutes)
}
