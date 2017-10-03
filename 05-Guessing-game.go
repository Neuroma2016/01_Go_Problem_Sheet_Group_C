//Guessing game
//Andrew Usevas-group C

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	guess := -1
	counter := 0

	//generate random number and save it to "num"
	rand.Seed(time.Now().UnixNano()) // found this all over the internet...it makes your "num" random every time
	num := rand.Intn(101)

	fmt.Println("Please guess a number from 0-100")

	// guessing stops only when the correct number is guessed
	for guess != num {
		counter++
		fmt.Scanf("%d", &guess)

		//checks users input against random number
		if guess == num {
			fmt.Println("Congrats! Correct number was :", num)
			fmt.Println("It took you ", counter, "guesses.")
			counter--
		} else if guess < 0 {
			fmt.Println("Please enter a valid number!(0-100)")
			counter--
		} else if guess > 100 {
			fmt.Println("Please enter a valid number!(0-100)")
			counter--
		} else if guess > num {
			fmt.Println("Too High!")
		} else if guess < num {
			fmt.Println("Too Low!")
		}
		
	} //for

}
