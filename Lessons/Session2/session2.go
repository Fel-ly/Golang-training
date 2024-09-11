package main

import "fmt"

func main(){

	arrays()

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

	fmt.Printf("There are 8 planets %v\n", planets)
	fmt.Printf("The biggest planet is %v\n", planets[4])
	fmt.Printf("The ring planet is %v", planets[5])

}


