// ch21/ex21.2/ex21.2.go
package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Create("test.txt") // ❶ 파일 생성
	if err != nil {                 // ❷ 에러 확인
		fmt.Println("Failed to create a file")
		return
	}

	defer fmt.Println("반드시 호출됩니다.") // ❸ 지연 수행될 코드 //파일, 소켓 처럼 OS내부 자원을 사용할 시에는 함수가 종료되기 직전에 사용되어야함
	// 그렇지 않으면 내부 자원이 고갈되어 네트워크 통신이 불가능 하거나, 파일을 사용하지 못할 수도
	defer f.Close()                // ➍ 지연 수행될 코드
	defer fmt.Println("파일을 닫았습니다") // ➎ 지연 수행될 코드

	fmt.Println("파일에 Hello World를 씁니다.")
	fmt.Fprintln(f, "Hello World") // ➏ 파일에 텍스트를 씁니다.
}
