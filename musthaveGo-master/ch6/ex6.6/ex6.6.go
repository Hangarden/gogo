// ch6/ex6.6/ex6.6.go
package main

import "fmt"

func main() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Printf("%f + %f == %f : %v\n", a, b, c, a+b == c) // ❶ float 연산은 다음과 같은 오차를 가지고 있습니다.
	fmt.Println(a + b)                                    // ❷
}
