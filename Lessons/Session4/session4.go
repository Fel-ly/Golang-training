package main

import "fmt"

//creating a struct
// Note: structs are created outside the main function for re-usability
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

type Calculator struct{  //Structs with Functions as Fields
	Add func(a, b int) int
}

func updateAge(p *Person, newAge int){	//Pointers to Structs

	p.Age = newAge
}

func main(){
	// using  a struct created
	person := Person{
		Name: "Jane",
		Age: 21,
		Gender: "Female",
		Address: Address{ //nested structs
			City: "Nairobi",
			Street: "Majani",
		},
	}

	fmt.Println(person) //prints all properties
	fmt.Println(person.Name, "is a", person.Age, "year old", person.Gender)

	person.Age = 25// changing a value
	fmt.Println(person.Name, "is a", person.Age, "year old", person.Gender, "who lives in", person.Address.City, "city,", person.Address.Street, "street.\n")


	//Structs with Functions as Fields
	calc := Calculator{
		Add: func(a, b int) int{
			return (a + b)
		},
	}

	result := calc.Add(20, 3)
	fmt.Println("The sum is", result, "\n")


	//Pointers to Structs
	updateAge(&person, 20)
	fmt.Println(person.Age)
	fmt.Println(person.Name, "is a", person.Age, "year old", person.Gender)
}