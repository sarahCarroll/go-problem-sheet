// Author    :   Sarah Carroll
// Student ID:   G00330821
// Adapter from: https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
// GO Program to reverse characters in an entered string.


package main

import "fmt"

func reverse(value string) string {
    // Convert string to rune slice.
    // ... This method works on the level of runes, not bytes.
    data := []rune(value)
    result := []rune{}

    // Add runes in reverse order.
    for i := len(data) - 1; i >= 0; i-- {
        result = append(result, data[i])
    }

    // Return new string.
    return string(result)
}

func main() {

	var input string
	fmt.Println("Enter string to be reversed:")
	fmt.Scanf("%s\n", &input)
  
    reversed := reverse(input)
    //fmt.Println(input)
    fmt.Println(reversed)

}