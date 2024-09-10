package main

import "fmt"
import "math"

func main(){
	//  Question 1
	i := 50
	f := float64(i)

	fmt.Printf("Integer i is %d\n", i)
	fmt.Printf("Floating point f is %f\n\n", f)

	
	question2() // calling the question2 function
	question3() // calling the question3 function

}

// Question 2
func question2(){
	const value = 16

	i := value
	f := value

	fmt.Printf("Integer value i is %d\n", i)
	fmt.Printf("Floating-point value f is %f\n\n", f)
}

//  Question 3
func question3(){
	var b byte = math.MaxUint8
	var smallI int32 = math.MaxInt32
	var bigI uint64 = math.MaxUint64

	b += 1
	smallI += 1
	bigI += 1

	fmt.Printf("b (byte) = %v\n", b)
	fmt.Printf("smallI (int32) = %v\n", smallI)
	fmt.Printf("bigI (uint64) = %v\n", bigI)

}