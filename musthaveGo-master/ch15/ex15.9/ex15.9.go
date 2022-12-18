// ch15/ex15.9/ex15.9.go
package main

import "fmt"

func main() {
	str := "Hello 월드!"      // ❶ 한영 문자가 섞인 문자열
	for _, v := range str { // ❷ range를 이용한 순회
		fmt.Printf("타입:%T 값:%d 문자:%c\n", v, v, v) // ❸ 출력
	}

	str2 := "뭐라카노"
	for _, i := range str2 {
		fmt.Printf("타입:%T, 값:%d, 문자: %c \n", i, i, i)
	}
	str3 := "rune도 사용가능"
	arr2 := []rune(str3)
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("타입 : %T, 값: %d, 문자: %c\n", arr2[i], arr2[i], arr2[i])
	}
}
