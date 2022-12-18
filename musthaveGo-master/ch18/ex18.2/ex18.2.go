// ch18/ex18.2/ex18.2.go
package main

import "fmt"

func main() {

	var slice = []int{1, 2, 3} // ❶ 요소가 3개인 슬라이스
	var slice3 = []int{1, 5: 2, 10: 3}

	slice2 := append(slice, 4) // ❷ 요소 추가

	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println(slice3)
}
