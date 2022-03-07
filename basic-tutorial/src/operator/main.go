package main

func main() {
	a := 10
	b := 20
	// 산술연산자 ( +, -, *, /, % ) + (++, --)
	c := (a + b) / 5
	c++
	println(c)

	// 관계연산자
	println(a == b)
	println(a != b)
	println(a >= b)

	// 논리연산자
	println(a > b && a == 10)
	println(a > b || a == b)

	// Bitwise 연산자 ( 바이너리 AND, OR, XOR와 바이너리 쉬프트 연산자 )
	// 2진수 ( a = 0011, b = 0101 )
	// AND => 둘다 1일때 ( a & b = 0001 )
	// OR => 둘중 하나 1일때 ( a | b = 0111 )
	// XOR => 두개가 다를때 1 ( a ^ b = 0110 )
	// 쉬프트 ( << , >> ) => 특정 비트씩 옆으로 이동 ( a << 1 = 0110, a >> 1 = 0001 )
	d := (a & b) << 5
	println(d)

	// 할당연산자
	a = 100
	a *= 100
	a >>= 2
	a |= 1

	// 포인트 연산자
	var k int = 10
	var p = &k
	println(*p)

}
