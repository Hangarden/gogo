package main

import "fmt"

func main() {

	for i := 2; i < 10; i++ {
		if 3 <= i && i <= 6 {
			continue
		}
		for j := 1; j < 10; j++ {
			fmt.Printf("%d * %d = %d \n", i, j, i*j)
		}
		fmt.Println("")
	}

}
