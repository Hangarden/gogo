// ch11/ex11.4/ex11.4.go
package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		for i := 0; i < 5; i++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
