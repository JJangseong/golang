package main

import (
	"fmt"
	"math"
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
	//greeting := "hello there friends"
	//fmt.Println(strings.Contains(greeting, "hello"))
	//fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	//fmt.Println(strings.ToUpper(greeting))
	//fmt.Println(strings.Index(greeting, "ll"))
	//fmt.Println(strings.Split(greeting, " "))
	//
	//// the original value is unchanged
	//fmt.Println("original String value = ", greeting)
	//
	//ages := []int{45, 20, 35, 30, 70, 60, 50, 25}
	//sort.Ints(ages)
	//fmt.Println(ages)
	//
	//index := sort.SearchInts(ages, 30)
	//fmt.Println(index)
	//
	//names := []string{"yoshi", "mario", "peach", "bowser"}
	//
	//sort.Strings(names)
	//fmt.Println(names)
	//
	//fmt.Println(sort.SearchStrings(names, ""))

	// Loops
	//x := 0
	//for x < 5 {
	//	fmt.Println("value of x is : ", x)
	//	x++
	//}
	//
	//for i := 0; i < 5; i++ {
	//	fmt.Println("value of i is :", i)
	//}
	//
	//names := []string{"yoshi", "mario", "peach", "bowser"}
	//
	//for _, v := range names {
	//	fmt.Println(v)
	//}

	// Booleans & Conditionals
	//
	//age := 45
	//
	//fmt.Println(age <= 50)
	//fmt.Println(age >= 50)
	//fmt.Println(age == 45)
	//fmt.Println(age != 50)
	//
	//if age < 30 {
	//	fmt.Println("age is less than 30")
	//} else if age < 40 {
	//	fmt.Println("age is less than 40")
	//} else {
	//	fmt.Println("age is not less than 45")
	//}
	//names := []string{"yoshi", "mario", "peach", "bowser"}
	//
	//for index, value := range names {
	//	if index == 1 {
	//		fmt.Println("continuing at pos ", index)
	//		continue
	//	}
	//
	//	if index > 2 {
	//		fmt.Println("breaking at pos ", index)
	//		break
	//	}
	//
	//	fmt.Printf("the value at pos %v is %v \n", index, value)
	//}

	// Using Functions
	//sayGreeting("mario")
	//sayBey("luigi")
	//cycleNames([]string{"yoshi", "mario", "peach", "bowser"}, sayGreeting)
	//cycleNames([]string{"yoshi", "mario", "peach", "bowser"}, sayBey)
	//
	//a1 := circleArea(10.5)
	//a2 := circleArea(15)
	//fmt.Println(a1, a2)
	//fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f", a1,a2)

	// Multiple Return Values
	//fn, sn := getInitials("tifa lockhart")
	//
	//fmt.Println(fn, sn)

	// Package Scope
	//sayHello("mario")
	//
	//for _, v := range points {
	//	fmt.Println(v)
	//}

	//// Maps
	//menu := map[string]float64{
	//	"soup": 4.99,
	//	"pie": 7.99,
	//	"salad": 6.99,
	//	"toffee pudding": 3.55,
	//}
	//
	//fmt.Println(menu)
	//fmt.Println(menu["soup"])
	//
	//// looping maps
	//for k, v := range menu {
	//	fmt.Println(k, " - ", v)
	//}
	//
	//// ints as key type
	//phonebook := map[int]string {
	//	1: "mario",
	//	2: "luigi",
	//}
	//
	//fmt.Println(phonebook)
	//fmt.Println(phonebook[1])
	//
	//phonebook[2] = "hello"
	//fmt.Println(phonebook)
	//
	//phonebook[1] = "world"
	//fmt.Println(phonebook)

	// Pass By Value
	//name := "tifa"
	//
	//name = updateName(name)
	//fmt.Println(name)
	//
	//menu := map[string]float64 {
	//	"pie": 5.6,
	//}
	//updateMenu(menu)
	//fmt.Println(menu)

	// Pointers
	name := "tifa"

	updateName(name)
	fmt.Println("memory address of name is : ", &name)
	m := &name
	fmt.Println("memory address:", &m)
	fmt.Println("Value at memory address:",*m)

	updateNameByPointer(m)
	fmt.Println(name)
}

func updateNameByPointer(x *string) {
	*x = "wedge"
}

func updateName(x string) string {
	x = "wedge"
	return x
}

func updateMenu(y map[string]float64) {
	y["pie"] = 1.1
}

func getInitials(n string) (string, string){
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayBey(n string) {
	fmt.Printf("Goodbye %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, value := range n {
		f(value)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}