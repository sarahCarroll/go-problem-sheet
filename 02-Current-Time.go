//Author Sarah Carroll
//Date :20/09/2017

// Adapter from:https://tour.golang.org/welcome/4 https://golang.org/pkg/time/

package main

//Package time provides functionality for measuring and displaying time.
import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())
}