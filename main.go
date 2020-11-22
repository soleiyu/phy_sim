package main

import "./lib"

func main() {

	vx := 2.0

	p1 := lib.NewParticle(-2, 0, vx, 0, 1)
	p2 := lib.NewParticle(-2, 2, vx, 0, 1)
	p3 := lib.NewParticle(-2, 4, vx, 0, 1)
	p4 := lib.NewParticle(-2, -2, vx, 0, 1)
	p5 := lib.NewParticle(-2, -4, vx, 0, 1)
	p6 := lib.NewParticle(-4, 3, vx, 0, 1)
	p7 := lib.NewParticle(-4, 1, vx, 0, 1)
	p8 := lib.NewParticle(-4, -1, vx, 0, 1)
	p9 := lib.NewParticle(-4, -3, vx, 0, 1)
	m1 := lib.NewMaterial(0, 0)
	m2 := lib.NewMaterial(4, -4)

	ps := []lib.Particle{p1, p2, p3, p4, p5, p6, p7, p8, p9}
	ms := []lib.Material{m1, m2}

	lib.PrintCoodPsMs(ps, ms)

	for i := 0; i < 10000; i++ {
		ps = lib.NextPsMsYg(ps, ms, 0.002, 5.0, -5.0)
		lib.PrintCoodPsMs(ps, ms)
	}
}
