package main

import "fmt"

var g int = 10 // package 전역 변수

func main() {
	var m int = 20

	{ // s 의 범위
		var s int = 50
		fmt.Println(m, s, g)
	}

	// s 범위 벗어나서 에러
	// m = s + 20
}
