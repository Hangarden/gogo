//ch11/ex11.9/ex11.9.go
package main

import "fmt"

func find45(a int) (int, bool) { // ❶ 곱해서 45가 되는 값을 찾는다
	for b := 1; b <= 9; b++ {
		if a*b == 45 {
			return b, true
		}
	}
	return 0, false
}

func main() {
	a := 1
	b := 0

	for ; a <= 9; a++ {
		var found bool                   //bool 타입으로 선언
		if b, found = find45(a); found { // ❷ 함수 호출, 함수를 통해서 int, bool리턴 받고 found가 true라면
			break //break실행 //b = 9, a = 5
		}
	}
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}
