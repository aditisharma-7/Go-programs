//Find the Frequency
//Given a vector of N positive integers and an integer X. The task is to find the frequency of X in vector.
package main

import "fmt"

func frequency(arr [6]int, x int) int {
	var count int
	count = 0
	for i := 0; i < 6; i++ {
		if arr[i] == x {
			count++
		}
	}
	//fmt.Printf("%d", count)
	return count
}
func main() {
	//arr :=[]int {}
	//arr := make([]int, n)
	var x int
	fmt.Println("enter which element's frequency you want to count")
	fmt.Scanf("%d", x)
	arr := [6]int{1, 1, 1, 2, 1, 2}
	f := frequency(arr, x)
	fmt.Println("the frequency is", f)
	//fmt.Printf("%d", f)
}
