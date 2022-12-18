// ch21/ex21.3/ex21.3.go
package main

import "fmt"

func add(a, b int) int { //함수 포인터는 func(int, int)로 표현가능
	return a + b
}

func mul(a, b int) int { //함수 포인터는 func(int, int)로 표현가능
	return a * b
}

func getOperator(op string) DoubleFunction { // ❶ op에 따른 함수 타입 반환 //int타입 인수 2개를 받고 int타입을 반환하는 함수 타입.
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
	} else { // ❷ +나 *가 아니면 nil 반환
		return nil
	}

	// 1. op 값을 통해 func(int, int) 가 +인지, *인지 판단한다.
	// 2. 판단된다면 add, mul중 하나 선택.
	// 3. 선택 된다면 해당 함수로 인자 값들을 계산
}

type DoubleFunction func(int, int) int

func main() {
	// ❸ int 타입 인수 2개를 받아서 int 타입 반환을 하는 함수 타입 변수
	var operator DoubleFunction //초기화를 안 했기에 기본값인 nil을 가진다.
	operator = getOperator("*") //getOperator("*") -> mul()을 반환한다.

	var result = operator(3, 4) // ❹ 함수 타입 변수를 사용해서 함수 호출 -> mul(3,4)
	fmt.Println(result)
}
