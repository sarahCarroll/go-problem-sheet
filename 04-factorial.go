// Author    :   Sarah Carroll
// Student ID:   G00330821
//
// GO Program to find the large value of 100! and sum the individual digits.
//
// Uses math/big package

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

	for r.Cmp(zero) > 0 { // loop until factorial reduced to zero
		//		fmt.Println("pre r  =",r)
		//		fmt.Println("TEN    =",ten)
		rem = rem.Mod(r, ten) // get MOD of number / 10
		//		fmt.Println("mod r  =",r)
		//		fmt.Println("rem =",rem)
		total.Add(total, rem) // total individual digits
		//quot = r
		//fmt.Println("quot =",quot)
		r.Div(r, ten) // reduce the factorial by facot of 10
		//r = quot
		//		fmt.Println("post r =",r)
	}

	fmt.Println("The Sum of the Factorial Digits of", f, "=", total)

}
