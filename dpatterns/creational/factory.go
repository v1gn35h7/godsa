package creational

import (
	"fmt"
)

type Igun interface {
	getRange() int
}

// base class
type Gun struct {
	srange int
}

func (g *Gun) getRange() int {
	return g.srange
}

// New class
type Ak47 struct {
	gun Gun
}

// New class
type Pistol struct {
	gun Gun
}

// Client code

func listRanges() {
	ak47 := &Ak47{
		gun: Gun{
			40,
		},
	}

	fmt.Println(ak47.gun.getRange())

	pistol := &Pistol{
		gun: Gun{
			100,
		},
	}

	fmt.Println(pistol.gun.getRange())

}
