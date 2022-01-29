//Given an array arr[] and an integer K where K is smaller than size of array,
//the task is to find the Kth smallest element in the given array. It is given that all array elements are distinct.

package main

import "fmt"

func check(arr [6]int, k int) {
	for l := 0; l < 6; l++ {
		for x := 0; x < 6; x++ {
			for y := 0; y < 6; y++ {
				if arr[y] > arr[x] {
					var temp int
					temp = arr[x]
					arr[x] = arr[y]
					arr[y] = temp
				}
			}
		}
	}
	for i := 0; i < 6; i++ {
		//return arr[i]
		fmt.Println("the new array is", arr[i])
		fmt.Printf("%d", arr[k]) //fmt.Printf("%d rd smallest element in array is %d", k, arr[k])
	}

}

func main() {
	var k int
	/*var arr[...] int
	var i int
	 fmt.Println("enter elemnets of array")
	 for i=0;i<=
	 fmt.Scanf(arr[])

	*/
	arr := [6]int{4, 6, 1, 5, 3, 2}
	fmt.Println("enter value for k")
	fmt.Scanf("%d", k)
	if k < 6 {
		check(arr, k)
	} else {
		fmt.Println("error")
	}

}
