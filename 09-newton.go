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
    times := 15
    for i := 0; i < times; i++ {
        squareroot := SquareRoot(float64(i))
        newton     := Newton(float64(i))
        fmt.Println(i, "squared:")
        fmt.Println("  SquareRoot:", squareroot)
        fmt.Println("  Newton    :", newton)
        fmt.Println("  Difference:", math.Abs(squareroot-newton))
    }
}