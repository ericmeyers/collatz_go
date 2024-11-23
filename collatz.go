package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

)

func collatz_length_brian(n uint) uint {
	if n == 1 {
		return 1
	} else if n%2 == 0 {
		return 1 + collatz_length_brian(n/2)

	} else {
		return 1 + collatz_length_brian(3*n+1)
	}
}

func main() {
	// ugh. really don't like this language
	//

	pts := make(plotter.XYs, 5000000)
	// and iterators are shit?
	for e := 0; e < len(pts); e++ {
		x := e + 1
		y := collatz_length_brian(uint(x))
		pts[e].X = float64(x)
		pts[e].Y = float64(y)
	}
	//	pts := plotter.XYs{{X: 0.5, Y: 0.5}, {X: 1, Y: 0.5}}
	p := plot.New()
	p.Add(plotter.NewGrid())
	p.Title.Text = "Collatz Sequence Length"
	p.X.Label.Text = "Start Value"
	p.Y.Label.Text = "Sequence Length"
	s, err := plotter.NewScatter(pts)
	if err != nil {
		panic(err)
	}
	p.Add(s)
	if err := p.Save(8*vg.Inch, 6*vg.Inch, "collatz.png"); err != nil {
		panic(err)
	}
}
