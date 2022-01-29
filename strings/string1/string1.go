//Given a string, say S, print it in reverse manner eliminating the repeated characters and spaces.
/*Input: S = "GEEKS FOR GEEKS"
Output: "SKEGROF"*/
package main

import "fmt"

/*func reverse(s1 string)string{
	//var temp byte
	for i:=0;i< len(s1);i++{
		for j:=len(s1);j>0;j--{
			temp := s1[i]
			s1[i] s1[j]
			s1[j] :=temp
		}
	}
	return s1
}*/
func reverse(s1 string) (reverse string) {

	for i := len(s1) - 1; i >= 0; i-- {
		reverse += string(s1[i])
	}
	return reverse
}

func main() {
	//var s1 string
	//fmt.Println("enter string you want to reverse")
	//fmt.Scanln(&s1)
	s := reverse(s1)
	//s := reverse("aditi divas")
	fmt.Println("reversed string is", s)
}
