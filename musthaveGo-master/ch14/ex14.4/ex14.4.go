// ch14/ex14.4/ex14.4.go
package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

func ChangeData(arg *Data) { // ❶ 파라미터로 Data 포인터를 받습니다.
	arg.value = 999     // ❸ arg 데이터 변경
	arg.data[100] = 999 //포인터로 받기에 arg와 data는 같은 메모리 주소를 공유합니다.
}

func main() {
	var data Data
	var user *Data = &Data{} //해당 방식으로 변수선언과 초기화를 동시에도 가능합니다.
	var user1 = &Data{}
	//더 간단한 방식
	var p2 = new(Data)
	fmt.Println(user, user1, p2)
	ChangeData(&data)                      // ❷ 인수로 data의 주소를 넘깁니다.
	fmt.Printf("value = %d\n", data.value) // ❹ data 필드값 출력
	fmt.Printf("data[100] = %d\n", data.data[100])
}
