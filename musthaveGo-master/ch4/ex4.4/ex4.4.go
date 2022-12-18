//ch4/ex4.4/ex4.4.go
package main

import "fmt"

func main() {
	a := 3              // int
	var b float64 = 3.5 // float64

	var c int = int(b)  // ❶ float64에서 int로 변환
	d := float64(a * c) // int에서 float64로 변환 계산이 먼저 이루어지고

	var e int64 = 7
	f := int64(d) * e // float64에서 int64로 변환 // int로 먼저 형 변환

	var g int = int(b * 3) // float64에서 int로 변환 10
	var h int = int(b) * 3 // float64에서 int로 변환 g와 값이 다릅니다. 9
	fmt.Println(g, h, f)   //10, 9, 63
}
