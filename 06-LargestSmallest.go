// Author    :   Sarah Carroll
// Student ID:   G00330821
// Adapter from:https://gist.github.com/pyk/10519339
// GO Program to print out Large and Smallest Values in a List.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//variables
	var i, n, large, small int
	var nonPrimes [100]int
	rand.Seed(time.Now().UnixNano())

	fmt.Println("\nNumber in the list.") // There are four spaces in front of print.
	fmt.Scanf("%d\n", &n)

	//list of n random numbers usign rand.seed
	for i = 0; i < n; i++ {
		nonPrimes[i] = (rand.Intn(100) + 1)
	}
	/*for i = 0; i < 10; i++ {
		fmt.Print(nonPrimes[i], ",")
	}*/
	//list in the form of an array
	//nonPrimes := [8]int{9,12,6,14,1,10,4,8}
	fmt.Print("List = ")

	for i = 0; i < n; i++ {
		fmt.Printf("%d ", nonPrimes[i])
	}

	//n = len(nonPrimes)

	fmt.Println("\nLength of List =", n)

	//set smallest and largest
	small = 999999
	large = -1

	//for loop to compare list with set smallest and largest
	for i = 0; i < n; i++ {

		if nonPrimes[i] < small {
			small = nonPrimes[i]
		}

		if nonPrimes[i] > large {
			large = nonPrimes[i]
		}

	}

	//print out results

	fmt.Println("Smallest Value =", small)
	fmt.Println("Largest  Value =", large)

}
