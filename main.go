package main

import "./lib"

func main() {

	vx := 3.0

	m1 := lib.NewMaterial(0, 0)
	m2 := lib.NewMaterial(0.2, -0.2)

//	ps := lib.NewParticleArray5x9(vx)
	ps := lib.NewParticleArray(11, 9, 8, 4, -10, vx)
	ms := []lib.Material{m1, m2}

	lib.PrintCoodPsMs(ps, ms)

	for i := 0; i < 10000; i++ {
		ps = lib.NextPsMsYg(ps, ms, 0.002, 5.0, -5.0)
		lib.PrintCoodPsMs(ps, ms)
	}
}
