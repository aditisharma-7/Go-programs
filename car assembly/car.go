package main

import "fmt"

//import "errors"

//to analyz success rate
func successrate(speed int) int {
	if speed == 0 {
		return 0
	}
	if speed >= 1 && speed <= 4 {
		return 100
	}
	if speed >= 5 && speed <= 8 {
		return 90
	}
	if speed >= 9 && speed <= 10 {
		return 77
	} else {
		fmt.Println("the speed you entered is wrong")
		return 0

	}
}

//to calculate production rate per our
func rateperhour(srate float64, speed int) float64 {
	var rph float64
	rph = 221 * (float64(speed) * srate)
	//fmt.Println(rph)
	return rph
}

//to analyz  rateperminute
func rateperminute(b float64) int {
	var rpm int
	var x int
	x = int(b)
	rpm = x / 60
	return rpm
}
func limitedproduct(b float64) float64 {
	if b <= 1000 {
		return b
	}
	return 1000
}

func main() {
	var speed int
	var srate float64
	fmt.Println("enter the speed of the assembly line")
	fmt.Scanf("%d", &speed)
	a := successrate(speed)
	srate = float64(a / 100)
	b := rateperhour(srate, speed)
	c := rateperminute(b)
	d := limitedproduct(b)
	fmt.Println("the success rate is ", a, "%")
	fmt.Println("the production rate per hour is ", b)
	fmt.Println("the production rate per minute is ", c)
	fmt.Println("the product limit is ", d)

}
