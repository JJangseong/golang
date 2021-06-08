package main

import (
	"fmt"
	"sort"
	"strings"
)

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
	//age := 26
	//name := "sung"
	//fmt.Printf("my age is %v and my name is %v \n", age, name)
	//fmt.Printf("my age is %q and my name is %q \n", age, name)
	//fmt.Printf("age is of type %T \n", age)
	//fmt.Printf("you scored %0.1f points! \n", 225.55)
	//
	////Springf (save formatted strings)
	//var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	//fmt.Println("the saved string is:", str)

	// Arrays & Slices
	//var ages [3]int = [3]int{20, 25, 30}
	//var ages = [3]int{20, 25, 30}
	//names := [4]string{"yoshi", "mario", "peach", "bowser"}
	//names[1] = "luigi"
	//
	//fmt.Println(ages, len(ages))
	//fmt.Println(names, len(names))
	//
	//// slices ( use arrays under the hood )
	//var scores = []int{100, 50, 60}
	//scores[2] = 25
	//scores = append(scores, 85)
	//fmt.Println(scores, len(scores))
	//
	//// slice ranges
	//rangeOne := names[1:3]
	//rangeTwo := names[2:]
	//rangeThree := names[:3]
	//fmt.Println(rangeOne, rangeTwo, rangeThree)
	//
	//rangeOne = append(rangeOne, "koopa")
	//fmt.Println(rangeOne)

	// The Standard Library

	greeting := "hello there friends"
	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ll"))
	fmt.Println(strings.Split(greeting, " "))

	// the original value is unchanged
	fmt.Println("original String value = ", greeting)

	ages := []int{45, 20, 35, 30, 70, 60, 50, 25}
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

	names := []string{"yoshi", "mario", "peach", "bowser"}

	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, ""))
}
