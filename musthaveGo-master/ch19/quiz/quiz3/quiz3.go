package main

import (
	"fmt"
	"strings"
)

type myString string

func (m myString) ToLower() myString {
	str := strings.ToLower(string(m))
	return myString(str)
}

func (m myString) ToUpper() myString {
	str := strings.ToUpper(string(m))
	return myString(str)
}

func main() {
	msg := myString("hello Go World")

	msg2 := msg.ToLower() // "hello go world"
	fmt.Println(msg2)

	msg3 := msg.ToUpper() // "HELLO GO WORLD"
	fmt.Println(msg3)

}
