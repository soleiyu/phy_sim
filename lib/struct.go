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

//num width ofset
func NewParticleArray(xn, yn int, xw, yw, xo, vx float64) []Particle {
	res := make([]Particle, xn * yn)

	xd := 2.0 * xw / float64(xn - 1)
	yd := 2.0 * yw / float64(yn - 1)

	for x := 0; x < xn; x++ {
		for y := 0; y < yn; y++ {
			res[yn * x + y] =	
				NewParticle(xo + xw - float64(x) * xd, yw - float64(y) * yd, vx, 0, 1)
		}
	}

	return res;
}

func NewParticleArray5x9 (vx float64) []Particle {
	res := make([]Particle, 45)

	res[0] = NewParticle(-2, 4, vx, 0, 1)
	res[1] = NewParticle(-2, 3, vx, 0, 1)
	res[2] = NewParticle(-2, 2, vx, 0, 1)
	res[3] = NewParticle(-2, 1, vx, 0, 1)
	res[4] = NewParticle(-2, 0, vx, 0, 1)
	res[5] = NewParticle(-2, -1, vx, 0, 1)
	res[6] = NewParticle(-2, -2, vx, 0, 1)
	res[7] = NewParticle(-2, -3, vx, 0, 1)
	res[8] = NewParticle(-2, -4, vx, 0, 1)

	res[9] = NewParticle(-3, 4, vx, 0, 1)
	res[10] = NewParticle(-3, 3, vx, 0, 1)
	res[11] = NewParticle(-3, 2, vx, 0, 1)
	res[12] = NewParticle(-3, 1, vx, 0, 1)
	res[13] = NewParticle(-3, 0, vx, 0, 1)
	res[14] = NewParticle(-3, -1, vx, 0, 1)
	res[15] = NewParticle(-3, -2, vx, 0, 1)
	res[16] = NewParticle(-3, -3, vx, 0, 1)
	res[17] = NewParticle(-3, -4, vx, 0, 1)

	res[18] = NewParticle(-4, 4, vx, 0, 1)
	res[19] = NewParticle(-4, 3, vx, 0, 1)
	res[20] = NewParticle(-4, 2, vx, 0, 1)
	res[21] = NewParticle(-4, 1, vx, 0, 1)
	res[22] = NewParticle(-4, 0, vx, 0, 1)
	res[23] = NewParticle(-4, -1, vx, 0, 1)
	res[24] = NewParticle(-4, -2, vx, 0, 1)
	res[25] = NewParticle(-4, -3, vx, 0, 1)
	res[26] = NewParticle(-4, -4, vx, 0, 1)

	res[27] = NewParticle(-5, 4, vx, 0, 1)
	res[28] = NewParticle(-5, 3, vx, 0, 1)
	res[29] = NewParticle(-5, 2, vx, 0, 1)
	res[30] = NewParticle(-5, 1, vx, 0, 1)
	res[31] = NewParticle(-5, 0, vx, 0, 1)
	res[32] = NewParticle(-5, -1, vx, 0, 1)
	res[33] = NewParticle(-5, -2, vx, 0, 1)
	res[34] = NewParticle(-5, -3, vx, 0, 1)
	res[35] = NewParticle(-5, -4, vx, 0, 1)

	res[36] = NewParticle(-6, 4, vx, 0, 1)
	res[37] = NewParticle(-6, 3, vx, 0, 1)
	res[38] = NewParticle(-6, 2, vx, 0, 1)
	res[39] = NewParticle(-6, 1, vx, 0, 1)
	res[40] = NewParticle(-6, 0, vx, 0, 1)
	res[41] = NewParticle(-6, -1, vx, 0, 1)
	res[42] = NewParticle(-6, -2, vx, 0, 1)
	res[43] = NewParticle(-6, -3, vx, 0, 1)
	res[44] = NewParticle(-6, -4, vx, 0, 1)

	return res
}


func NewParticleArray3x9 (vx float64) []Particle {
	res := make([]Particle, 27)

	res[0] = NewParticle(-2, 4, vx, 0, 1)
	res[1] = NewParticle(-2, 3, vx, 0, 1)
	res[2] = NewParticle(-2, 2, vx, 0, 1)
	res[3] = NewParticle(-2, 1, vx, 0, 1)
	res[4] = NewParticle(-2, 0, vx, 0, 1)
	res[5] = NewParticle(-2, -1, vx, 0, 1)
	res[6] = NewParticle(-2, -2, vx, 0, 1)
	res[7] = NewParticle(-2, -3, vx, 0, 1)
	res[8] = NewParticle(-2, -4, vx, 0, 1)

	res[9] = NewParticle(-3, 4, vx, 0, 1)
	res[10] = NewParticle(-3, 3, vx, 0, 1)
	res[11] = NewParticle(-3, 2, vx, 0, 1)
	res[12] = NewParticle(-3, 1, vx, 0, 1)
	res[13] = NewParticle(-3, 0, vx, 0, 1)
	res[14] = NewParticle(-3, -1, vx, 0, 1)
	res[15] = NewParticle(-3, -2, vx, 0, 1)
	res[16] = NewParticle(-3, -3, vx, 0, 1)
	res[17] = NewParticle(-3, -4, vx, 0, 1)

	res[18] = NewParticle(-4, 4, vx, 0, 1)
	res[19] = NewParticle(-4, 3, vx, 0, 1)
	res[20] = NewParticle(-4, 2, vx, 0, 1)
	res[21] = NewParticle(-4, 1, vx, 0, 1)
	res[22] = NewParticle(-4, 0, vx, 0, 1)
	res[23] = NewParticle(-4, -1, vx, 0, 1)
	res[24] = NewParticle(-4, -2, vx, 0, 1)
	res[25] = NewParticle(-4, -3, vx, 0, 1)
	res[26] = NewParticle(-4, -4, vx, 0, 1)

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

func PrintCoodPsMs(p []Particle, m []Material) {
	for i := 0; i < len(p); i++ {
		fmt.Printf("%f, %f, ", p[i].X, p[i].Y)
	}

	for i := 0; i < len(m); i++ {
		fmt.Printf("%f, %f, ", m[i].X, m[i].Y)
	}
	fmt.Println()
}
