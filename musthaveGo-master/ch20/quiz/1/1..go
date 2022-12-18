package main

func main() {
	// 0. ReadWriter 인터페이스 선언 Read(), Write() 메서드 포함
	// 1. File 구조체
	//2. File 구조체 Read() 메서드 포함
	//3. ReadWrite함수 선언, 인수 ReadWriter인터페이스 받음. Read(), Write() 함수 호출
	//4. File구조체 생성
	//5. ReadWrite에 File을 매개변수로 사용
	//6. 불가능 File 구조체는 Read() 만 가지고 있음으로 ReadWriter 인터페이스로 대체 불가능

	// 해결 방법 : File 에 Write() 메서드 추가
}
