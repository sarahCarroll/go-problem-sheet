// Author    :   Sarah Carroll
// Student ID:   G00330821
//
// GO Program to Merge 2 Sorted Lists.

package main

import "fmt"

func main() {

	var i, x, y int

	primes := [6]int{2, 3, 5, 7, 11, 13}
	nonPrimes := [8]int{1, 4, 6, 8, 9, 10, 12, 14}
	merged := [14]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	fmt.Println("primes    =", primes)
	fmt.Println("nonPrimes =", nonPrimes)
	fmt.Println("merged    =", merged)

	p := len(primes)
	n := len(nonPrimes)
	m := len(merged)

	fmt.Println("Length of primes    =", p)
	fmt.Println("Length of nonPrimes =", n)
	fmt.Println("Length of merged    =", m)

	i = 0 // merged index
	x = 0 // primes index
	y = 0 // non-primes index

	for x < p && y < n {

		if primes[x] < nonPrimes[y] {
			merged[i] = primes[x]
			x++
		} else {
			merged[i] = nonPrimes[y]
			y++
		}

		i++

	}

	for x < p {
		merged[i] = primes[x]
		i++
		x++
	}

	for y < n {
		merged[i] = nonPrimes[y]
		i++
		y++
	}

	fmt.Println("New merged=", merged)
}
