// Author    :   Sarah Carroll
// Student ID:   G00330821
//Adapter from: http://www.golangpro.com/2016/01/check-string-palindrome-go.html
// GO Program to check if an entered string is a plaindrome.

package main

import (
	"fmt"
	"strings"
)

func main() {

	var input string
	fmt.Println("Enter string:")
	fmt.Scanf("%s\n", &input)
	
	input = strings.ToLower(input)			// make sure it's all in lowercase
	fmt.Println(isP(input))
}

//Function to test if the string entered is a Palindrome

func isP(s string) string {
	mid  := len(s) / 2
	last := len(s) - 1
 
	for i := 0; i < mid; i++ {
		if s[i] != s[last-i] {
			return "NO. It's not a Palindrome."
		}
	}
	
	return "YES! You've entered a Palindrome"
}
