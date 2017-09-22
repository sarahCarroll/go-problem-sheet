//Author Sarah Carroll
//Date :20/09/2017

// Adapter from: http://wiki.c2.com/?FizzBuzzTest

package main

	import "fmt"
    //import "math/rand"

    func main() {
        guessesTaken := 0
        guess :=0;

        number:=5
       
        //rand.Seed(20) 

        for guessesTaken < 6{

            fmt.Println("\nTake a guess.") // There are four spaces in front of print.
            fmt.Scanf("%d",&guess)
        
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