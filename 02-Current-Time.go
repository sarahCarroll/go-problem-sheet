//Author Sarah Carroll
//Date :20/09/2017

// Adapter from:https://tour.golang.org/welcome/4 https://golang.org/pkg/time/ https://stackoverflow.com/questions/5885486/how-do-i-get-the-current-timestamp-in-go

package main

//Package time provides functionality for measuring and displaying time.
import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	//current time date printed in default format
	fmt.Println("The current date and time is", time.Now())

	//format time currently,in hours and mins
	fmt.Println(t.Format("\n15.04"))
}
