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

		p.Vx += vmpx * f
		p.Vy += vmpy * f
	}

	p.X = nx
	p.Y = ny

	return p
}

func NextPsMsYg (p []Particle, m []Material, dura, ypth, ynth float64) []Particle {
	pres := make([]Particle, len(p))

	kabeksu := 1.0
	pksu := 0.2

	for i := 0; i < len(p); i++ {
		prp := NewParticle(0, 0, p[i].Vx, p[i].Vy, p[i].M)
		prp.X = p[i].X + p[i].Vx * dura
		prp.Y = p[i].Y + p[i].Vy * dura

		pdy := math.Pow(p[i].Y - ypth, 2.0)
		ndy := math.Pow(p[i].Y - ynth, 2.0)
		fpdy := kabeksu * dura / (pdy * p[i].M)
		fndy := kabeksu * dura / (ndy * p[i].M)
		prp.Vy += (p[i].Y - ypth) * fpdy + (p[i].Y - ynth) * fndy

		for pi := 0; pi < len(p); pi++ {
			if i == pi {
				continue
			}

			dsq := math.Pow(p[i].X - p[pi].X, 2.0) + math.Pow(p[i].Y - p[pi].Y, 2.0)
			f := dura / (dsq * p[i].M * p[pi].M)

			vmpx := p[i].X - p[pi].X
			vmpy := p[i].Y - p[pi].Y

			prp.Vx += pksu * vmpx * f
			prp.Vy += pksu * vmpy * f
		}

		for mi := 0; mi < len(m); mi++ {
			dsq := math.Pow(p[i].X - m[mi].X, 2.0) + math.Pow(p[i].Y - m[mi].Y, 2.0)
			f := dura / (dsq * p[i].M)

			vmpx := p[i].X - m[mi].X
			vmpy := p[i].Y - m[mi].Y

			prp.Vx += vmpx * f
			prp.Vy += vmpy * f
		}

		pres[i] = prp
	}

	return pres
}

func NextPsMs (p []Particle, m []Material, dura float64) []Particle {
	pres := make([]Particle, len(p))

	for i := 0; i < len(p); i++ {
		prp := NewParticle(0, 0, p[i].Vx, p[i].Vy, p[i].M)
		prp.X = p[i].X + p[i].Vx * dura
		prp.Y = p[i].Y + p[i].Vy * dura

		for pi := 0; pi < len(p); pi++ {
			if i == pi {
				continue
			}

			dsq := math.Pow(p[i].X - p[pi].X, 2.0) + math.Pow(p[i].Y - p[pi].Y, 2.0)
			f := dura / (dsq * p[i].M * p[pi].M)

			vmpx := p[i].X - p[pi].X
			vmpy := p[i].Y - p[pi].Y

			prp.Vx += vmpx * f
			prp.Vy += vmpy * f
		}

		for mi := 0; mi < len(m); mi++ {
			dsq := math.Pow(p[i].X - m[mi].X, 2.0) + math.Pow(p[i].Y - m[mi].Y, 2.0)
			f := dura / (dsq * p[i].M)

			vmpx := p[i].X - m[mi].X
			vmpy := p[i].Y - m[mi].Y

			prp.Vx += vmpx * f
			prp.Vy += vmpy * f
		}

		pres[i] = prp
	}

	return pres
}

func NewPos (inp Particle, dura float64) Particle {
	inp.X += inp.Vx * dura
	inp.Y += inp.Vy * dura

	return inp
}
