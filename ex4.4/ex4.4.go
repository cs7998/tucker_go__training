package main

import "fmt"

func main() {
	a := 3 // int - 64bit = int64
	var b float64 = 3.5

	var c int = int(b) // 3.5 => 3
	d := float64(a) * b

	var e int64 = 7
	f := a * int(e)

	fmt.Println(a, b, c, d, e, f)
}
