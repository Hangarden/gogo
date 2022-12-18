// ch18/ex18.7/ex18.7.go
package main

import "fmt"

func main() {
	array := [7]int{1, 2, 3, 4, 5, 6, 7}

	slice := array[1:2]
	slice2 := array[2:3:4] // ❶ 슬라이싱 [2,3] [시작인덱스 : 끝 인덱스 : 최대 인덱스]

	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))    //[2], 1, 5? 왜 4? slice는 앞서 말한대로 pointer(메모리 주소까지 가져온다) 6 - 2 + 1 = 5
	fmt.Println("slice:", slice2, len(slice2), cap(slice2)) //[2], 1, 1? 왜 4? slice는 앞서 말한대로 pointer(메모리 주소까지 가져온다)
	//cap이 4인 이유는 array를 슬라이싱 할 때, cap의 길이는 array인덱스 1에서부터 배열 마지막 인덱스까지 길이를 갖게 됨. array 요소 5개 갖는 배열 인덱스 1부터 마지막 배열을 사용함으로
	//cap은 4 append할 수 있는 길이는 cap - len이고 추가할 수 있는 배열의 길이는 마지막 인덱스 순서 7 -  시작 인덱스 순서 2로 = 5입니다
	array[1] = 100 // ❷ array의 두 번째 값 변경

	fmt.Println("After change second element")
	fmt.Println("array:", array)                         // [1, 100, 3, 4, 5]
	fmt.Println("slice:", slice, len(slice), cap(slice)) //그대로 X [100]1, 4

	slice = append(slice, 500) // ❸ slice에 값 추가

	fmt.Println("After append 500")
	fmt.Println("array:", array)                         //그대로
	fmt.Println("slice:", slice, len(slice), cap(slice)) //[2,3,100], 3, 3   [100 500], 2, 4
	slice = append(slice, 5000)                          // ❸ slice에 값 추가

	fmt.Println("slice:", slice, len(slice), cap(slice)) //[2,3,100], 3, 3   [100 500], 2, 4

}
