package main

import "fmt"

func AAA() {
	fmt.Println("start AAA()") //1
	BBB()                      //2호출 -> 출력
	fmt.Println("end AAA()")   //출력
}
func BBB() {
	fmt.Println("BBB()")
}

func main() {
	AAA()

}
