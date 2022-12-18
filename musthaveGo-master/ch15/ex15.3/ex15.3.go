// ch15/ex15.3/ex15.3.go
package main

import "fmt"

func main() { //go는 utf-8을 표준 문자코드로 사용, ANSI와 1대1대응 된다.
	var char rune = '한' //rune과 int32는 이름만 다를 뿐 같은 타입

	fmt.Printf("%T\n", char) // ❶ char 타입 출력
	fmt.Println(char)        // ❷ char값 출력
	fmt.Printf("%c\n", char) // ❸ 문자 출력
}
