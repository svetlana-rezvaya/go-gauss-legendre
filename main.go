package main

import (
	"flag"
	"fmt"
	"math"
)

func main() {
	epsilon := flag.Float64("e", 0.000001, "epsilon")
	flag.Parse()

	a := 1.0
	b := 1.0 / math.Sqrt(2)
	t := 1.0 / 4
	p := 1.0
	n := 0
	for math.Abs(a-b) > *epsilon {
		aNext := (a + b) / 2
		bNext := math.Sqrt(a * b)
		tNext := t - p*math.Pow(a-aNext, 2)
		pNext := 2 * p

		a = aNext
		b = bNext
		t = tNext
		p = pNext

		n = n + 1
	}

	pi := math.Pow(a+b, 2) / (4 * t)
	fmt.Println("pi =", pi)

	fmt.Println("number of iterations =", n)
}
