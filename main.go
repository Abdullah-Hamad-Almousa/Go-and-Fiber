package main

import "fmt"

type Player struct {
	Health int
}

type Counter struct {
	count int
}

func (p *Player) TakeDamage(amount int) {
	p.Health -= amount
}

func (c *Counter) Increment() {
	c.count++
}

func main() {

	hero := Player{Health: 100}

	fmt.Println("Hero Health:", hero.Health)

	for i := range 3 {
		i++
		hero.TakeDamage(30)
		fmt.Println("Hero Health:", hero.Health)
	}

	for x := range 37 {
		fmt.Print("-")
		if x == 36 {
			fmt.Println()
		}
		x++
	}

	myCounter := Counter{count: 0}
	//myCounter.Increment()

	//fmt.Println(myCounter.count)
	for i := range 3 {
		i++
		fmt.Printf("Counter: %v\n", myCounter.count)
		myCounter.Increment()
	}

}
