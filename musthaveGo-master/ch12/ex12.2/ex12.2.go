// ch12/ex12.2/ex12.2.go
package main

const Y int = 3 // ❶ 상수

func main() {
	x := 5                     // ❷ 변수
	a := [x]int{1, 2, 3, 4, 5} // ❸ 배열의 길이로는 상수만 선언 가능합니다.

	b := [Y]int{1, 2, 3} // ➍

	var c [10]int // ➎
}
