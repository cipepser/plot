package myutil

import (
	"log"
	"os/exec"
	"strconv"

	"github.com/cipepser/plot/plotter"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
)

// MySinglePlot is a wrapper of Line of package plotter with slice of float64 x.
func MySinglePlot(x []float64) {
	data := make(plotter.XYs, len(x))
	for i := 0; i < len(x); i++ {
		data[i].X = float64(i)
		data[i].Y = x[i]
	}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	l, err := plotter.NewLine(data)
	if err != nil {
		panic(err)
	}

	p.Add(l)

	file := "img.png"
	if err = p.Save(10*vg.Inch, 6*vg.Inch, file); err != nil {
		panic(err)
	}

	if err = exec.Command("open", file).Run(); err != nil {
		panic(err)
	}
}

// MyPlot is a wrapper of Line of package plotter with slice of float64 x and y.
func MyPlot(x, y []float64) {
	if len(x) != len(y) {
		log.Fatal("length of x and y have to same.")
	}

	data := make(plotter.XYs, len(x))
	for i := 0; i < len(x); i++ {
		data[i].X = x[i]
		data[i].Y = y[i]
	}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	l, err := plotter.NewLine(data)
	if err != nil {
		panic(err)
	}

	p.Add(l)

	file := "img.png"
	if err = p.Save(10*vg.Inch, 6*vg.Inch, file); err != nil {
		panic(err)
	}

	if err = exec.Command("open", file).Run(); err != nil {
		panic(err)
	}
}

// MySingleScatter is a wrapper of Scatter of package plotter with slice of float64 x.
func MySingleScatter(x []float64) {
	data := make(plotter.XYs, len(x))
	for i := 0; i < len(x); i++ {
		data[i].X = float64(i)
		data[i].Y = x[i]
	}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	s, err := plotter.NewScatter(data)
	if err != nil {
		panic(err)
	}

	s.Radius = vg.Length(1)

	p.Add(s)

	file := "img.png"
	if err = p.Save(10*vg.Inch, 6*vg.Inch, file); err != nil {
		panic(err)
	}

	if err = exec.Command("open", file).Run(); err != nil {
		panic(err)
	}

}

// MyScatter is a wrapper of Scatter of package plotter with slice of float64 x and y.
func MyScatter(x, y []float64) {
	if len(x) != len(y) {
		log.Fatal("length of x and y have to same.")
	}

	data := make(plotter.XYs, len(x))
	for i := 0; i < len(x); i++ {
		data[i].X = x[i]
		data[i].Y = y[i]
	}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	s, err := plotter.NewScatter(data)
	if err != nil {
		panic(err)
	}

	s.Radius = vg.Length(2)

	p.Add(s)

	file := "img.png"
	if err = p.Save(10*vg.Inch, 6*vg.Inch, file); err != nil {
		panic(err)
	}

	if err = exec.Command("open", file).Run(); err != nil {
		panic(err)
	}
}

// MyPlotWithScatter draw plot and scatter at once.
func MyPlotWithScatter(x, y []float64) {
	if len(x) != len(y) {
		log.Fatal("length of x and y have to same.")
	}

	data := make(plotter.XYs, len(x))
	for i := 0; i < len(x); i++ {
		data[i].X = x[i]
		data[i].Y = y[i]
	}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	s, err := plotter.NewScatter(data)
	if err != nil {
		panic(err)
	}

	s.Radius = vg.Length(2)
	p.Add(s)

	l, err := plotter.NewLine(data)
	if err != nil {
		panic(err)
	}

	p.Add(l)

	file := "img.png"
	if err = p.Save(10*vg.Inch, 6*vg.Inch, file); err != nil {
		panic(err)
	}

	if err = exec.Command("open", file).Run(); err != nil {
		panic(err)
	}
}

// MyCandleChart draw the candle chart with data.
// ts represents the time which used as label.
func MyCandleChart(ts []string, data [][]float64, bu plotter.BarUnit) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	cc, err := plotter.NewCandleChart(data)
	if err != nil {
		panic(err)
	}

	p.Add(cc)

	// tunit := "min"
	cunit := "yen"
	p.Title.Text = "Candle Chart"
	p.X.Label.Text = "Time"
	p.X.Label.Text = "Time [" + strconv.Itoa(bu.T) + " " + plotter.TransFormat2Unit(bu.Unit) + "]"
	p.Y.Label.Text = "Price [" + cunit + "]"

	// fmt.Println("")
	// p.X.Tick.Marker.Ticks(0, 0)
	// fmt.Println(p.X.Tick.Marker.Ticks(0, 0)[0])
	// fmt.Println(p.Y.Tick.Marker.Ticks(0, 0)[0])
	p.Y.Tick.Marker = plotter.RawTicks{}
	// p.Y.Tick.Marker = commaTicks{}

	p.NominalX(ts...)

	p.X.Min = -0.5
	p.X.Max = float64(len(data)) * 1.1

	file := "img.png"
	if err = p.Save(10*vg.Inch, 6*vg.Inch, file); err != nil {
		panic(err)
	}

	if err = exec.Command("open", file).Run(); err != nil {
		panic(err)
	}

}
