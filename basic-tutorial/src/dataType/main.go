package main

func main() {
	var boolean bool = false
	// int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64 uintptr
	// int = 음수, 양수 포함
	// uint = 양수만
	var integer int = 10
	// float32, float64, complex64, complex128
	var float float32 = 1.
	/*
		기타 타입
		byte: uint8과 동일, 바이트 코드에 사용
		rune: int32와 동일, 유니코드 코드포인트에 사용
	*/

	/*
		문자열은 `hello`, "hello" 이 두가지로 선언
		` 로 선언한 문자열은 입력된 문자열 그대로 출력 ( 예 - `Hello \n World`, 출력 결과 - Hello \n World )
		- 복수 라인을 많이 사용할때 사용함
	*/
	// immutable 타입임 ( 한번 선언되면 수정 불가)
	var str string = "stringType"

	println(boolean, str, integer, float)

	// Type Conversion
	var i int = 100
	var u uint = uint(i)
	var f float32 = float32(i)
	println(f, u)

	str1 := "ABC"
	bytes := []byte(str1)
	str2 := string(bytes)
	println(bytes, str2)

}
