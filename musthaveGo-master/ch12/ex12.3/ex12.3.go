// ch12/ex12.3/ex12.3.go
package main

import "fmt"

func main() {
	nums := [...]int{10, 20, 30, 40, 50, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2} // ❶ ..을 사용하면 배열의 길이는 선언한 만큼?

	nums[2] = 300 // ❷

	for i := 0; i < len(nums); i++ { // ❸
		fmt.Println(nums[i]) // ➍
	}
	fmt.Println(len(nums))
}
