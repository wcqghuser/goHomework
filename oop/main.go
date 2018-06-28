package main

import (
	"fmt"
	"math/rand"
)

const (
	GoOn   = iota
	Left
	Right
	Arrive
)

type car struct {
	key       uint
	isStarted bool
}

func (c *car) carStar(key uint) {
	if key == c.key {
		c.isStarted = true
		fmt.Println("car started")
	}
}
func (c *car) goOn() {
	fmt.Println("go on")
}

func (c *car) turnLeft() {
	fmt.Println("turn left")
}
func (c *car) turnRight() {
	fmt.Println("turn left")
}
func (c *car) CarStop() {
	fmt.Println("arrive")
}

type directer struct {
}

func (d *directer) getNextDirect() int {
	v := rand.Intn(60)
	if v != 40 {
		v = v % 4
	}
	return v

}

type card struct {
	id uint
	amount int
}

type person struct {
	name string
	key  uint
	cardId uint
	amount int
	c    *car
	d    *directer
	cd   *card
}

func (p *person) driveCar() {
	p.c.carStar(p.key)
	if p.c.isStarted {
		for {
			v := p.d.getNextDirect()
			switch v {
			case GoOn:
				p.c.goOn()
			case Left:
				p.c.turnLeft()
			case Right:
				p.c.turnRight()
			case Arrive:
				p.c.CarStop()
				return
			}
		}
	}
}

func (p *person) withdraw(c *card, amount int) {
	if p.cardId != c.id {
		fmt.Print("it is not your card")
		return
	}
	if amount > c.amount {
		fmt.Println("card has no such amount")
		return
	}
	c.amount = c.amount - amount
	p.amount = p.amount + amount
}

func main() {
	zhangsan := person{name: "zhangsan"}
	car := car{key: 100, isStarted: false}
	zhangsan.c = &car
	zhangsan.key = 100
	d := directer{}
	zhangsan.d = &d
	cd := card{11, 50000}
	zhangsan.cd = &cd
	zhangsan.cardId = 11
	zhangsan.amount = 100

	zhangsan.driveCar()

	zhangsan.withdraw(zhangsan.cd, 1000)

	fmt.Printf("card amount is %d, person amount is %d", cd.amount, zhangsan.amount)

}
