// ch19/ex19.3/ex19.3.go
package main

import "fmt"

type account struct { // account 구조체 정의
	balance   int    //잔고
	firstName string //성
	lastName  string //이름
}

// 포인터 타입 메서드
func (a1 *account) withdrawPointer(amount int) {
	a1.balance -= amount //account의 메서드.
}

// 값 타입 메서드
func (a2 account) withdrawValue(amount int) {
	a2.balance -= amount
}

// 변경된 값을 반환하는 값 타입 메서드
func (a3 account) withdrawReturnValue(amount int) account {
	a3.balance -= amount
	return a3
}

func main() {
	var mainA = account{100, "Joe", "Park"} //account 구조체 선언 &를 안쓰면 어케됨?
	mainA.withdrawPointer(30)               // 30만큼 뺀다.
	fmt.Println(mainA.balance)              // 70 이 출력

	mainA.withdrawValue(20)    // 포인터 변수 값타입 메서드 호출 - ①
	fmt.Println(mainA.balance) // 여전히 70이 출력

	var mainB account = mainA.withdrawReturnValue(20)
	fmt.Println(mainB.balance) // 50이 출력

	mainB.withdrawPointer(30)  // 값 변수 포인터타입 메서드 호출 - ②
	fmt.Println(mainB.balance) // 20이 출력
}
