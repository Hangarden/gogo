package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		for j := 1; j < 10; j++ {
			if i == j {
				fmt.Printf("%d * %d = %d \n", i, j, i*j)
			}
		}
	}
	//1 * 1= 1
	//3 * 3 = 9
	//...
	//9 * 9 = 81
}
