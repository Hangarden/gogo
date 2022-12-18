// ch19/ex19.2/ex19.2.go
package main

import "fmt"

//메서드의 사용이유 -> 소속 함수는 어디에도 속하지 않지만 메서드는 리시버에 속함. a.메서드명, b.메서드명 etc..

//결합도를 낮추고 응집도를 높여야 좋은 프로그래밍. 메서드는 관련 기능을 묶기에 응집도를 높이는 역할

// ❶ 사용자 정의 별칭 타입
type myInt int

// ❷ myInt 별칭 타입을 리시버로 갖는 메서드
func (a myInt) add(b int) int {
	return int(a) + b
}

func main() {
	var a myInt = 10       // myInt 타입 변수
	fmt.Println(a.add(30)) // ❸ myInt 타입의 add() 메서드 호출
	var b int = 20
	fmt.Println(myInt(b).add(50)) // ❹ int 타입을 타입변환
}
