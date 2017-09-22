//Author Sarah Carroll
//Date :21/09/2017

package main

import "fmt"

func factorial(x uint64) uint64 { //uint64 works for large integers
	if x == 0 {
		return 1
	}
	
	return x * factorial(x-1)
}

func main() {
	var calcFactorial uint64
	var total uint64
	var rem uint64
	
	x := uint64(15)
	
	calcFactorial = factorial(x)
	fmt.Println("The Factorial of", x, " = ", calcFactorial)
	
	total = 0
	
	//for loop doesnt need ()
	for calcFactorial > 0 {
		rem = calcFactorial % 10
		total = total + rem
		calcFactorial = calcFactorial / 10
	}
	
	fmt.Println("The Sum of the Factorial of", x, "=", total)
	
}