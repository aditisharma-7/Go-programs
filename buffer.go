package main

import (
	"bytes"
	"fmt"
)
//importing the bytes package so that buffer can be used

func main() {
	var x bytes.Buffer
	k := "aditi"

	x.WriteString(k) //stores data in buffer variable from variable k 
	x.WriteString(" sharma") //stores data in buffer variable
	fmt.Println(x.String())
}
