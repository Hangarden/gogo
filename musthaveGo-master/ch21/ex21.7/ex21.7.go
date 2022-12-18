// ch21/ex21.7/ex21.7.go
package main

import (
	"fmt"
	"os"
)

type Writer func(string) // 함수 리터럴

func writeHello(writer Writer) { // ❷ writer 함수타입 변수 호출
	writer("Hello World") //writeHello 는 writer를 호출 시 그것이 파일에 쓰여질지, 프린터로 갈지 네트워크로 쏘아질지 알 수 없다 외부에서 로직을 주입하는 것을 의존성 주입이라 칸다.
}

func main() {
	f, err := os.Create("test.txt")
	if err != nil { //에러 발생시 ( 파일이 생성되지 않을 때 )
		fmt.Println("Failed to create a file")
		return
	}
	defer f.Close() //종료 직전에 실행할 값

	writeHello(func(msg string) {
		fmt.Fprintln(f, msg) // ❶ 함수 리터럴 외부 변수 f 사용
	})
}
