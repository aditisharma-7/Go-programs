package main

import (
	"errors"
	"fmt"
)

func main() {

	var s1 string
	var s2 string

	//var s1[17] string{"G","A","G","C","C","T","A","C","T","A","A","C","G","G","G","A","T"}
	//var s1[17] string{"C","A","T","C","G","T","A","A","T","G","A","C","G","G","C","C","T"}

	fmt.Println("enter value for strand 1")
	fmt.Scanln(&s1)
	fmt.Println("enter value for strand 2")
	fmt.Scanln(&s2)
	dist, err := Distance(s1, s2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("the haming distance is ", dist)
	}
}

func Distance(s1, s2 string) (int, error) {
	var count int
	count = 0
	var i int

	if len(s1) == len(s2) {
		for i = 0; i < len(s1); i++ {
			if s1[i] != s2[i] {
				count++
			}
		}
		return count, nil
	}
	return count, errors.New("both strings are not equal")
	//panic("Implement the distance functions")
}
