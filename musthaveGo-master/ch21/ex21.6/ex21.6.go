// ch21/ex21.6/ex21.6.go
package main

import "fmt"

func CaptureLoop() {
	f := make([]func(), 3) //함수 리터럴(타입, 21.4참조) 3개를 가진 슬라이스
	fmt.Println("CaptureLoop")
	for i := 0; i < 3; i++ {
		f[i] = func() { //f[0] = func() : { 터미널에 0 출력 하는 함수},
			fmt.Println(i) //i변수가 참조로 캡처 되기 때문에 함수 리터럴이 호출되는 순간 i의 값은 3 값 복사가 아닌 인스턴스 복사임으로 다 3으로
		}
	}

	for i := 0; i < 3; i++ {
		f[i]() //함수들 호출해서 찍어 내도록
	}
}

func CaptureLoop2() {
	f := make([]func(), 3)      // 3개의 slice생성
	fmt.Println("CaptureLoop2") //CaptureLoop2
	for i := 0; i < 3; i++ {    //0 , 1 , 2
		v := i          //i변수를 저장하는 v변수를 선언하여 내부 변수 임으로
		f[i] = func() { //캡처하면 된다.
			fmt.Println(v) // 캡처 되는 순간 v값들이 캡쳐되어 0,1,2 찍힘
		}
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func main() {
	CaptureLoop()
	CaptureLoop2()
}
