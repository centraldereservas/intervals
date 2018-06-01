package interval

import (
	"fmt"
	"image/color"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/palette"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

type PlotType int

const (
	PlotTypeIntervals PlotType = iota
	PlotTypeGaps
	PlotTypeOverlapped
	PlotTypeMerged
)

const (
	openBracket      = "["
	closeBracket     = "]"
	openParenthesis  = "("
	closeParenthesis = ")"
)

type IntervalPlot interface {
	PlotData(path string, intervals Intervals, addIntervals, addGaps, addOverlapped, addMerges bool) error
	CreatePlot(title string, xys plotter.XYs, plotType PlotType) (*Superplot, error)
	AlignPlots(plotItems []*Superplot, minLow, maxHigh int) *vgimg.Canvas
	CreateFileFromCanvas(path string, img *vgimg.Canvas) error
}

type intervalPlot struct {
	lowInclusive    bool
	highInclusive   bool
	intervalOpening string
	intervalClosing string
}

func NewPlot(lowInclusive bool, highInclusive bool) IntervalPlot {
	opening := ""
	if lowInclusive {
		opening = openBracket
	} else {
		opening = openParenthesis
	}

	closing := ""
	if highInclusive {
		closing = closeBracket
	} else {
		closing = closeParenthesis
	}

	ip := &intervalPlot{
		lowInclusive:    lowInclusive,
		highInclusive:   highInclusive,
		intervalOpening: opening,
		intervalClosing: closing,
	}
	return ip
}

type Superplot struct {
	Plot        *plot.Plot
	NumElements int
}

func (ip *intervalPlot) PlotData(path string, intervals Intervals, addIntervals, addGaps, addOverlapped, addMerges bool) error {
	plots := []*Superplot{}
	var err error

	if addIntervals {
		// create Intervals plot
		xysIntervals := ip.convertToPlotterXYs(intervals.GetIntervals())
		p1, err := ip.CreatePlot("Intervals", xysIntervals, PlotTypeIntervals)
		if err != nil {
			return fmt.Errorf("could not create plot: %v", err)
		}
		plots = append(plots, p1)
	}

	if addGaps {
		// create Gaps plot
		xysGaps := ip.convertToPlotterXYs(intervals.Gaps())
		p2, err := ip.CreatePlot("Gaps", xysGaps, PlotTypeGaps)
		if err != nil {
			return fmt.Errorf("could not create plot: %v", err)
		}
		plots = append(plots, p2)
	}

	if addOverlapped {
		// create Overlapped `plot
		xysOverlapped := ip.convertToPlotterXYs(intervals.Overlapped())
		p3, err := ip.CreatePlot("Overlapped", xysOverlapped, PlotTypeOverlapped)
		if err != nil {
			return fmt.Errorf("could not create plot: %v", err)
		}
		plots = append(plots, p3)
	}

	if addMerges {
		// create Merged plot
		xysMerged := ip.convertToPlotterXYs(intervals.Merge())
		p4, err := ip.CreatePlot("Merged", xysMerged, PlotTypeMerged)
		if err != nil {
			return fmt.Errorf("could not create plot: %v", err)
		}
		plots = append(plots, p4)
	}

	// join all plots, align them
	canvas := ip.AlignPlots(plots, intervals.GetMinLow(), intervals.GetMaxHigh())
	err = ip.CreateFileFromCanvas("out.png", canvas)
	if err != nil {
		return err
	}
	return nil
}

func (ip *intervalPlot) CreateFileFromCanvas(path string, img *vgimg.Canvas) error {
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

func (ip *intervalPlot) AlignPlots(plotItems []*Superplot, minLow, maxHigh int) *vgimg.Canvas {
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

func (ip *intervalPlot) CreatePlot(title string, xys plotter.XYs, plotType PlotType) (*Superplot, error) {
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

	ip.plotIntervals(p, plotType, xys)
	return &Superplot{p, xys.Len()}, nil
}

func (ip *intervalPlot) plotIntervals(p *plot.Plot, plotType PlotType, xys plotter.XYs) error {
	var ps []plot.Plotter
	colors := ip.getColors()
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
			label = fmt.Sprintf("%s%v,%v%s", ip.intervalOpening, xy.X, xy.Y, ip.intervalClosing)
			s1, err = plotter.NewScatter(plotter.XYs{pXYLow})
			if err != nil {
				return fmt.Errorf("unable to create new scatter: %v", err)
			}
			if !ip.lowInclusive {
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
			if !ip.highInclusive {
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
			if !ip.lowInclusive && !ip.highInclusive {
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

func (ip *intervalPlot) getColors() []color.Color {
	palette := palette.Rainbow(10, 0, 1, 1, 1, 1)
	return palette.Colors()
}

func (ip *intervalPlot) convertToPlotterXYs(intervals []*Interval) plotter.XYs {
	pxys := plotter.XYs{}
	for _, intvl := range intervals {
		pxys = append(pxys, struct{ X, Y float64 }{X: float64(intvl.Low), Y: float64(intvl.High)})
	}
	return pxys
}
