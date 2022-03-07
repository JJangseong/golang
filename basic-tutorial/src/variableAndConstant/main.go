package main

func main() {
	var a float32 = 10.
	var i, j, k int = 1, 2, 3
	name := "Go Lang" // 지역번수로만 선언 가능

	// 상수키워드는 const 사용
	const c int = 10
	const s = "HI"
	const (
		Visa   = "Visa"
		Master = "MasterCard"
	)
	// iota 라는 identifier 사용하면 0부터 1씩 증가된 값을 부여받는다.
	const (
		Apple  = iota // 0
		Grape         // 1
		Orange        // 2
	)

	println(a, i, j, k, name)
}
