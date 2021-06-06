package main

import "fmt"

func main() {
	//// strings
	//var nameOne string = "SungJin"
	//var nameTwo = "Jang"
	//var nameThree string
	//
	//fmt.Println(nameOne, nameTwo, nameThree)
	//
	//nameOne = "peach"
	//nameThree = "bowser"
	//
	//fmt.Println(nameOne, nameTwo, nameThree)
	//
	//nameFour := "yoshi"
	//fmt.Println(nameFour)
	//
	//var ageOne int = 20
	//var ageTwo = 30
	//ageThree := 40
	//
	//fmt.Println(ageOne, ageTwo, ageThree)
	//
	//var numOne int8 = 127
	//var numTwo int8 = -128
	//var numThree uint8 = 255
	//fmt.Println(numOne, numTwo, numThree)
	//
	//var scoreOne float32 = 25.98
	//var scoreTwo float64 = 1238174813749173249812.98
	//scoreThree := 1.5
	//fmt.Println(scoreOne, scoreTwo, scoreThree)

	// Printf (formatted strings)
	age := 26
	name := "sung"
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %q and my name is %q \n", age, name)
	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("you scored %0.1f points! \n", 225.55)

	//Springf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("the saved string is:", str)

}
