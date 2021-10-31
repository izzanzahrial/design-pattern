package main

type gun struct {
	name  string
	power int
}

func (gun *gun) setName(name string) {
	gun.name = name
}

func (gun *gun) setPower(power int) {
	gun.power = power
}

func (gun *gun) getName() string {
	return gun.name
}

func (gun *gun) getPower() int {
	return gun.power
}
