package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {

	fmt.Println(s)

	const n = 500000000
	const d = 3e20 / n
	fmt.Println("Value of d is ", d)

	fmt.Println("INT64 value of d is ", int64(d))

	fmt.Println(math.Sin(n))

}

// A const statement can appear anywhere a var statement can.
