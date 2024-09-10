package main

import "fmt"

func main(){
	i := 50
	f := float64(i)

	fmt.Printf("Integer i is %d\n", i)
	fmt.Printf("Floating point f is %f\n\n", f)

	// calling the question2 function
	question2()
}

//Question 2
func question2(){
	const value = 16

	i := value
	f := value

	fmt.Printf("Integer value i is %d\n", i)
	fmt.Printf("Floating-point value f is %f\n", f)
}

