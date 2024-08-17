package goiterpatterns

import (
	"fmt"
)

func ExampleSquare() {
	// Set up the pipeline.
	c := Generate(2, 3)
	out := Square(c)

	// Consume the output.
	for n := range out {
		fmt.Println(n)
	}
	// Output: 4
	// 9
}
