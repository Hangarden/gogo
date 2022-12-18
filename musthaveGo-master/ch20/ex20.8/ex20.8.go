// ch20/ex20.8/ex20.8.go
package main

type Stringer interface {
	String() string
}

type Student struct {
}

func main() {
	var stringer Stringer
	stringer.(*Student) // Student구조체가 String() 메서드를 가지고 있지 않기 때문.  ❶ 빌드 타임 에러 발생
}
