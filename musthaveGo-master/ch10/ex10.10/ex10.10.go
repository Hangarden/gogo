// ch10/ex10.10/ex10.10.go
package main

import "fmt"

func main() {

	a := 3

	switch a {
	case 1:
		fmt.Println("a == 1")
		break
	case 2:
		fmt.Println("a == 2")
	case 3:
		fmt.Println("a == 3")
		fallthrough // 아래까지 같이 실행 해 버린다.
	case 4:
		fmt.Println("a == 4")
	case 5:
		fmt.Println("a == 5")
	default: //만족하는 케이스가 없을 시 실행하지만 필수는 아닙니다.
		fmt.Println("a > 4")
	}
}
