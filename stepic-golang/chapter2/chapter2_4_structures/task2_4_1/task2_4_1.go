package main

type Test struct {
	On          bool
	Ammo, Power int
}

func (t *Test) Shoot() bool {
	if !t.On {
		return false
	}
	if t.Ammo != 0 {
		t.Ammo--
		return true
	}
	return false

}

func (t *Test) RideBike() bool {
	if !t.On {
		return false
	}
	if t.Power != 0 {
		t.Power--
		return true
	}
	return false
}

func main() {
	testStruct := new(Test)
	testStruct.Shoot()
	testStruct.RideBike()
}
