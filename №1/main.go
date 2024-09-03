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

	var cost, cap, costMB, nextCap int
	_, _ = fmt.Fscan(in, &cost, &cap, &costMB, &nextCap)

	var result int
	if cap >= nextCap {
		result = cost
	} else {
		result = cost + (nextCap-cap)*costMB
	}
	fmt.Fprint(out, result)
}
