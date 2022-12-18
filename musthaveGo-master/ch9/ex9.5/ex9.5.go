// ch9/ex9.5/ex9.5.go
package main

import "fmt"

// 친구 중 부자가 있는가 반환 - 무조건 true 반환
func HasRichFriend() bool {
	return true
}

// 같이 간 친구 숫자를 반환 - 무조건 3을 반환
func GetFriendsCount() int {
	return 3
}

func main() {
	price := 55000

	//가격이 5만원이 넘을 때
	if price > 50000 {
		if HasRichFriend() {
			fmt.Println("신발끈을 묶는다")
		} else {
			fmt.Println("돈을 나누어 낸다, ")
		}
	} else if price >= 30000 {
		if GetFriendsCount() > 3 {
			fmt.Println("신발끈을 묶는다")
		} else {
			money := price / GetFriendsCount()
			fmt.Println("돈을 나눠낸다 :", money, " 만원 씩")
		}
	} else {
		fmt.Println("내가 다 낸다")
	}
	//
	//가격이 5만원 이하고 3만원 이상일 때
	//	1. 친구가 3명 초과일 때
	//
	//	2. 친구가 3명 이하일 때
	//3만원 미만일 때
	//
	//

	//if price > 50000 { // ❶
	//	if HasRichFriend() {
	//		fmt.Println("앗 신발끈이 풀렸네")
	//	} else {
	//		fmt.Println("나눠내자")
	//	}
	//} else if price >= 30000 && price <= 50000 { // ❷
	//	if GetFriendsCount() > 3 { // ❸
	//		fmt.Println("어이쿠 신발끈이..")
	//	} else {
	//		fmt.Println("사람도 얼마 없는데 나눠내자")
	//	}
	//} else {
	//	fmt.Println("오늘은 내가 쏜다")
	//}
}
