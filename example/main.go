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

type PlotType int

const (
	PlotTypeIntervals PlotType = iota
	PlotTypeGaps
	PlotTypeOverlapped
	PlotTypeMerged
)

type Superplot struct {
	Plot        *plot.Plot
	NumElements int
}

var (
	openBracket      = "["
	closeBracket     = "]"
	openParenthesis  = "("
	closeParenthesis = ")"
	intervalOpening  = openBracket
	intervalClosing  = closeBracket
	lowInclusive     bool
	highInclusive    bool
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
		var low, high int
		_, err := fmt.Sscanf(s.Text(), "%d,%d", &low, &high)
		if err != nil {
			log.Printf("discarding bad data point %v: %v", s.Text(), err)
			continue
		}
		if low > high {
			log.Printf("discarding bad data point (low, high)=(%v): low can not be greater than high", s.Text())
			continue
		}
		xys = append(xys, xy{low, high})
	}
	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("could not scan: %v", err)
	}
	return xys, nil
}

func initIntervals(xys []xy) interval.Intervals {
	// initialize Intervals
	minLow := MinX
	maxHigh := MaxX
	lowInclusive = true
	highInclusive = true
	selfAdjustMinLow := false
	selfAdjustMaxHigh := true
	intervals := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	if intervals.IsLowInclusive() {
		intervalOpening = "["
		lowInclusive = true
	} else {
		intervalOpening = "("
		lowInclusive = false
	}

	if intervals.IsHighInclusive() {
		intervalClosing = "]"
		highInclusive = true
	} else {
		intervalClosing = ")"
		highInclusive = false
	}

	for _, xy := range xys {
		err := intervals.AddInterval(&interval.Interval{Low: xy.x, High: xy.y})
		if err != nil {
			fmt.Printf("invalid interval discarded: %v\n", err)
		}
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

func alignPlots(plotItems []*Superplot, minLow, maxHigh int) *vgimg.Canvas {
	rows, cols := len(plotItems), 1
	plots := make([][]*plot.Plot, rows)
	for j := 0; j < rows; j++ {
		plots[j] = make([]*plot.Plot, cols)
		for i := 0; i < cols; i++ {
			p := plotItems[j]

			// make sure the horizontal scales match
			p.Plot.X.Min = float64(minLow)  // MinX
			p.Plot.X.Max = float64(maxHigh) //MaxX

			plots[j][i] = p.Plot
		}
	}

	img := vgimg.New(vg.Points(512), vg.Points(float64(200*rows)))
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

func createPlot(title string, xys plotter.XYs, plotType PlotType) (*Superplot, error) {
	p, err := plot.New()
	if err != nil {
		return nil, fmt.Errorf("could not create plot: %v", err)
	}

	// Draw a grid behind the data
	p.Add(plotter.NewGrid())
	p.Title.Text = title
	p.HideY()
	p.X.Padding = vg.Length(5)
	p.Y.Padding = vg.Length(20)

	plotIntervals(p, plotType, xys)
	return &Superplot{p, xys.Len()}, nil
}

func plotData(path string, intervals interval.Intervals) error {
	plots := []*Superplot{}

	// create Intervals plot
	xysIntervals := convertToPlotterXYs(intervals.GetIntervals())
	p1, err := createPlot("Intervals", xysIntervals, PlotTypeIntervals)
	if err != nil {
		return fmt.Errorf("could not create plot: %v", err)
	}
	plots = append(plots, p1)

	// create Gaps plot
	xysGaps := convertToPlotterXYs(intervals.Gaps())
	p2, err := createPlot("Gaps", xysGaps, PlotTypeGaps)
	if err != nil {
		return fmt.Errorf("could not create plot: %v", err)
	}
	plots = append(plots, p2)

	// create Overlapped `plot
	xysOverlapped := convertToPlotterXYs(intervals.Overlapped())
	p3, err := createPlot("Overlapped", xysOverlapped, PlotTypeOverlapped)
	if err != nil {
		return fmt.Errorf("could not create plot: %v", err)
	}
	plots = append(plots, p3)

	// create Merged plot
	xysMerged := convertToPlotterXYs(intervals.Merge())
	p4, err := createPlot("Merged", xysMerged, PlotTypeMerged)
	if err != nil {
		return fmt.Errorf("could not create plot: %v", err)
	}
	plots = append(plots, p4)

	// join all plots, align them
	canvas := alignPlots(plots, intervals.GetMinLow(), intervals.GetMaxHigh())
	err = createFileFromCanvas("out.png", canvas)
	if err != nil {
		return err
	}
	return nil
}

func plotIntervals(p *plot.Plot, plotType PlotType, xys plotter.XYs) error {
	var ps []plot.Plotter
	colors := getColors()
	numColors := len(colors)

	crossShape := draw.CrossGlyph{}
	ringShape := draw.RingGlyph{}
	legendWithCrossShape := false
	legendWithRingShape := false
	for i, xy := range xys {
		pXYLow := struct{ X, Y float64 }{xy.X, float64(i)}
		pXYHigh := struct{ X, Y float64 }{xy.Y, float64(i)}
		pXYs := plotter.XYs{pXYLow, pXYHigh}
		color := colors[i%numColors]

		var s1, s2 *plotter.Scatter
		var err error
		label := ""
		if plotType == PlotTypeIntervals {
			label = fmt.Sprintf("%s%v,%v%s", intervalOpening, xy.X, xy.Y, intervalClosing)
			s1, err = plotter.NewScatter(plotter.XYs{pXYLow})
			if err != nil {
				return fmt.Errorf("unable to create new scatter: %v", err)
			}
			if !lowInclusive {
				s1.GlyphStyle.Shape = crossShape
				if !legendWithCrossShape {
					p.Legend.Add("Exclusive", s1)
					legendWithCrossShape = true
				}
			} else {
				s1.GlyphStyle.Shape = ringShape
				if !legendWithRingShape {
					p.Legend.Add("Inclusive", s1)
					legendWithRingShape = true
				}
			}

			s2, err = plotter.NewScatter(plotter.XYs{pXYHigh})
			if err != nil {
				return fmt.Errorf("unable to create new scatter: %v", err)
			}
			if !highInclusive {
				s2.GlyphStyle.Shape = crossShape
				if !legendWithCrossShape {
					p.Legend.Add("Exclusive", s2)
					legendWithCrossShape = true
				}
			} else {
				s2.GlyphStyle.Shape = ringShape
				if !legendWithRingShape {
					p.Legend.Add("Inclusive", s2)
					legendWithRingShape = true
				}
			}
		} else {
			label = fmt.Sprintf("%s%v,%v%s", openBracket, xy.X, xy.Y, closeBracket)
			s1, err = plotter.NewScatter(pXYs)
			if err != nil {
				return fmt.Errorf("unable to create new scatter: %v", err)
			}
			if !lowInclusive && !highInclusive {
				s1.GlyphStyle.Shape = crossShape
			}
		}

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
			s1.Color = color
			if s2 != nil {
				s2.Color = color
			}
			p.Legend.Add(label, s1)
		}

		if err != nil {
			return fmt.Errorf("could not create a new scatter: %v", err)
		}
		ps = append(ps, s1)
		if s2 != nil {
			ps = append(ps, s2)
		}
	}
	p.Legend.Left = false
	p.Add(ps...)
	return nil
}

func getColors() []color.Color {
	palette := palette.Rainbow(10, 0, 1, 1, 1, 1)
	return palette.Colors()
}
