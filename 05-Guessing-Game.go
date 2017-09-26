//Author Sarah Carroll
//Date :21/09/2017

/* Adapter from: https://stackoverflow.com/questions/20895552/how-to-read-input-from-console-line
https://golang.org/pkg/fmt/
https://www.youtube.com/watch?v=gh1yOouqFs0

*/

package main

	import "fmt"
    //import "math/rand"

    func main() {
        guessesTaken := 0
        guess :=0;

        number:=5
       
        //rand.Seed(20) 

        for guessesTaken < 6{
            //fmt.Scanf()
            fmt.Println("\nTake a guess.") // There are four spaces in front of print.
            fmt.Scanf("%d\n",&guess)
            //fmt.Scanf()
        
            //guessesTaken++;
        
            if guess < number{
                fmt.Println("Your guess is too low.")
            }
        
            if guess > number{
                fmt.Println("Your guess is too high.")
            }
        
            if guess == number{
                fmt.Println("yippeee, You guessed correctly!! ")
                break

            }
                guessesTaken++
                fmt.Printf("Guesses taken= %d",guessesTaken)
         }

    }