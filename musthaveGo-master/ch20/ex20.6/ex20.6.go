// ch20/ex20.6/ex20.6.go
package main

type Attacker interface {
	Attack()
}

func main() {
	var att Attacker // ❶ 인터페이스의 기본값은 nil입니다.
	att.Attack()     // ❷ att(interface)가 nil이기 때문에 런타임 에러가 발생합니다.
	//문법에는 오류가 없으나 실행중에 생기는 에러를 런타임 에러라고 한다.
}
