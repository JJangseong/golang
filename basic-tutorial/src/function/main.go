package main

func main() {
	msg := "Hello"
	say(msg)
	say2(&msg)
	println(msg)
	say3("H", "e", "l", "l", "o")
	count, total := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	println(count, total)
}

func say(msg string) {
	println(msg)
}

func say2(msg *string) {
	println(*msg)
	*msg = "Changed" //메시지 변경
}

// 가변인자함수
func say3(msg ...string) {
	for _, s := range msg {
		println(s)
	}
}

func sum(nums ...int) (count int, total int) {
	for _, n := range nums {
		total += n
	}
	count = len(nums)
	return
}
