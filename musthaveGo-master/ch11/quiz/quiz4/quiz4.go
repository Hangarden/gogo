package main

import "fmt"

func main() {
	for j := 5; j > 1; j-- {
		for i := 0; i < j; i++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
	//for j := 0; j < 5; j++ {
	//	for i := 0; i < 5-j; i++ {
	//		fmt.Print("*")
	//	}
	//	fmt.Println()
	//}
}
