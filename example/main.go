package main

import (
	"bufio"
	"fmt"
	"image/color"
	"log"
	"os"

	"bitbucket.org/differenttravel/interval"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/palette"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

const (
	MinX = 0
	MaxX = 40
)

func main() {
	filename := "data.txt"
	xys, err := readData(filename)
	if err != nil {
		log.Fatalf("could not read %s: %v", filename, err)
	}
	intervals := initIntervals(xys)
	err = plotData("out.png", intervals)
	if err != nil {
		log.Fatalf("could not plot data: %v", err)
	}
}

type xy struct{ x, y int }

func readData(path string) ([]xy, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var xys []xy

	// read line by line using a scanner (because we don't know if the file will be huge)
	s := bufio.NewScanner(f)
	for s.Scan() {
		var x, y int
		_, err := fmt.Sscanf(s.Text(), "%d,%d", &x, &y)
		if err != nil {
			log.Printf("discarding bad data point %v: %v", s.Text(), err)
		}
		xys = append(xys, xy{x, y})
	}
	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("could not scan: %v", err)
	}
	return xys, nil
}

func initIntervals(xys []xy) interval.Intervals {
	intervals := interval.NewIntervals(MinX, MaxX)
	for _, xy := range xys {
		intervals.Add(&interval.Interval{Low: xy.x, High: xy.y})
	}
	intervals.Sort()
	return intervals
}

func convertToPlotterXYs(intervals []*interval.Interval) plotter.XYs {
	pxys := plotter.XYs{}
	for _, intvl := range intervals {
		pxys = append(pxys, struct{ X, Y float64 }{X: float64(intvl.Low), Y: float64(intvl.High)})
	}
	return pxys
}

func alignPlots(plotItems []*plot.Plot) *vgimg.Canvas {
	rows, cols := len(plotItems), 1
	plots := make([][]*plot.Plot, rows)
	for j := 0; j < rows; j++ {
		plots[j] = make([]*plot.Plot, cols)
		for i := 0; i < cols; i++ {
			p := plotItems[j]

			// make sure the horizontal scales match
			p.X.Min = MinX
			p.X.Max = MaxX

			plots[j][i] = p
		}
	}

	img := vgimg.New(vg.Points(512), vg.Points(float64(128*rows)))
	dc := draw.New(img)

	t := draw.Tiles{
		Rows: rows,
		Cols: cols,
	}

	canvases := plot.Align(plots, t, dc)
	for j := 0; j < rows; j++ {
		for i := 0; i < cols; i++ {
			if plots[j][i] != nil {
				plots[j][i].Draw(canvases[j][i])
			}
		}
	}
	return img
}

func createFileFromCanvas(path string, img *vgimg.Canvas) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create %s: %v", path, err)
	}

	png := vgimg.PngCanvas{Canvas: img}
	if _, err := png.WriteTo(f); err != nil {
		return fmt.Errorf("could not write to %s: %v", path, err)
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("could not close %s: %v", path, err)
	}
	return nil
}

func createPlot(title string, xys plotter.XYs) (*plot.Plot, error) {
	p, err := plot.New()
	if err != nil {
		return nil, fmt.Errorf("could not create plot: %v", err)
	}

	// Draw a grid behind the data
	p.Add(plotter.NewGrid())
	p.Title.Text = title
	p.HideY()
	// p.X.Label.Text = "values"
	p.X.Padding = vg.Length(5)
	p.Y.Padding = vg.Length(20)
	plotIntervals(p, xys)
	return p, nil
}

func plotData(path string, intervals interval.Intervals) error {
	plots := []*plot.Plot{}

	// create Intervals plot
	xysIntervals := convertToPlotterXYs(intervals.Get())
	p1, err := createPlot("Intervals", xysIntervals)
	if err != nil {
		return fmt.Errorf("could not create plot: %v", err)
	}
	plots = append(plots, p1)

	// create Gaps plot
	xysGaps := convertToPlotterXYs(intervals.Gaps())
	p2, err := createPlot("Gaps", xysGaps)
	if err != nil {
		return fmt.Errorf("could not create plot: %v", err)
	}
	plots = append(plots, p2)

	// create Overlapped `plot
	xysOverlapped := convertToPlotterXYs(intervals.Overlapped())
	p3, err := createPlot("Overlapped", xysOverlapped)
	if err != nil {
		return fmt.Errorf("could not create plot: %v", err)
	}
	plots = append(plots, p3)

	// create Merged plot
	xysMerged := convertToPlotterXYs(intervals.Merge())
	p4, err := createPlot("Merged", xysMerged)
	if err != nil {
		return fmt.Errorf("could not create plot: %v", err)
	}
	plots = append(plots, p4)

	// join all plots, align them
	canvas := alignPlots(plots)
	err = createFileFromCanvas("out.png", canvas)
	if err != nil {
		return err
	}
	return nil
}

func plotIntervals(p *plot.Plot, xys plotter.XYs) error {
	var ps []plot.Plotter
	colors := getColors()
	numColors := len(colors)
	for i, xy := range xys {
		label := fmt.Sprintf("(%v,%v)", xy.X, xy.Y)
		pXYs := plotter.XYs{{xy.X, float64(i)}, {xy.Y, float64(i)}}
		color := colors[i%numColors]

		s, err := plotter.NewScatter(pXYs)

		if xy.X != xy.Y {
			l, err := plotter.NewLine(pXYs)
			l.Color = color
			l.Width = vg.Points(10)
			if err != nil {
				return fmt.Errorf("could not create a new line: %v", err)
			}
			ps = append(ps, l)
			p.Legend.Add(label, l)
		} else {
			s.Color = color
			p.Legend.Add(label, s)
		}

		if err != nil {
			return fmt.Errorf("could not create a new scatter: %v", err)
		}
		ps = append(ps, s)

	}
	p.Legend.Left = false
	p.Add(ps...)
	return nil
}

func getColors() []color.Color {
	palette := palette.Rainbow(10, 0, 1, 1, 1, 1)
	return palette.Colors()
}
