package main

import "fmt"

type OpFunc func(int, int) int //함수 리터럴
func Process(a, b int, op OpFunc) {
	fmt.Println("Result : ", op(a, b))
}

func main() {
	op := func(a, b int) int { //람다
		return a * b
	}

	Process(5, 6, op) //30
}
