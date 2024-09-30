// Maps

package main

import "fmt"

func main(){

	maps()
}

func maps(){
	//creating a map
	eplteamPoints:= map[string]int{"MANCITY": 91, "ARSENAL": 89, "LIVERPOOL": 82, "ASTONVILLA": 66}
	fmt.Printf("The EPL team points are %v\n", eplteamPoints)

	// Adding content into the map
	eplteamPoints["SPURS"] = 66
	eplteamPoints["CHELSEA"] = 60
	fmt.Printf("The EPL team points are %v\n", eplteamPoints)

	//Reading into the map
	mancityPoints := eplteamPoints["MANCITY"]
	fmt.Printf("Man city points are %d\n", mancityPoints)

	chelseaPoints := eplteamPoints["CHELSEA"]
	fmt.Printf("Chelsea points are %d\n", chelseaPoints)

	//Giving a key that doesn't exist in the map gives a zero value
	manuPoints :=  eplteamPoints["MANU"]
	fmt.Printf("Man United points are %d\n", manuPoints)


	// Using a slice as a map's value
	manuLegends := []string{"Ronaldo", "Maguire", "Rooney", "Pogba", "Chicharito", "Rashford"}
	chelseaLegends := []string{"Lampard", "Drogba", "Peter", "Terry", "Kalou", "Ballack"}

	teamLegends := map[string][]string{"MANU":manuLegends, "CHELSEA":chelseaLegends}
	fmt.Printf("Team players are %v\n", teamLegends)


	//Using a map as another map's value
	manuPlayerScores := map[string]int{"Ronaldo": 19, "Maguire": 2, "Pogba": 10, "Chicharito": 15, "Rashford": 13}
	chelseaPlayerScores := map[string]int{"Lampard": 20, "Drogba": 19, "Peter": 0, "Terry": 9, "Ballack": 5}

	teamPlayerScores := map[string]map[string]int{"MANU": manuPlayerScores, "CHELSEA": chelseaPlayerScores}
	fmt.Printf("Team player scores are %v\n", teamPlayerScores)


	//Checking the map length
	fmt.Printf("EPL team points map length is %v\n", len(eplteamPoints))
	fmt.Printf("Team player scores map length is %v\n", len(teamPlayerScores))

	//deleting a value from a map
	delete(eplteamPoints, "ASTONVILLA")
	fmt.Printf("The EPL team points are %v\n", eplteamPoints)
	fmt.Printf("EPL team points map length is %v\n", len(eplteamPoints))
}