package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20
	var f float32 = 1.2

	fmt.Print("a:", a, "b:", b)
	fmt.Println("a:", a, "b:", b, "f: ", f)
	fmt.Printf("a: %v b: %v, f: %v \n", a, b, f)
}
