package lib

import "math"

func NextPM (p Particle, m Material, dura float64) Particle {
	dsq := math.Pow(p.X - m.X, 2.0) + math.Pow(p.Y - m.Y, 2.0)
	f := dura / (dsq * p.M)

	vmpx := p.X - m.X
	vmpy := p.Y - m.Y

	p.X += p.Vx * dura
	p.Y += p.Vy * dura

	p.Vx += vmpx * f
	p.Vy += vmpy * f

	return p
}

func NextPMs (p Particle, m []Material, dura float64) Particle {
	nx := p.X + p.Vx * dura
	ny := p.Y + p.Vy * dura

	for i := 0; i < len(m); i++ {
		dsq := math.Pow(p.X - m[i].X, 2.0) + math.Pow(p.Y - m[i].Y, 2.0)
		f := dura / (dsq * p.M)

		vmpx := p.X - m[i].X
		vmpy := p.Y - m[i].Y

		p.X += p.Vx * dura
		p.Y += p.Vy * dura

		p.Vx += vmpx * f
		p.Vy += vmpy * f
	}

	p.X = nx
	p.Y = ny

	return p
}



func NewPos (inp Particle, dura float64) Particle {
	inp.X += inp.Vx * dura
	inp.Y += inp.Vy * dura

	return inp
}


