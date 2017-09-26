// Author    :   Sarah Carroll
// Student ID:   G00330821
// Adapter from: https://golang.org/pkg/sort/
// GO Program to Merge 2 Sorted Lists.


package main

import "fmt"

func main() {

	var i, x, y int
	
	primes := [6]int{2, 3, 5, 7, 11, 13}
	nonPrimes := [8]int{1,4,6,8,9,10,12,14}
	merged := [14]int{0,0,0,0,0,0,0,0,0,0,0,0,0,0}
	
	fmt.Println("primes    =",primes)
	fmt.Println("nonPrimes =",nonPrimes)
	fmt.Println("merged    =",merged)
	
	p := len(primes)
	n := len(nonPrimes)
	m := len(merged)
	
	fmt.Println("Length of primes    =", p)
	fmt.Println("Length of nonPrimes =", n)
	fmt.Println("Length of merged    =", m)
	
	x = 0
	y = 0
	
	for i = 0; i < 14; i++ {
			fmt.Println("Loop:",i,"x=",x,"y=",y)
			if x != -1 && nonPrimes[y] > primes[x] {
				merged[i] = primes[x]
				x++
				if x == p {
					x = -1
				}
			} else {
				merged[i] = nonPrimes[y]
				y++
				if y == n {
					y = -1
				}
			}
	}
	
	fmt.Println("New merged=",merged)
}
