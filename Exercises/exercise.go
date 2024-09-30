package main

import "fmt"
import "math"

func main(){

	exercise1()
	exercise2()
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
	fmt.Printf("bigI (uint64) = %v\n\n\n", bigI)
}


//Exeercise 2
func exercise2(){

	fmt.Printf("This is exercise 2, on maps\n")

	subjectAverages := map[string]int{"MATHS": 68, "ENGLISH": 76, "kISWAHILI": 62, "PHYSICS": 60, "CHEMISTRY": 58, "BUSINESS": 70}
	fmt.Printf("The subject averages for this class are %v\n", subjectAverages)

	englishAverage := subjectAverages["ENGLISH"]
	fmt.Printf("The average score for english was %v\n", englishAverage)

	//using slices as map values
	bestStudents := []string{"Ian", "Talia", "Hannah", "Paul", "Renee"}
	averageStudents := []string{"William", "John", "Faith", "Humphrey", "Paula", "Samantha","Valentine"}
	belowAverageStudents := []string{"Tim", "Amna", "Wesy"}

	studentPerformance := map[string][]string{"Best": bestStudents,"Average": averageStudents, "Below average": belowAverageStudents}
	fmt.Printf("The student performance in this class was %v\n", studentPerformance)

	//using a maps as map values
	bestStudentsMarks := map[string]int{"Ian": 84, "Talia": 84, "Hannah": 80, "Paul": 78, "Renee":76}
	averageStudentsMarks := map[string]int{"William": 58, "John":54, "Faith": 56, "Humphrey":50, "Paula":48, "Samantha":46,"Valentine":46}
	belowAverageStudentsMarks :=map[string]int{"Tim":38, "Amna":32, "Wesy": 28}

	studentMarks := map[string]map[string]int{"Best Marks": bestStudentsMarks, "Average Marks": averageStudentsMarks, "Below average Marks": belowAverageStudentsMarks}
	fmt.Printf("The students scored as follows %v\n", studentMarks)

}
