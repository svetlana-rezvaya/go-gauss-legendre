package main

import (
	"fmt"
	"math"
)

func main() {
	a := 1.0
	b := 1.0 / math.Sqrt(2)
	t := 1.0 / 4.0
	p := 1.0
	for math.Abs(a-b) > 0.000001 {
		aNext := (a + b) / 2.0
		bNext := math.Sqrt(a * b)
		tNext := t - p*math.Pow(a-aNext, 2.0)
		pNext := 2.0 * p

		a = aNext
		b = bNext
		t = tNext
		p = pNext
	}

	pi := math.Pow(a+b, 2) / (4 * t)
	fmt.Println(pi)
}
