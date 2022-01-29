package main

import "fmt"

//to check for fastattack
func fastattack(isknightsleep bool) bool {
	return isknightsleep
}

//to check for spying
func canspy(isknightsleep, isprisnorsleep, isarchersleep bool) bool {
	return !(isknightsleep && isprisnorsleep && isarchersleep)
}

//to check for signal sending condition
func signal(isprisnorsleep, isarchersleep bool) bool {
	return (!isprisnorsleep) && isarchersleep
}

//to check for sneaking in and freeing prisnor
func freeprisnor(ishasdog, isarchersleep, isprisnorsleep, isknightsleep bool) bool {
	return (ishasdog && isarchersleep) || (!isprisnorsleep && isarchersleep && isknightsleep)
}

/*func main() {
	var isknightsleep bool
	var isprisnorsleep bool
	var isarchersleep bool
	var ishasdog bool
	isarchersleep, isprisnorsleep, isknightsleep, ishasdog = true, true, true, true //asigning values
	a := fastattack(isknightsleep)
	b := canspy(isknightsleep, isprisnorsleep, isarchersleep)
	c := signal(isprisnorsleep, isarchersleep)
	d := freeprisnor(ishasdog, isarchersleep, isprisnorsleep, isknightsleep)
	fmt.Println("the result is :", a, b, c, d)
}*/
func main() {
	var isknightsleep bool
	var isprisnorsleep bool
	var isarchersleep bool
	var ishasdog bool
	var action string
	//isarchersleep, isprisnorsleep, isknightsleep, ishasdog = true, true, true, true //asigning values
	fmt.Println("what action do you want to perform type attack / spy / signal / free")
	fmt.Scanln(&action)
	fmt.Println("is the knight sleeping? true/false")
	fmt.Scanln(&isknightsleep)
	fmt.Println("is the Archer sleeping? true/false")
	fmt.Scanln(&isarchersleep)
	fmt.Println("is the prisoner sleeping? true/false")
	fmt.Scanln(&isprisnorsleep)
	fmt.Println("do you have a dog? true/false")
	fmt.Scanln(&ishasdog)
	if action == "attack" {
		if fastattack(isknightsleep) == true {
			fmt.Println("you can do a fastattack")
		}
	} else {
		fmt.Println("you cant make a fastattack")
	}
	if action == "spy" {
		if fastattack(isknightsleep) == true {
			fmt.Println("you can spy")
		}
	} else {
		fmt.Println("you can not spy ")
	}
	if action == "signal" {
		if fastattack(isknightsleep) == true {
			fmt.Println("you can send a signal")
		}
	} else {
		fmt.Println("you can not send a signal")
	}
	if action == "free" {
		if fastattack(isknightsleep) == true {
			fmt.Println("you can free the prisoner")
		}
	} else {
		fmt.Println("you can not free the prisoner")
	}
	/*a := fastattack(isknightsleep)
	b := canspy(isknightsleep, isprisnorsleep, isarchersleep)
	c := signal(isprisnorsleep, isarchersleep)
	d := freeprisnor(ishasdog, isarchersleep, isprisnorsleep, isknightsleep)
	fmt.Println("the result is :", a, b, c, d) */
	
}
