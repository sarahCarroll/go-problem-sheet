// Author    :   Sarah Carroll
// Student ID:   G00330821
//
// GO Program to demonstrate Newtons' method for Square Roots.


package main

import (
    "fmt"
    "math"
)

func Newton(x float64) float64 {
    if x == 0 { return 0 }
    z := 1.0
    for i := 0; i < int(x); i++ {
        z = z - ((math.Pow(z, 2) - x) / (2 * z))
    }
    return z
}

func SquareRoot(x float64) float64 {
    return math.Sqrt(x)
}

func main() {
    times := 16
	delta := float64(0.000001)
    for i := 0; i < times; i++ {
        squareroot := SquareRoot(float64(i))
        newton     := Newton(float64(i))
        fmt.Println(i, "squared:")
        fmt.Println("  SquareRoot:", squareroot)
        fmt.Println("  Newton    :", newton)
		diff := math.Abs(squareroot-newton)
		if diff == 0.0 {
	        fmt.Println("  Difference: 0.0")
		} else if diff < delta {
	        fmt.Println("  Difference: 0.0 [", diff, "< DELTA of",delta,"]")
		} else {
	        fmt.Println("  Difference:", diff)
		}
	}
}
