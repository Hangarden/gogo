// ch22/ex22.1/ex22.1.go
package main

import (
	"container/list"
	"fmt"
)

func main() {
	v := list.New()       // ❶ 새로운 리스트 생성
	e4 := v.PushBack(4)   // ❷ 리스트 뒤에 요소 추가
	e1 := v.PushFront(1)  // ❸ 리스트 앞에 요소 추가
	v.InsertBefore(3, e4) // ❹ e4 요소 앞에 요소 삽입
	v.InsertAfter(2, e1)  // ➎ e1 요소 뒤에 요소 삽입

	for e := v.Front(); e != nil; e = e.Next() { // ➏ 각 요소 순회
		fmt.Print(e.Value, " ")
	}

	fmt.Println()
	for e := v.Back(); e != nil; e = e.Prev() { // ➐ 각 요소 역순 순회
		fmt.Print(e.Value, " ")
	}

	//v.PushBack() -> 맨뒤에
	//v.PushFront() -> 맨 앞에
	//v.InsertBefore(index) -> index 앞에 요소 삽입
	//v.InsertBefore(index) -> index 뒤에 요소 삽입
	//v.Front() -> 리스트 맨 앞에
	//v.Back() -> 리스트 맨 뒤에
	//index.Next -> 다음
	//index.Prev -> 전으로

	// 배열과 인덱스의 차이점은 ? 알고리즘 Big-O로 접근 하면
	// 배열은 값을 추가하려면 한칸 씩 다 뒤로 밀어야함 O(N)
	// 리스트는 Next, Prev를 통해 삽입하면 됨으로 상수 시간 O(1)이 걸림
	// 배열은 값을 찾을 때 주소를 통해 가면 됨으로 O(1) 이지만
	// 리스트는 Prev, Next를 통해 이동해야 함으로 O(N)이다.
	// 따라서 값이 추가만 되고 사람을 추출해서 사용할 일이 적다 -> list
	// 특정 인덱스를 추출해서 사용할 일이 많다 -> array

	//컴퓨터는 연산할 때 읽어온 데이터를 캐시라는 임시 보관소에 저장한다. 필요한 데이터만 가져오는 것이 아닌 그 주변데이터도 가져오는데 그 이유는
	//높은 확률로 주변데이터에 대한 연산이 이어지기 때문, 필요한 데이터가 인접해 있을 수록 처리 속도가 빨라짐 -> sort로 하는 것이 빠를지 직접 해보도록
	//이를 데이터 지역성이 좋다고 말한다.
	//배열은 연속된 메모리로 이뤄진 자료구조
	//리스트는 불연속으로
	//배열이 리스트에 비해서 데이터 지역성이 월등하게 좋습니다. -> 데이터 규모가 작은 경우 배열이 리스트보다 더 효율적

}
