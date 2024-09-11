package main

import "fmt"

func main(){

	arrays() // calling the arrays function
	slices() //calling the slices function

}

//Arrays
func arrays(){
	var planets [8]string

	planets[0] = "Mercury"
	planets[1] = "Venus"
	planets[2] = "Earth"
	planets[3] = "Mars"
	planets[4] = "Jupiter"
	planets[5] = "Saturn"
	planets[6] = "Uranus"
	planets[7] = "Neptune"

	fmt.Printf("The biggest planet is %v\n", planets[4])
	fmt.Printf("The ring planet is %v\n", planets[5])

	fmt.Printf("There are %d planets in the universe %v\n\n", len(planets), planets)

}

//Slices
func slices(){
	var planets[]string

	planets = append(planets, "Mercury")
	planets = append(planets, "Venus")
	planets = append(planets, "Earth")
	planets = append(planets, "Mars")

	fmt.Printf("We have %d planets listed %v\n",  len(planets), planets)

}
