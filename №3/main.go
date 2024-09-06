package main

import (
	"bufio"
	"fmt"
	"math"
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

	neighbour := int(math.Min(float64(employees[index]-employees[0]), float64(employees[len(employees)-1]-employees[index])))

	if neighbour <= t || index == 0 || index == len(employees)-1 {
		minutes = employees[len(employees)-1] - employees[0]
	} else {
		minutes = neighbour + (employees[len(employees)-1] - employees[0])
	}
	fmt.Fprintln(out, minutes)
}
