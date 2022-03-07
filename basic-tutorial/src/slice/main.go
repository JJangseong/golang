package main

import "fmt"

func main() {
	var a []int
	a = []int{1, 2, 3}
	a[1] = 10
	fmt.Println(a)

	s := make([]int, 5, 10)
	fmt.Println(len(s), cap(s), s)

	var b []int
	if b == nil {
		println("Nil Slice")
	}

	println(len(b), cap(b))

	c := []int{0, 1, 2, 3, 4, 5}
	c = c[2:5]     // 2, 3, 4
	c = c[1:]      // 3, 4
	fmt.Println(c) // 3, 4 출력

	d := []int{0, 1}
	// 하나 확장
	d = append(d, 2) // 0, 1, 2
	// 복수 요소들 확장
	d = append(d, 3, 4, 5) // 0,1,2,3,4,5
	fmt.Println(d)

	// len=0, cap=3 인 슬라이스
	sliceA := make([]int, 0, 3)

	// 계속 한 요소씩 추가
	for i := 1; i <= 15; i++ {
		sliceA = append(sliceA, i)
		// 슬라이스 길이와 용량 확인
		fmt.Println(len(sliceA), cap(sliceA))
	}

	fmt.Println(sliceA) // 1 부터 15 까지 숫자 출력

	sliceB := []int{1, 2, 3}
	sliceC := []int{4, 5, 6}

	sliceB = append(sliceB, sliceC...)
	//sliceB = append(sliceA, 4, 5, 6)

	fmt.Println(sliceB) // [1 2 3 4 5 6] 출력

	source := []int{0, 1, 2}
	target := make([]int, len(source), cap(source)*2)
	copy(target, source)
	fmt.Println(target)               // [0 1 2 ] 출력
	println(len(target), cap(target)) // 3, 6 출력
}
