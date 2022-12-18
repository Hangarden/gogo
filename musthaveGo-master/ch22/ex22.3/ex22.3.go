// ch22/ex22.3/ex22.3.go
package main

//Queue는 선입선출 Stack 선입후출
import (
	"container/list"
	"fmt"
)

type Stack struct {
	v *list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func (s *Stack) Push(val interface{}) {
	s.v.PushBack(val) // ❶ 맨 뒤에 요소 추가 (Queue와 동일)
}

func (s *Stack) Pop() interface{} { // Stack과 Queue의 차이는 Pop에서 맨 앞을 빼느냐, 맨 뒤를 빼느냐
	back := s.v.Back() // ❷ 맨 뒤에서 요소를 반환
	if back != nil {
		return s.v.Remove(back)
	}
	return nil
}

func main() {
	stack := NewStack()

	for i := 1; i < 5; i++ {
		stack.Push(i)
	}

	val := stack.Pop()
	for val != nil {
		fmt.Printf("%v -> ", val)
		val = stack.Pop()
	}
}
