package main

import (
	"fmt"
	"image/color"
	"os"
	"time"

	"github.com/montanaflynn/stats"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

var anscombe = map[string][]float64{
	"x1": {10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
	"x2": {10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
	"x3": {10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
	"x4": {8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8},
	"y1": {8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68},
	"y2": {9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74},
	"y3": {7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73},
	"y4": {6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89},
}

func getDataset(i int, data map[string][]float64) ([]float64, []float64) {
	x := data[fmt.Sprintf("x%d", i)]
	y := data[fmt.Sprintf("y%d", i)]

	return x, y
}

func makePoints(x, y []float64) plotter.XYs {
	points := make(plotter.XYs, len(x))
	for i := range x {
		points[i].X = x[i]
		points[i].Y = y[i]
	}
	return points
}

func makePlot(i int, points plotter.XYs, title string) (*plot.Plot, error) {
	p := plot.New()
	p.Title.Text = title
	p.X.Label.Text = fmt.Sprintf("x%d", i)
	p.Y.Label.Text = fmt.Sprintf("y%d", i)
	p.X.Min = 0
	p.X.Max = 20
	p.Y.Min = 0
	p.Y.Max = 14

	scatter, err := plotter.NewScatter(points)
	if err != nil {
		return nil, err
	}

	scatter.GlyphStyle.Color = color.RGBA{B: 255, A: 255}
	scatter.GlyphStyle.Radius = vg.Points(3)
	scatter.GlyphStyle.Shape = draw.CircleGlyph{}
	p.Add(scatter)

	return p, err
}

func plotData(data map[string][]float64, filename string) error {
	plots := make([]*plot.Plot, 4)

	for i := 1; i <= 4; i++ {
		x, y := getDataset(i, data)

		points := makePoints(x, y)

		p, err := makePlot(i, points, fmt.Sprintf("Set %d", i))
		if err != nil {
			return err
		}
		plots[i-1] = p
	}
	img := vgimg.New(8*vg.Inch, 8*vg.Inch)
	dc := draw.New(img)

	tiles := draw.Tiles{
		Rows: 2,
		Cols: 2,
	}

	for i, p := range plots {
		if p == nil {
			fmt.Println(plots)
			return fmt.Errorf("plot %d is nil", i)
		}
		row := i % 2
		col := i / 2
		canvas := tiles.At(dc, row, col)
		p.Draw(canvas)
	}

	w, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer w.Close()

	png := vgimg.PngCanvas{Canvas: img}
	_, err = png.WriteTo(w)
	return err
}

func makeSeries(x, y []float64) ([]stats.Coordinate, error) {
	if len(x) != len(y) {
		return nil, fmt.Errorf("x and y are not the same length")
	}

	data := make([]stats.Coordinate, 0, len(x))
	for i := 0; i < len(x); i++ {
		data = append(data, stats.Coordinate{
			X: x[i],
			Y: y[i],
		})
	}

	return data, nil
}

func runRegression(x, y []float64) (slope, intercept float64, err error) {
	input, err := makeSeries(x, y)
	if err != nil {
		return 0, 0, err
	}

	result, err := stats.LinearRegression(input)
	if err != nil {
		return 0, 0, err
	}

	sl, in, err := getRegressionStats(result)
	if err != nil {
		return 0, 0, err
	}

	return sl, in, nil
}

func getRegressionStats(s stats.Series) (slope, intercept float64, err error) {
	if len(s) == 0 {
		return 0, 0, stats.EmptyInputErr
	}

	var sumX, sumY, sumXX, sumXY float64

	for i := range s {
		sumX += s[i].X
		sumY += s[i].Y
		sumXX += s[i].X * s[i].X
		sumXY += s[i].X * s[i].Y
	}

	n := float64(len(s))

	denominator := (n*sumXX - sumX*sumX)
	if denominator == 0 {
		return 0, 0, fmt.Errorf("division by zero")
	}

	slope = (n*sumXY - sumX*sumY) / denominator
	intercept = (sumY / n) - slope*(sumX/n)

	return slope, intercept, nil
}

func main() {
	plotData(anscombe, "fig_anscombe_Go.png")

	n := 1000

	for i := 1; i <= 4; i++ {
		x, y := getDataset(i, anscombe)

		var slopes float64
		var intercepts float64
		var runtimes time.Duration

		for j := 0; j < n; j++ {
			start := time.Now()

			slope, intercept, err := runRegression(x, y)
			if err != nil {
				fmt.Printf("dataset %d error: %v\n", i, err)
				continue
			}

			elapsed := time.Since(start)

			slopes += slope
			intercepts += intercept
			runtimes += elapsed

		}

		avgSlope := slopes / float64(n)
		avgIntercept := intercepts / float64(n)
		avgRuntime := runtimes / time.Duration(n)

		fmt.Printf(
			"Dataset %d -> avg slope: %.4f, avg intercept: %.4f, avg runtime: %v\n",
			i, avgSlope, avgIntercept, avgRuntime,
		)
	}

}
