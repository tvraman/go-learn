//newton Rapson in Go 


package main

import (
	"fmt"
	"math"
)

const delta = 1e-6

func Sqrt(x float64) float64 {
	z := x
	n :=  0.0
	count:= 0
	for   math.Abs(n-z) > delta {
		n, z = z, z-(z*z-x)/(2*z)
		fmt.Println(n)
		count++
	}
	fmt.Println("Converged after: ", count)
	return z
}

func main() {
	for x :=2; x < 10; x++ {
		nr, msq := Sqrt(float64(x)), math.Sqrt(float64(x))
	fmt.Println(nr, msq, nr-msq)
	}
}
