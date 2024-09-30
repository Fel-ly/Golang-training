package main

import "fmt"

//creating a struct
type Person struct{
	Name string
	Age int
	Gender string
	Address Address
}
type Address struct{
	City string
	Street string
}

func main(){
	// using  a struct created
	person := Person{
		Name: "Felly",
		Age: 21,
		Gender: "Female",
		Address: Address{ //nested structs
			City: "Nairobi",
			Street: "Majani",
		},
	}

	fmt.Println(person) //prints all properties
	fmt.Println(person.Name, "is a", person.Age, "year old", person.Gender)

	fmt.Println(person.Name, "is a", person.Age, "year old", person.Gender, "who lives in", person.Address.City, "city,", person.Address.Street, "street.")


}