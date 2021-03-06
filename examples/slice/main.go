package main

import (
	"fmt"

	r "github.com/sniperkit/rangetype"
)

func main() {
	// Outputs 1 to 10, with 0 digits after "."
	fmt.Println(r.New("1..10").Join(", ", 0))

	// Outputs 2 and 4
	for _, x := range r.Slice([]float64{1.0, 2.0, 3.0, 4.0}, "1..3 step 2") {
		fmt.Print(x, " ")
	}
	fmt.Println()

	// Also outputs 2 and 4
	r.New("(0:6:2)").ForEach(func(x float64) {
		fmt.Print(x, " ")
	})
	fmt.Println()
}
