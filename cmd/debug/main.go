package main

import (
	"fmt"
	"math"
)

type Range struct {
	Min int64
	Max int64
}

func main() {
	mult := int64(math.MaxInt16) / int64(math.MaxInt8)
	fmt.Println(mult)
	// a := math.Pow(float64(math.MaxUint8), 2)
	// b := float64(math.MaxUint16)

	// fmt.Println(a == b, a, b)

	// n := 25

	// rands := make([]float64, n)
	// for idx, _ := range rands {
	// 	rands[idx] = rand.Float64()
	// }

	// // make 8bit

}
