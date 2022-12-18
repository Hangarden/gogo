// ch20/ex20.7/ex20.7.go
package main

import "fmt"

type Stringer interface { // ❶ 인터페이스
	String() string
}

type Student struct { // ❷ 구조체
	Age int
}

func (s *Student) String() string { // ❸ Student 타입의 String() 메서드 Student 구조체의 메서드
	return fmt.Sprintf("Student Age:%d", s.Age) //문자열로 변환 출력은 아님
}

func PrintAge(stringer Stringer) { //Student 구조체는 String() 메서드를 가지고 있기에 Stringer 인터페이스로 사용가능 ➍ //함수

	s := stringer.(*Student)       //stringer를 *Student로 타입 변환.  //➎ *Student 타입으로 타입 변환
	fmt.Printf("Age: %d\n", s.Age) // ➏ s.Age 출력 //

}

func main() {
	s := &Student{15} // ➐ *Student 타입 변수 s 선언 및 초기화 구조체 선언

	PrintAge(s) // ➑ 변수 s 를 인터페이스 인수로 PrintAge() 함수 호출 //printAge는 인수로 stringer를 받지만 Student가 String()을 메서드로 포함하고 있기에 사용 가능.
}
