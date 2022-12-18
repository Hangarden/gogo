// ch18/ex18.13/ex18.13.go
package main

import (
	"fmt"
	"math/bits"
	"sort"
)

type Student struct {
	Name string
	Age  int
}

// ❶ []Student의 별칭 타입 Students
type Students []Student

func (s Students) Len() int {
	return len(s)
}
func (s Students) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}                                // ❷ 나이 비교
func (s Students) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func Sort(data Interface) { //아무 타입이나 올 수 있음
	n := data.Len() //여기서 Len이 사용 데이터 길이를 구하고
	if n <= 1 {
		return
	} // 데이터 가 1이하면 종료.
	limit := bits.Len(uint(n))
	pdqsort(data, 0, n, limit)
}

func main() {
	s := []Student{
		{"화랑", 31}, {"백두산", 52}, {"류", 42},
		{"켄", 38}, {"송하나", 18}}

	sort.Sort(Students(s)) // ❸ 정렬 Sort를 사용하기 위해서는 Len(), Less(), Swap 3메서드가 필요.
	fmt.Println(s)
}
