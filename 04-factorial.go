package main

import (
	"fmt"
	"math/big"
)

func calcFactorial(x *big.Int) *big.Int {
	n := big.NewInt(1)
	if x.Cmp(big.NewInt(0)) == 0 {
		return n
	}
	return n.Mul(x, calcFactorial(n.Sub(x, n)))
}

func main() {

	f := big.NewInt(100)  // Create a Big Int, value 100
	r := calcFactorial(f) // Calculate 100!

	fmt.Println("The Factorial of", f, " = ", r)

	total := big.NewInt(0)
	ten := big.NewInt(10)
	zero := big.NewInt(0)
	rem := big.NewInt(0)
	//quot  := big.NewInt(0)

	for r.Cmp(zero) > 0 {
		//		fmt.Println("pre r  =",r)
		//		fmt.Println("TEN    =",ten)
		rem = rem.Mod(r, ten)
		//		fmt.Println("mod r  =",r)
		//		fmt.Println("rem =",rem)
		total.Add(total, rem)
		//quot = r
		//fmt.Println("quot =",quot)
		r.Div(r, ten)
		//r = quot
		//		fmt.Println("post r =",r)
	}

	fmt.Println("The Sum of the Factorial Digits of", f, "=", total)

}
