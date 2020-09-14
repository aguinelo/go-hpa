package main

import (
	"fmt"
	"math"
)

func main() {
	x := 0.0001

	for i := 0; i < 1000000; i++ {
		x += math.Sqrt(x)
	}
	fmt.Println("Code.education Rocks!")
}
