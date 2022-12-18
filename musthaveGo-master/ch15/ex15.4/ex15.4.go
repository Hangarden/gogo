// ch15/ex15.4/ex15.4.go
package main

import "fmt"

func main() {
	str1 := "가나다라마"  // ❶ 한글 문자열
	str2 := "abcdeF" // ❷ 영문 문자열

	fmt.Printf("len(str1) = %d\n", len(str1)) // 한글 문자열 크기 한글은 한 글자당 3바이트
	fmt.Printf("len(str2) = %d\n", len(str2)) // 영문 문자열 크기 영문은 한 글자당 1바이트, 대문자도 마찬가지
}
