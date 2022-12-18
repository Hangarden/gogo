// ch8/ex8.2/ex8.2.go
package main

import "fmt"

func main() {
	const PI1 float64 = 3.141592653589793238 // ❶
	var PI2 float64 = 3.141592653589793238   // ❷

	// PI1 = 4  // ❸ 주석 풀면 cannot assign 에러 발생
	PI2 = 4 // ➍

	fmt.Printf("원주율: %f\n", PI1) //%f임으로 6자리 까지만 출력 될 것
	fmt.Printf("원주율: %f\n", PI2) // ➎
}
