// ch12/ex12.6/ex12.6.go
package main

import "fmt"

func main() {
	a := [][]int{ // ❶ 이중 배열 선언 [5] 배열의 2개 있는 것 [...][...] X [][]로 사용 할 것
		{1, 2, 3, 4, 5},
		{5, 6, 7, 8, 9}, // ❷ 여러 줄에 걸쳐 초기화할 때는 쉼표를 찍자!
	}
	for _, arr := range a { // ❸ arr값은 순서대로 a[0] -> 첫번째 [5]int배열 a[1] -> 두번째 [5]int의 배열
		for _, v := range arr { // ❹ v값은 순서대로 첫번째 배열의 0,1,2,3,4 인덱스 출력
			fmt.Print(v, " ") // ➎ v값 출력
		}
		fmt.Println()
	}
}
