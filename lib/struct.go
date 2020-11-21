package lib

import "fmt"

type Particle struct {
	X, Y float64 // Position
	Vx, Vy float64 // Velocity
	M float64 // Mass
}

type Material struct {
	X, Y float64 // Position
}

func NewParticle (x, y, vx, vy, m float64) Particle {
	var res Particle

	res.X = x
	res.Y = y
	res.Vx = vx
	res.Vy = vy
	res.M = m

	return res
}

func NewMaterial (x, y float64) Material {
	var res Material

	res.X = x
	res.Y = y

	return res
}

func PrintCood(inp Particle) {
	fmt.Printf("%f, %f\n", inp.X, inp.Y)
}

func PrintCoodPM(p Particle, m Material) {
	fmt.Printf("%f, %f, %f, %f\n", p.X, p.Y, m.X, m.Y)
}

func PrintCoodPMs(p Particle, m []Material) {
	fmt.Printf("%f, %f, ", p.X, p.Y)

	for i := 0; i < len(m); i++ {
		fmt.Printf("%f, %f, ", m[i].X, m[i].Y)
	}
	fmt.Println()
}
