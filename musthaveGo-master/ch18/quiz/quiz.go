package main

import (
	"fmt"
	"sort"
)

type Player struct {
	name         string
	age          int
	scores       int
	passAccuracy float64
}

type Players []Player

func (p Players) Less(i, j int) bool {
	return p[i].scores < p[j].scores
}
func (p Players) Swap(i, j int) {
	p[j], p[i] = p[i], p[j]
}
func (p Players) Len() int {
	return len(p)
}

func main() {

	p := []Player{{"나동키", 13, 45, 78.4},
		{"오맹태", 16, 24, 67.4},
		{"오동도", 18, 54, 50.8},
		{"황금산", 16, 36, 89.7},
	}
	sort.Sort(sort.Reverse(Players(p)))
	fmt.Println(p)
	//slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//fmt.Println(slice[0 : len(slice)-2])

}
