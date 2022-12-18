// ch18/ex18.11/ex18.11.go
package main

import "fmt"

func main() {

	//1. 위치와 값이 모두 들어가야 함.
	//2. 들어가려면 cap, len을 늘려야 하는데 얼마나 늘어나도 되는지는 상관없음, append로 cap을 2배 늘리고 len도 늘리자.
	//3. 맨 끝 인덱스에 추가하고 추가하려는 인덱스 까지 값들을 이동시킨다.
	//4. 2번 인덱스에 2값을 삽입하고자 한다면
	//
	//5. 1, 2, 3, 4, 5, 6, 2
	//
	//slice[2] = 3 이 slice[3]으로 움직여야 한다.

	//방법1
	//slice := []int{1, 2, 3, 4, 5, 6}
	//idx, val := 2, 2
	//slice = append(slice[:idx], append([]int{val}, slice[idx:]...)...)
	//fmt.Println(slice)

	//방법2
	//slice := []int{1, 2, 3, 4, 5, 6}
	//idx, val := 2, 2
	//slice = append(slice, val)
	//for i := len(slice) - 1; i > idx; i-- {
	//	slice[i] = slice[i-1]
	//}
	//slice[idx] = val
	//fmt.Println(slice)

	//방법3
	slice := []int{1, 2, 3, 4, 5, 6}
	idx, val := 0, 0
	slice = append(slice, 0)
	copy(slice[idx+1:], slice[idx:])
	slice[idx] = val //값을 할당하는 것이 나은가 그 자리에서 선언시키는 것이 나은가 그 자리에서 선언 시키는 것이 메모리상으로는 효율적일 수 있으나, 코드의 가독성은 위가 나을듯
	fmt.Println(slice)
	//fmt.Println(slice, idx)

	//val := 3
	//pos := 3
	//slice = append(slice, val)
	//fmt.Println(slice)
	////lastIndex := len(slice)
	//slice2 := append(slice[:pos], val)
	//fmt.Println(slice2)
	//slice2:
	//	append(slice2)
	//slice := append(slice2, slice[pos:lastIndex])

	//1. 추가하려는 인덱스, 값 설정
	//idx := 3
	////val := 3
	////그전 값은 그대로 있어도 됨
	//fmt.Println(slice[:idx])
	//for i := idx; i < len(slice)+1; i++ {
	//	slice[i+1] = slice[i]
	//}
	//
	//fmt.Println(slice)
	//2. 추가하려는 인덱스 뒤에 값들을 한칸씩 밀어야 됨. 처음에 추가하려면 [1] = [0], [2] = [1], [마지막 + 1] = [마지막]

	//3. 마지막 인덱스는 len() - 1 임으로 마지막 인덱스는 len으로 표현될 수 있도록 하자 그러면 len + 1까지 for문으로 표현하면 된다.

	//
	//// ❶ 맨 뒤에 요소 추가
	//slice = append(slice, 0)
	//
	//idx := 2 // 추가하려는 위치
	//
	//// ❷ 맨 뒤부터 추가하려는 위치까지 값을 하나씩 옮겨줍니다.
	//for i := len(slice) - 2; i >= idx; i-- {
	//	slice[i+1] = slice[i]
	//}
	//
	//// ❸ 값 변경
	//slice[idx] = 100
	//
	//fmt.Println(slice)
}
