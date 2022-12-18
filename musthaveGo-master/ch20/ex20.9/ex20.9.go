// ch20/ex20.9/ex20.9.go
package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
}

func (s *Student) String() string { //stdent구조체는 String메서드를 가짐
	return "Student"
}

type Actor struct {
}

func (a *Actor) String() string { //Actor 구조체는 String메서드를 가짐
	return "Actor"
}

func ConvertType(stringer Stringer) { // stringer 인터페이스를 인수로 받음 stringer인터페이스는 String()메서드를 포함
	// Actor, Student모두 사용 가능
	// ❷ 런타임 에러 발생: *Student 타입은 Stringer 인터페이스로 쓰일 수 있지만
	// stringer값이 *Student 타입이 아니기 때문에 에러가 발생합니다.
	student := stringer.(*Student)
	fmt.Println(student)
}

func main() {
	// ❶ *Actor 구조체 값을 ConvertType() 함수의 인수로 사용합니다.
	//actor := &Actor{} //Actor를 넣으면 stringer 인터페이스 안에 있는 String() 메서드 내부에서는 인자로 *Actor를 가리키고 있다.
	//ConvertType(actor) //따라서 ConverType실행시 *Student가 아님으로 에러 발생
	student := &Student{}
	ConvertType(student)
}
