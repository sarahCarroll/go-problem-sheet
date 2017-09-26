// Author    :   Sarah Carroll
// Student ID:   G00330821
// Adapter from:https://gist.github.com/pyk/10519339
// GO Program to print out Large and Smallest Values in a List.


package main

import "fmt"

func main() {

	//variables
	var i, n, large, small int
	
	//list in the form of an array
	nonPrimes := [8]int{9,12,6,14,1,10,4,8}
	
	fmt.Println("List =",nonPrimes)
	
	n = len(nonPrimes)
	
	fmt.Println("Length of List =", n)
	
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
	fmt.Println("Smallest Value =",small)
	fmt.Println("Largest  Value =",large)
	
}