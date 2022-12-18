// ch18/ex18.6/ex18.6.go
package main

import "fmt"

func main() {
	slice1 := make([]int, 5, 5) // ❶ len:3 cap:3 슬라이스 생성
	// := -> 선언 + 할당
	// = -> only assign
	for i := 0; i < 3; i++ {
		slice1[i] = i + 1
	}
	slice2 := slice1
	for i := 3; i < 5; i++ {
		slice2[i] = i + 1
	}
	slice3 := append(slice1, 4, 5) // ❷ append() 함수로 요소 추가, 빈공간이 충분한지 확인, 충분치 않다면 새로운 큰 배열 생성, 따로 생성됨으로 주소가 다름

	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2)) //4, 5를 추가할 빈 공간이 있다.
	fmt.Println("slice3:", slice3, len(slice3), cap(slice3)) // 추가로 4, 5를 추가할 빈 공간이 없음으로 새로운 변수를 만든다.

	slice1[1] = 100 // ❸ slice1 요솟값 변경, slice1, slice2는 주소값이 다름으로 변경해도 서로 아무 영향을 미치지 않음

	fmt.Println("After change second element")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1)) // slice1,2는 모두 같은 메모리 주소를 공유하기에 값이 바뀌지만 slice3는 다른 주소를 가르키기에 바뀌지 않는다.
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))
	fmt.Println("slice3:", slice3, len(slice3), cap(slice3))

	slice1 = append(slice1, 500) // ➍ slice1 요솟값 변경

	fmt.Println("After append 500")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1)) //왜 cap이 10으로 늘어났지?
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))
	fmt.Println("slice3:", slice3, len(slice3), cap(slice3))
	slice1 = append(slice1, 500) // ➍ slice1 요솟값 변경

	fmt.Println("After append 500")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1)) //왜 cap이 10으로 늘어났지?

	slice2 = []int{1, 2, 3, 4}
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	slice2 = append(slice2, 4) //append는 cap의 2배를 해주나 보다.
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

}
