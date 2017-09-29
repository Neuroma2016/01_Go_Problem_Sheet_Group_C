//Fizz Buzz program
//Andrew Usevas-group C

package main

import (
	"fmt"
)

func main() {

	var num int

	//user input
	fmt.Println("Please enter the number:")
	fmt.Scanf("%d", &num)

	//for loop with other conditions
	for i := 1; i <= num; i++ {

		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
