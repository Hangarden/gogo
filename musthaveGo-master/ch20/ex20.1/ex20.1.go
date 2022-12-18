// ch20/ex20.1/ex20.1.go
package main

import "fmt"

type Stringer interface { // ❶ Stringer 인터페이스 선언
	String() string //매개변수 반환이 다르더라도 이름이 같은 메서드는 존재 X
	//인터페이스에서는 메서드 구현 X, 이름없는 메서드는 없다.
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string { // ❷ Student의 String() 메서드

	return fmt.Sprintf("안녕! 나는 %d살 %s라고 해", s.Age, s.Name) // ❸ 문자열 만들기 sprinf 를 사용하면 된다.
}

func main() {
	student := Student{"철수", 12} // Student 타입
	var stringer Stringer        // Stringer 타입

	stringer = student // 4 stringer값으로 student 대입

	fmt.Printf("%s\n", stringer.String()) // 5 stringer의 String() 메서드 호출 //stringer를 통해 구조체로 존재하던 student를 문자열 형태로 바꾸어 준다.
}
