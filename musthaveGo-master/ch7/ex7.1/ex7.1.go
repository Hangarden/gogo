// ch7/ex7.1/ex7.1.go
package main

import "fmt"

func Add(a int, b int) int { // ❶ 함수선언 정수 인수 a,b를 입력받고 리턴값은 int
	return a + b // ❷ a + b = 9 a,b 는 ㅎ마ㅜㅅ가 종료되면 사라진다.
}

func main() {
	c := Add(3, 6) // ❸ 매개변수를 인수로 복사함
	fmt.Println(c) // ➍
}
