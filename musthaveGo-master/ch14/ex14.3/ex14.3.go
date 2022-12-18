// ch14/ex14.3/ex14.3.go
package main

import "fmt"

type Data struct { // ❶ Data형 구조체
	value int
	data  [200]int
}

func ChangeData(arg Data) { // ❷ 파라미터로 Data를 받습니다.
	arg.value = 999
	arg.data[100] = 999 //arg와 data는 서로 다른 메모리 공간을 가지는 변수
}

func main() {
	var data Data
	//변수 대입이나 함수 인수 전달은 항상 값을 복사 -> 메모리 공간 사용
	//다른 공간으로 복사되어 변경사항 적용되지도 않음
	ChangeData(data)                               // ❸ 인수로 data를 넣습니다. //arg와 data는 서로 다른 메모리 공간을 가지는 변수
	fmt.Printf("value = %d\n", data.value)         //따라서 data.value는 초기값인 0이고
	fmt.Printf("data[100] = %d\n", data.data[100]) // ❹ data 필드 출력 //data.data[100]도 마찬가지이다.
}
