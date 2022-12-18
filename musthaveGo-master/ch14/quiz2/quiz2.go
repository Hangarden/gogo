package main

import "fmt"

type Actor struct {
	Name  string
	HP    int
	Speed float64
}

func NewActor(name string, hp int, speed float64) *Actor {
	var a = Actor{name, hp, speed}
	return &a
	//retrun &Actor(name, hp, speed)
	//a := new(Actor)
	// a.Name = name, a.HP = hp, a.Speed = speed retrun a
}

func main() {
	var actor = NewActor("금토끼", 99, 100)
	fmt.Println(actor.Speed)
	fmt.Println(actor.HP)

}
