// ch5/ex5.1/ex5.1.go
package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20
	var f float64 = 32799438743.8297
	var c bool = false
	var d bool = true
	var e int = 65
	//fmt.Print("a:", a, "b:", b)               // ❶
	//fmt.Println("a:", a, "b:", b, "f:", f)    // ❷
	fmt.Printf("a: %d b: %d f:%f\n", a, b, f) // ❸

	fmt.Printf("a: %c b: %c f:%f\n", a, e, f) // ❸

	fmt.Printf("a: %v b: %v f:%v\n", a, b, f) // ❸

	fmt.Printf("a: %v b: %v f:%T\n", c, d, d) // ❸
}
