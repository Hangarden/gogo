package main

import "fmt"

type ParkingLot struct {
	LotSize int
}

func ParkCar(lot *ParkingLot, carSize int) {
	lot.LotSize -= carSize
}

func (p *ParkingLot) ParkCar(carSize int) {
	p.LotSize -= carSize
}

func (p *ParkingLot) GetParkCar() int {
	return p.LotSize
}

func main() {
	lot := ParkingLot{100}
	//ParkCar(lot, 10)
	lot.ParkCar(10)
	fmt.Println("남은 주차 공간은 : ", lot.GetParkCar(), " 입니다")
}
