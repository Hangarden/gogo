// ch7/ex7.3/ex7.3.go
package main

import "fmt"

type Student struct {
	name    string
	math    int
	eng     int
	history int
}

func printAvgScore(student Student) { // ❶ 타입은  main.Student
	total := student.history + student.math + student.eng
	//fmt.Printf("Type is %T", student)
	avg := total / 3
	fmt.Println(student.name, "님 평균 점수는", avg, "입니다.")
}

func main() {
	student := []Student{
		{"김일동", 80, 74, 95},
		{"송이동", 88, 92, 53},
		{"박삼동", 78, 73, 78},
	}
	for i := 0; i < 3; i++ {
		printAvgScore(student[i])
	}
	//printAvgScore("김일등", 80, 74, 95) // ❷
	//printAvgScore("송이등", 88, 92, 53)
	//printAvgScore("박삼등", 78, 73, 78)
}
