//Largest and smallest in list
//Andrew Usevas-group C
//SOURCE: https://golang.org/pkg/container/list/#List.Front
//took me a while to figure this out...i've read that slice is easier and better(more efficient)
//but i made myself a chalenge to get this one working

package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//generate random number and save it to "num"
	rand.Seed(time.Now().UnixNano()) // found this all over the internet...it makes your "num" random every time

	//create a list
	l := list.New()

	//for loop to insert 10 random(numbers) elements into the list
	for i := 0; i < 10; i++ {
		num := rand.Intn(101)
		l.PushFront(num)
	}

	// Iterate through list and print its contents
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	//set highest/smallest to the first element in the list
	highest := l.Front().Value.(int)
	smallest := l.Front().Value.(int)

	//iterate through list and update highest/smallest
	for e := l.Front(); e != nil; e = e.Next() {
		temp := e.Value.(int)
		if highest < temp {
			highest = temp
		}
		if smallest > temp {
			smallest = temp
		}
	}
	//display result
	fmt.Println("Highest number in this List is: ", highest, ", while smallest is: ", smallest)
}
