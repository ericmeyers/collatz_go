package main

import (
	//	"gonum.org/v1/plot"
	//	"gonum.org/v1/plot/plotter"
	//	"gonum.org/v1/plot/vg"
	grob "github.com/MetalBlueberry/go-plotly/generated/v2.31.1/graph_objects"
	"github.com/MetalBlueberry/go-plotly/pkg/offline"
	"github.com/MetalBlueberry/go-plotly/pkg/types"


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

	//	pts := make(plotter.XYs, 5000000)
	// and iterators are shit?
	//
	xs := make([]uint,5000000)
	ys := make([]uint,5000000)
	for e := 0; e < len(xs); e++ {
		x := e + 1
		y := collatz_length_brian(uint(x))
		//pts[e].X = float64(x)
		//pts[e].Y = float64(y)
		xs[e] = uint(x)
		ys[e] = y
	}
	//
/* gonum plot
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
*/
	// plotly. oh my WebGL!!
	fig := &grob.Fig{
		Data: []types.Trace{
			&grob.Scattergl{
				X: types.DataArray(xs),
				Y: types.DataArray(ys),
				Mode: grob.ScatterglModeMarkers,
			},
		},
		Layout: &grob.Layout{
			Title: &grob.LayoutTitle{
				Text: "Collatz Sequence Length",
			},
		},
	}

	offline.Show(fig)
	// renders @ localhost:8080
	//offline.Serve(fig)

}
