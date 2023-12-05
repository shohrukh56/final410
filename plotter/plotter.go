// plotter.go
package plotter

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"log"
	"runtime"
	"testing"
)

// CreatePlot generates and saves a line plot for benchmark results
func CreatePlot(title string, results testing.BenchmarkResult, filename string) {
	// Extract data points from benchmark results
	var dataPoints plotter.XYs

	for i := 0; i < results.N; i++ {
		dataPoints = append(dataPoints, plotter.XY{X: float64(i), Y: float64(results.T.Nanoseconds()) / 1e6})
	}

	// Create plot
	p := plot.New()

	p.Title.Text = title
	p.X.Label.Text = "Number of Iterations"
	p.Y.Label.Text = "Time Taken (ms)"

	// Create a line plot
	line, err := plotter.NewLine(dataPoints)
	if err != nil {
		log.Fatal(err)
	}

	// Set line style
	line.LineStyle.Width = vg.Points(1)
	line.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}

	p.Add(line)

	// Save the plot to a file
	if err := p.Save(8*vg.Inch, 4*vg.Inch, filename); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Line plot saved to %s\n", filename)
}

func GeneratePlots() {
	// Set GOMAXPROCS to utilize multiple cores
	runtime.GOMAXPROCS(runtime.NumCPU())
}
