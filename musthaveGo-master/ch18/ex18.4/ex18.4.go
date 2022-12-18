// ch18/ex18.4/ex18.4.go
package main

import "fmt"

func changeArray(array2 [5]int) { // ❶ 배열을 받아서 세 번째 값 변경
	array2[2] = 200
}

//array는 배열의 모든 값이 복사 되어 int(8byte) * 5 = 40 byte가 복사된다.
//복사 된 값들은 전혀 다른 메모리 주소를 가지고 있는 array2로 복사된다. 따라서 원래 array값은 변하지 않는다.

func changeSlice(slice2 []int) { // ❷ 슬라이스를 받아서 세 번째 값 변경
	slice2[2] = 200
}

//slice의 기본 구조는 포인터, cap, len으로 이루어져 있는데 포인터는 메모리 주소로 8바이트, cap, len은 int임으로 각각 8바이트
//총 24바이트로 이루어져 있으며 어떤 길이의 slice길이를 호출하더라도 항상 24바이트만 복사됩니다. 주소값이 복사됨으로 slice2와 slice는 같은 인스턴스를 가르킵니다
//따라서 slice2를 바꾸면 slice도 바뀌게 됩니다.

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	changeArray(array)
	changeSlice(slice)

	fmt.Println("array:", array) // 1, 2, 3, 4, 5
	fmt.Println("slice:", slice) // 1, 2, 200, 4, 5
}
