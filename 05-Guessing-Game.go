//Author Sarah Carroll
//Date :21/09/2017

/* Adapter from: https://stackoverflow.com/questions/20895552/how-to-read-input-from-console-line
https://golang.org/pkg/fmt/
https://www.youtube.com/watch?v=gh1yOouqFs0

*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	guessesTaken := 0
	guess := 0
	//random number generater
	rand.Seed(time.Now().UnixNano())
	var number int

	//pick raandom number between 1 and 100
	number = rand.Intn(100) + 1

	//fmt.Println(number)

	//hard code game for initial testing
	//fmt.Print("Random Number is",number)
	//number:=5

	//rand.Seed(20)

	for guessesTaken < 6 {

		fmt.Println("\nTake a guess.") // There are four spaces in front of print.
		fmt.Scanf("%d\n", &guess)

		if guess < number {
			fmt.Println("Your guess is too low.")
		}

		if guess > number {
			fmt.Println("Your guess is too high.")
		}

		if guess == number {
			fmt.Println("yippeee, You guessed correctly!! ")
			break

		}
		guessesTaken++
		fmt.Printf("Guesses taken= %d\n", guessesTaken)
	}

	fmt.Print("You guessed too many times:Random Number was ", number)

}
