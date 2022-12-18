package main

import "time"

type Courier struct {
	Name string
}

type Product struct {
	Name string
	ID   int
}

type Parcel struct {
	Pdt           *Product
	ShippedTime   time.Time
	DeliveredTime time.Time
}

func (c *Courier) SendProduct(product *Product) *Parcel {
	p := &Parcel{}
	p.Pdt = product
	p.ShippedTime = time.Now()
	return p
}

func (p *Parcel) Delivered() *Product {
	p.DeliveredTime = time.Now()
	return p.Pdt
}

func main() {

}
