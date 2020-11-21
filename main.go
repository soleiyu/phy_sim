package main

import "./lib"

func main() {

	p1 := lib.NewParticle(0, 0, 1, 0, 1)
	m1 := lib.NewMaterial(1, 1)
	m2 := lib.NewMaterial(4, -4)
	m3 := lib.NewMaterial(1, -5)

	ms := []lib.Material{m1, m2, m3}

	lib.PrintCoodPMs(p1, ms)

	for i := 0; i < 10000; i++ {
		p1 = lib.NextPMs(p1, ms, 0.01)
		lib.PrintCoodPMs(p1, ms)
	}


}
