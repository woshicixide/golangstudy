package main

import (
	"fmt"
	"math"
)

type MyFloat float64

// You can only declare a method with a receiver whose type is defined in the same package as the method.
func (self MyFloat) Abs() float64 {
	if self < 0 {
		return float64(-self)
	}
	return float64(self)
}
func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
