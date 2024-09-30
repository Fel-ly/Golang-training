package main

import "fmt"
import "math"

func main(){

	exercise1()
}

//Exercise 1
func exercise1(){
	//  Question 1
	i := 50
	f := float64(i)

	fmt.Printf("Integer i is %d\n", i)
	fmt.Printf("Floating point f is %f\n\n", f)


	// Question 2

	const value = 16

	int := value
	flt := value

	fmt.Printf("Integer value i is %d\n", int)
	fmt.Printf("Floating-point value f is %f\n\n", flt)


	//  Question 3

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