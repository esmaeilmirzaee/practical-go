package main

import (
	"fmt"
)

type point struct {
	X int
	Y int
}

func main() {
	game()
}

func game() {
	var p1 point
	fmt.Println(p1)
	fmt.Printf("%#v\n", p1)

	p2 := point{1, 2}
	fmt.Printf("%#v\n", p2)

	p3, _ := newPoint(10, 20)
	p3.move(20, 10)
	p3.status()
	fmt.Printf("This is type of p3: %T", *p3)

	// somePoints := []mover{
	//	&p1,
	//	&p2,
	//}
	//	moveAll(somePoints, point{0, 0})
	fmt.Println("Copper: ", Copper)
	fmt.Println("Key(20): ", Key(20))
}

type Key byte

const (
	Jade   Key = iota + 1
	Copper     = iota + 5
	Crystal
)

type player struct {
	name string
	p    point
	keys []Key
}

func (p *player) FoundKey(k Key) error {
	if k < Jade || k >= Crystal {
		return fmt.Errorf("error: Key is unknown: %s", k)
	}

	if !containKeys(p.keys, k) {
		p.keys = append(p.keys, k)
	}

	return nil
}

func containKeys(keys []Key, k Key) bool {
	for _, k2 := range keys {
		if k == k2 {
			return true
		}
	}
	return false
}

func (k Key) String() string {
	switch k {
	case Jade:
		return "jade"
	case Copper:
		return "copper"
	case Crystal:
		return "crystal"
	}
	return fmt.Sprintf("%d", k)
}

type mover interface {
	move(x, y int)
	// move(int, int) alternative
}

func moveAll(ms []mover, p point) {
	for _, np := range ms {
		np.move(p.X, p.Y)
	}
}

func (p *point) status() {
	fmt.Printf("You're in (%d, %d)\n", p.X, p.Y)
}

func (p *point) move(x, y int) {
	p.X = x
	p.Y = y
}

const (
	maxX = 1000
	maxY = 700
)

func newPoint(x, y int) (*point, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		return nil, fmt.Errorf("%d, %d out of bounds %d, %d", x, y, maxX, maxY)
	}
	return &point{
		X: x,
		Y: y,
	}, nil
}
