package main

import (
	"encoding/json/v2"
	"fmt"
	"os"
)

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

type Player struct {
	Health int
}

type Counter struct {
	count int
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (p *Player) TakeDamage(amount int) {
	p.Health -= amount
}

func (c *Counter) Increment() {
	c.count++
}

func main() {
	/*
		i := 1
		fmt.Println("initial:", i)

		zeroval(i)
		fmt.Println("zeroval:", i)

		zeroptr(&i)
		fmt.Println("zeroptr:", i)

		fmt.Println("pointer:", &i)

		p := new(42)
		fmt.Println("value at *p:", *p)
		zeroptr(p)
		fmt.Println("value at *p:", *p)
	*/
	for x := range 37 {
		fmt.Print("-")
		if x == 36 {
			fmt.Println()
		}
		x++
	}

	hero := Player{Health: 100}
	hero.TakeDamage(30)
	fmt.Println("Hero Health:", hero.Health)

	for x := range 37 {
		fmt.Print("-")
		if x == 36 {
			fmt.Println()
		}
		x++
	}
	myCounter := Counter{count: 0}
	myCounter.Increment()
	myCounter.Increment()
	fmt.Println(myCounter.count)

	for x := range 37 {
		fmt.Print("-")
		if x == 36 {
			fmt.Println()
		}
		x++
	}

	jsonData := []byte(`{"name": "Abdullah", "age": 27}`)
	var u User

	err := json.Unmarshal(jsonData, &u)
	if err != nil {
		fmt.Println("JSON Error:", err)
		return
	}

	fmt.Printf("%s is %d\n", u.Name, u.Age)

	data, _ := json.Marshal(u)
	_ = os.WriteFile("user.json", data, 0644)
}
