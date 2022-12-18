//ch11/ex11.2/ex11.2.go
package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for {
		time.Sleep(time.Second) //1초마다
		fmt.Println(i)          //프린트하고
		i++                     //다시 올리고
	}
}
