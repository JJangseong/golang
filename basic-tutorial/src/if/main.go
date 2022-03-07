package main

import "fmt"

func main() {
	// () 중괄호를 안써도 된다.
	// 일부 언어들 처럼 1,0,"",null 등을 입력해서는 안된다.
	if 1 == 2 {
		println("Two")
	} else if 1 == 1 {
		println("One")
	}

	// Optional Statement와 함께 실행도 가능 대신 사용된 조건문의 스코프 안에서만 사용 가능
	if val := 1 * 2; val < 20 {
		println(val)
	}

	var name string
	var category = 1

	/*
		go에서 switch 문 특징
		1. case안에 복잡한 expression이 들어갈 수 있다.
		2. switch 뒤에 expression이 없을 수 있다.
		3. break를 안써도 다음으로 넘어가지 않는다. ( default break를 무시하고 쭉 실행될려면 fallthrough 키워드를 사용하면 된다. )
		4. 값 자체말고 변수의 Type에 따라 분기할 수 있다.
	*/
	switch category {
	case 1:
		name = "Paper Book"
	case 2:
		name = "eBook"
	case 3, 4:
		name = "Blog"
	case 2 + 3:
		name = "Test"
	default:
		name = "Other"
	}
	println(name)

	// Expression을 사용한 경우 ( 1 = 0001, 1 << 2 = 0100 = 4, x = 4)
	switch x := category << 2; x - 1 {
	//...
	}

	grade(1)

}

// Expression 사용 안한 경우
func grade(score int) {
	switch {
	case score >= 90:
		println("A")
	case score >= 80:
		println("B")
	case score >= 70:
		println("C")
	case score >= 60:
		println("D")
	default:
		println("No Hope")
	}
}

// 타입 체크
func typeCheck() {
	switch v.(type) {
	case int:
		println("int")
	case bool:
		println("bool")
	case string:
		println("string")
	default:
		println("unknown")
	}
}

// fallthrough 키워드 사용 법
func check(val int) {
	switch val {
	case 1:
		fmt.Println("1 이하")
		fallthrough
	case 2:
		fmt.Println("2 이하")
		fallthrough
	case 3:
		fmt.Println("3 이하")
		fallthrough
	default:
		fmt.Println("default 도달")
	}
}
