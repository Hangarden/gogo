// ch13/ex13.8/ex13.8.go
package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	A int8 // 1바이트
	C int8 // 1바이트
	E int8 // 1바이트 배열은 같은 타입을 묶어 응집도를 높이고 공간을 효율적으로 활용할 수 있도록 한다.
	B int  // 8바이트
	D int  // 8바이트
}

func main() {
	user := User{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(user))
}
