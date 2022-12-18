// ch11/ex11.6/ex11.6.go
package main

import "fmt"

func main() {
	gugu := 2
	dan := 1
	for {
		for i := gugu; i < 10; i++ {
			for j := dan; j < 10; j++ {
				fmt.Printf("%d * %d = %d\n", i, j, i*j)
			}
			dan++
			if dan == 10 {
				break
			}
		}
	}
	fmt.Println("for문이 종료되었습니다.")
}
