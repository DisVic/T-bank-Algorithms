package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var str string
	_, _ = fmt.Fscan(in, &str)
	n := new(big.Int)
	n.SetString(str, 10)
	var result int

	number := big.NewInt(1)
	for number.Cmp(n) < 0 {
		result++
		number.Mul(number, big.NewInt(2))
	}
	fmt.Fprint(out, result)
}
