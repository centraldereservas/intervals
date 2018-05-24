package interval

import (
	"fmt"
	"math"
	"sort"
)

const (
	defaultMinLow  = 0
	defaultMaxHigh = math.MaxInt64
)

// Intervals is an interface to handle Interval structures discovering the existence of gaps or overlays
type Intervals interface {
	Add(itvl *Interval)

	// Sort sorts the intervals list by the Low property (ascending)
	Sort()

	// Gaps first sorts (if necessary) and then returns the interval gaps
	Gaps() []*Interval

	// Overlapped first sorts (if necessary) and then returns the overlapped intervals
	Overlapped() []*Interval

	// FindIntervalsForValue returns all the intervals which contains the passed value
	FindIntervalsForValue(value int) []*Interval

	// Print first sorts (if necessary) and then displays graphically the interval sequence
	Print() string
}

// intervals implements Intervals interface
type intervals struct {
	Intervals []*Interval
	MinLow    int
	MaxHigh   int
	Sorted    bool
}

// NewIntervalsDefault is a constructor that returns an instance of the Intervals interface with default values
func NewIntervalsDefault() Intervals {
	return NewIntervals(defaultMinLow, defaultMaxHigh)
}

// NewIntervals is a constructor that returns an instance of the Intervals interface
func NewIntervals(minLow int, maxHigh int) Intervals {
	return &intervals{
		MinLow:    minLow,
		MaxHigh:   maxHigh,
		Intervals: []*Interval{},
		Sorted:    false,
	}
}

func (intvls *intervals) Add(itvl *Interval) {
	intvls.Intervals = append(intvls.Intervals, itvl)
	intvls.Sorted = false
}

func (intvls *intervals) FindIntervalsForValue(value int) []*Interval {
	intvls.Sort()
	var matches []*Interval
	for _, intvl := range intvls.Intervals {
		if intvl.Low > value {
			// due to the intervals are sorted, we can confirm that we will not find more matches
			break
		}
		if inBetweenInclusive(value, intvl.Low, intvl.High) {
			matches = append(matches, intvl)
		}
	}
	return matches
}

func (intvls *intervals) Sort() {
	if !intvls.Sorted {
		sort.Sort(ByLow(intvls.Intervals))
	}
	intvls.Sorted = true
}

func (intvls *intervals) Gaps() []*Interval {
	intvls.Sort()
	gaps := []*Interval{}
	lastMaxHigh := intvls.MinLow
	for _, intvl := range intvls.Intervals {
		if intvl.Low > lastMaxHigh {
			gaps = append(gaps, &Interval{Low: lastMaxHigh, High: intvl.Low - 1})
		}
		// lastHigh = intvl.High + 1
		if intvl.High >= lastMaxHigh {
			lastMaxHigh = intvl.High + 1
		}
	}
	if lastMaxHigh < intvls.MaxHigh {
		gaps = append(gaps, &Interval{Low: lastMaxHigh, High: intvls.MaxHigh})
	}
	return gaps
}

func (intvls *intervals) Overlapped() []*Interval {
	intvls.Sort()
	list := []*Interval{}
	lastMinLow := math.MaxInt64
	lastMaxHigh := math.MinInt64
	for i, intvl := range intvls.Intervals {
		if i > 0 {
			lowInBetween := inBetweenInclusive(lastMinLow, intvl.Low, intvl.High) || inBetweenInclusive(intvl.Low, lastMinLow, lastMaxHigh)
			highInBetween := inBetweenInclusive(lastMaxHigh, intvl.Low, intvl.High) || inBetweenInclusive(intvl.High, lastMinLow, lastMaxHigh)
			if lowInBetween || highInBetween {
				greaterLow := max(intvl.Low, lastMinLow)
				lowerHigh := min(intvl.High, lastMaxHigh)
				list = append(list, &Interval{Low: greaterLow, High: lowerHigh})
			}
		}
		if intvl.Low < lastMinLow {
			lastMinLow = intvl.Low
		}
		if intvl.High > lastMaxHigh {
			lastMaxHigh = intvl.High
		}
	}
	return list
}

func (intvls *intervals) isOverlapping(value int, overlapped []*Interval) bool {
	for _, ovrlp := range overlapped {
		if inBetweenInclusive(value, ovrlp.Low, ovrlp.High) {
			return true
		}
	}
	return false
}

func (intvls *intervals) Print() string {
	intvls.Sort()

	// Available Symbols:  ( ◯ ◌ ◍ ◎ ● ◉ ) , ( □ ■ ), ( ░ ▒ ▓ █ )
	emptySymbol := "◌"
	fullSymbol := "◎"
	overlapSymbol := "●"
	leadingSymbol := "├"
	trailingSymbol := "┤"
	separator := "|"

	introText := fmt.Sprintf("\n==================================\n SUMMARY (minLow=%d, maxHigh=%d)\n==================================", intvls.MinLow, intvls.MaxHigh)
	legend := fmt.Sprintf("\n • Legend: %v (empty), %v (full), %v (overlap)", emptySymbol, fullSymbol, overlapSymbol)
	intervalText := "\n • Intervals: "
	gapsText := "\n • Gaps: "
	overlapText := "\n • Overlapped: "
	graph := ""
	index := intvls.MinLow
	blockSize := 10
	numSeparators := 0

	overlapped := intvls.Overlapped()
	for i, ovrlp := range overlapped {
		if i != 0 {
			overlapText += ", "
		}
		overlapText += fmt.Sprintf("[%d,%d]", ovrlp.Low, ovrlp.High)
	}

	for j, intvl := range intvls.Intervals {
		if j != 0 {
			intervalText += ", "
		}
		intervalText += fmt.Sprintf("[%d,%d]", intvl.Low, intvl.High)
		for i := index; i < intvl.Low; i++ {
			index++
			graph += emptySymbol
			if index%blockSize == 0 {
				graph += separator
				numSeparators++
			}
		}

		for i := index; i <= intvl.High; i++ {
			if intvls.isOverlapping(index, overlapped) {
				graph += overlapSymbol
			} else {
				graph += fullSymbol
			}
			index++
			if index%blockSize == 0 {
				graph += separator
				numSeparators++
			}
		}
	}
	gaps := intvls.Gaps()
	for i, gap := range gaps {
		if i != 0 {
			gapsText += ", "
		}
		gapsText += fmt.Sprintf("[%d,%d]", gap.Low, gap.High)
	}

	for i := index + 1; i < intvls.MaxHigh; i++ {
		graph += emptySymbol
		if i%blockSize == 0 {
			graph += separator
			numSeparators++
		}
	}
	axisLegend := " "
	numSeparators = 0
	for i := intvls.MinLow; i < intvls.MaxHigh/blockSize; i++ {
		mark := fmt.Sprintf("%d", i*blockSize)
		axisLegend += mark
		limit := (blockSize - len(mark)) + 1
		for j := 0; j < limit; j++ {
			axisLegend += " "
		}
	}
	axisLegend += fmt.Sprintf("%v", intvls.MaxHigh)
	graphText := fmt.Sprintf("\n\n%s\n%s%s%s", axisLegend, leadingSymbol, graph, trailingSymbol)
	return "\n" + introText + legend + intervalText + gapsText + overlapText + graphText + "\n"
}
