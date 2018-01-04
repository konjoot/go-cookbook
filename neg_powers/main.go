package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	var (
		base  int
		power int
	)

	flag.IntVar(&base, "base", 2, "base")
	flag.IntVar(&power, "power", 10, "max negative power")
	flag.Parse()

	numbers := [...]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	d := make([]int, power)

	d[0] = 1 / base
	r := 1 % base
	var i int
	for k := 0; k < power; k++ {
		fmt.Fprint(os.Stdout, "0.")
		for i = 0; i < power; i++ {
			r = 10*r + d[i]
			d[i] = r / base
			r = r % base
			fmt.Fprint(os.Stdout, numbers[d[i]])
		}
		r = 0
		fmt.Fprint(os.Stdout, "\n")
	}
}
