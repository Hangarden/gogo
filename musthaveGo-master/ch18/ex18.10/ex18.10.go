// ch18/ex18.10/ex18.10.go
package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}

	//1. 삭제할 인덱스 설정
	idx := 2 // 삭제할 인덱스
	//2. 각 인덱스를 한칸씩앞으로 1번째를 삭제한다면 2번째 -> 1 번째, 3번째-> 2번째 마지막-> 마지막 -1번째 마지막은 그대로
	for i := idx; i < len(slice)-1; i++ { //len(slice) = 6 slice[5] = slice[4]
		slice[i] = slice[i+1]
	}
	slice = append(slice[:len(slice)-1])
	fmt.Println(slice)

	//3. 마지막 인덱스는 삭제한다.

	//for i := 0; i > 5; i++ { // 실행자체가 안된다
	//	fmt.Println("과연")
	//}
}
