package main

import "fmt"

// 자바는 클래스 타입을 힙에, 기본타입을 스택에 할당
// Go는 탈출 검사를 해서 어느 메모리에 할당할 지 정합니다.
type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	var u = User{name, age}
	return &u // 1 탈출 분석으로 u 메모리가 사라지지 않음
}

func main() {
	userPointer := NewUser("AAA", 23) //이미 사라진 메모리를 가르키는 오류를 댕글링 오류라고 합니다.
	fmt.Println(userPointer)
}
