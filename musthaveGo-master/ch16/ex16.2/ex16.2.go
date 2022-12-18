// ch16/ex16.2/ex16.2.go
package main

import (
	"fmt"

	"ch16/ex16.2/publicpkg"
)

func main() {
	fmt.Println("PI:", publicpkg.PI)
	//fmt.Println("pi:", publicpkg.pi) //접근 불가
	publicpkg.PublicFunc()

	var myint publicpkg.MyInt = 10
	//var myint2 publicpkg.myString = "10"
	fmt.Println("myint:", myint)

	var mystruct = publicpkg.MyStruct{Age: 18}
	//var mystruct = publicpkg.MyStruct{Age: 18, name:"garden"} 접근 불가
	fmt.Println("mystruct:", mystruct)
	var s1 = new(publicpkg.MyStruct)
	s1.PublicMethod()
	//s1.privateFunc()

}
