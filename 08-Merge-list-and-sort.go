//Merge list and Sort
//Andrew Usevas-group C
//SOURCE: https://gobyexample.com/sorting
//https://blog.golang.org/go-slices-usage-and-internals

package main

import (
	"fmt"
	"sort"
)

func main() {

	var num int

	//two same size slices
	s1 := make([]int, 3)
	s2 := make([]int, 3)
	//slice to store other two(merged) slices
	s3 := make([]int, len(s1)+len(s2))

	//user input for s1
	fmt.Println("Please enter ", len(s1), " numbers to fill the s1 list: ")

	for i := 0; i < len(s1); i++ {
		fmt.Scanf("%d", &num)
		s1[i] = num
	}

	//user input for s2
	fmt.Println("Please enter ", len(s2), " numbers to fill the s2 list: ")

	for i := 0; i < len(s2); i++ {
		fmt.Scanf("%d", &num)
		s2[i] = num
	}

	//appends two (s1,s2) into 1 slice(s3)
	s3 = append(s1, s2...)
	//sorts slice in ascending order
	sort.Ints(s3)

	//display slices
	fmt.Println("s1 slice: ", s1)
	fmt.Println("s2 slice: ", s2)
	fmt.Println("s3 slice(Merged and sorted): ", s3)
}
