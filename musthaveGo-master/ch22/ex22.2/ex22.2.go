// ch22/ex22.2/ex22.2.go
package main

import (
	"container/list"
	"fmt"
)

type Queue struct { // ❶ Queue 구조체 정의
	v *list.List //내부 필드를 리스트로 가지고 있어 추가 삭제 할 시 리스트 사용
}

func (q *Queue) Push(val interface{}) { // ❷ 요소 추가
	q.v.PushBack(val) //Queue의 Push(요소 추가) method 추가. 맨 뒤에 추가
}

func (q *Queue) Pop() interface{} { // ❸ 요소을 반환하면서 삭제
	front := q.v.Front() // 맨 앞에
	if front != nil {    //맨 앞에가 nil이 아니라면
		return q.v.Remove(front) //remove(맨앞)
	}
	return nil //맨 앞이 nil이라면 nil리턴
}

func NewQueue() *Queue {
	return &Queue{list.New()} //Queue 생성
}

func main() {
	queue := NewQueue() // ❹ 새로운 큐 생성

	for i := 1; i < 5; i++ { // ➎ 요소 입력
		queue.Push(i) // 1, 2, 3, 4, 5
	}
	v := queue.Pop()
	for v != nil { // ➏ 요소 출력
		fmt.Printf("%v -> ", v) //nil이 아니면 출력
		v = queue.Pop()         //for를 통하여 반복하다 더 이상 pop할 게 없으면 nil이 되어 종료 !
	}
}
