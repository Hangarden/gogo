package main

import "fmt"

type Direction int

const (
	None Direction = iota
	North
	East
	South
	West
)

func DirectionToString(d Direction) string {
	switch d {
	case North:
		return "North"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	default:
		return "None"
	}
}

func GetDirection(angle float64) Direction {
	switch {
	case 315 <= angle && 0 >= angle && angle < 45:
		return North
	case 45 <= angle && angle < 135:
		return East
	case 135 <= angle && angle <= 225:
		return South
	case 225 <= angle && angle <= 315:
		return West
	default:
		return None
	}
}
func main() {
	fmt.Println(DirectionToString(GetDirection(38.3)))
	fmt.Println(DirectionToString(GetDirection(235.8)))
	fmt.Println(DirectionToString(GetDirection(94.2)))
	fmt.Println(DirectionToString(GetDirection(-30)))
}
