// ch20/ex20.10/ex20.10.go
package main

type Reader interface {
	Read() //Read() 메서드를 가진 인터페이스 선언
}

type Closer interface {
	Close() //Close() 메서드를 가진 인터페이스 선언
}

type File struct { // File구조체 선언
	//Read() 메서드 포함
}

func (f *File) Read() { //
}

//func (f *File) Close() { //
//}

func ReadFile(reader Reader) { //Reader 인터페이스를 인수로 받음. Read() 메서드를 가진
	// ❷ Reader 인터페이스 변수를 Closer 인터페이스로 타입 변환합니다.
	// 런타임 에러가 발생합니다.
	//c, ok := reader.(Closer)
	if c, ok := reader.(Closer); ok {
		c.Close()
	} // reader인터페이스에서 Close 인터페이스로 변환. -> 불가 Close메서드가 없자타
}

func main() {
	// ❶ File 포인터 인스턴스를 ReadFile() 함수의 인수로 사용합니다.
	file := &File{} //파일 구조체 선언
	ReadFile(file)  // Readfile인수로 선언 가능하지만 안에서 reader.(Closer)에서 에러 발생
}
