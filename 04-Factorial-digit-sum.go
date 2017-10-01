//Factorial program
//Andrew Usevas-group C
//Influenced By: SOURCES: http://www.programmingsimplified.com/java/source-code/java-program-find-factorial

package main

import (
	"fmt"
	"math/big"
)

//calculate factorial....code from https://play.golang.org/p/5lZMq3bd8a works for any (int) number
func factorial(n int64) *big.Int {
	if n < 0 {
		return big.NewInt(1)
	}
	if n == 0 {
		return big.NewInt(1)
	}
	bigN := big.NewInt(n)
	return bigN.Mul(bigN, factorial(n-1))
}

func main() {

	var num int64

	/*	var num , i int64                          // My Original code....works for numbers up to 19
		var fact int64 = 1

		fmt.Print("Please enter number to calculate it's Factorial: ")
		fmt.Scanf("%d", &num)

		if num < 0 {
			fmt.Println("Please enter a valid, non negative number!")
		}else if num == 0{
			fmt.Println("Factorial of ", num ," is = " ,fact)
		}else{
			for i = 1; i <= num; i++{
				fact = fact * i
			}
			fmt.Println("Factorial of ",num," is = ",fact)
		}*/

	//user input
	fmt.Print("Please enter number to calculate it's Factorial: ")
	fmt.Scanf("%d", &num)

	//factorial call and result display
	fmt.Println(factorial(num))

}
