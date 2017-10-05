// Author    :   Sarah Carroll
// Student ID:   G00330821
// https://play.golang.org/p/Ma2GXvj3XP  https://golang.org/src/sort/sort.go
// GO Program to Merge 2 Sorted Lists.

package main

import (
	"fmt"
	"sort"
)

func main() {

	var i, x, y int

	primes := []int{13, 11, 7, 5, 3, 2}           // unsorted
	nonPrimes := []int{14, 12, 10, 9, 8, 6, 4, 1} // unsorted

	merged := [14]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	fmt.Println("Lists at Start")
	fmt.Println("--------------")

	fmt.Println("primes    =", primes)
	fmt.Println("nonPrimes =", nonPrimes)
	fmt.Println("merged    =", merged)

	p := len(primes)
	n := len(nonPrimes)
	m := len(merged)

	fmt.Println("Length of primes    =", p)
	fmt.Println("Length of nonPrimes =", n)
	fmt.Println("Length of merged    =", m)

	// Sort the 2 arrays before merging

	sort.Ints(primes)
	sort.Ints(nonPrimes)

	i = 0 // merged index
	x = 0 // primes index
	y = 0 // non-primes index

	// Merge lists in ascending int order

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

	fmt.Println("Lists at Finish")
	fmt.Println("---------------")
	fmt.Println("primes    =", primes)
	fmt.Println("nonPrimes =", nonPrimes)
	fmt.Println("New merged=", merged)
}
