// ch19/ex19.1/ex19.1.go
package main

import "fmt"

type account struct {
	balance int
}

func withdrawFunc(a *account, amount int) { // 일반 함수 표현 *를 사용하지 않으면 인수만 값이 바뀜
	a.balance -= amount
}

func (a *account) withdrawMethod(amount int) { // 메서드 표현 메서드 명 앞에 붙으면 리시버라고 함.❶
	a.balance -= amount //withdrawMethod는 account타입에 속한다.
}

func main() {
	a := &account{100} // balance가 100인 account 포인터 변수 생성

	withdrawFunc(a, 30) // 함수 형태 호출

	a.withdrawMethod(30) // 메서드 형태 호출 ❷

	fmt.Printf("%d \n", a.balance)
}
