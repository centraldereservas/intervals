package intervals

import "fmt"

var blockSize = 10

type symbols struct {
	emptySymbol    string
	fullSymbol     string
	overlapSymbol  string
	leadingSymbol  string
	trailingSymbol string
	separator      string
}

func newSymbols() symbols {
	// Available Symbols:  ( ◯ ◌ ◍ ◎ ● ◉ ) , ( □ ■ ), ( ░ ▒ ▓ █ )
	return symbols{
		emptySymbol:    "◌",
		fullSymbol:     "◎",
		overlapSymbol:  "●",
		leadingSymbol:  "├",
		trailingSymbol: "┤",
		separator:      "|",
	}
}

func (intvls *intervals) report() string {
	intvls.Sort()
	symbols := newSymbols()
	intro := intvls.buildHeading()
	legend := intvls.buildLegend(symbols)
	overlapText := intvls.reportOverlapped()
	gapsText := intvls.reportGaps()
	mergeText := intvls.reportMerged()
	intervalText, graph := intvls.buildGraph(symbols)
	axisLegend := intvls.buildAxisLegend()
	graphWithLegend := fmt.Sprintf("\n\n%s\n%s%s%s", axisLegend, symbols.leadingSymbol, graph, symbols.trailingSymbol)
	return "\n" + intro + intervalText + gapsText + overlapText + mergeText + legend + graphWithLegend + "\n"
}

func (intvls *intervals) buildGraph(symbols symbols) (string, string) {
	index := intvls.MinLow
	numSeparators := 0
	graph := ""
	intervalText := ""
	intervalText, index, graph, numSeparators = intvls.reportIntervals(index, symbols, numSeparators)
	graph, numSeparators, index = intvls.fillUntilTheEnd(index, symbols, graph, numSeparators)
	return intervalText, graph
}

func (intvls *intervals) buildHeading() string {
	introText := fmt.Sprintf("\n==================================\n REPORT (minLow=%d, maxHigh=%d)\n==================================", intvls.MinLow, intvls.MaxHigh)
	return introText
}

func (intvls *intervals) buildLegend(symbols symbols) string {
	legend := fmt.Sprintf("\n • Legend: %v (empty), %v (full), %v (overlap)", symbols.emptySymbol, symbols.fullSymbol, symbols.overlapSymbol)
	return legend
}

func (intvls *intervals) reportIntervals(index int, symbols symbols, numSeparators int) (string, int, string, int) {
	graph := ""
	intervalText := "\n • Intervals: "
	for j, intvl := range intvls.Intervals {
		if j != 0 {
			intervalText += ", "
		}
		intervalText += fmt.Sprintf("[%d,%d]", intvl.Low, intvl.High)
		for i := index; i < intvl.Low; i++ {
			index++
			graph += symbols.emptySymbol
			if index%blockSize == 0 {
				graph += symbols.separator
				numSeparators++
			}
		}

		for i := index; i <= intvl.High; i++ {
			if intvls.valueIsOverlapping(index, intvls.OverlappedList) {
				graph += symbols.overlapSymbol
			} else {
				graph += symbols.fullSymbol
			}
			index++
			if index%blockSize == 0 {
				graph += symbols.separator
				numSeparators++
			}
		}
	}
	return intervalText, index, graph, numSeparators
}

func (intvls *intervals) fillUntilTheEnd(index int, symbols symbols, graph string, numSeparators int) (string, int, int) {
	for i := index + 1; i <= intvls.MaxHigh+1; i++ {
		graph += symbols.emptySymbol
		if i%blockSize == 0 {
			graph += symbols.separator
			numSeparators++
		}
	}
	return graph, numSeparators, index
}

func (intvls *intervals) reportOverlapped() string {
	overlapText := "\n • Overlapped: "
	overlapped := intvls.Overlapped()
	for i, ovrlp := range overlapped {
		if i != 0 {
			overlapText += ", "
		}
		overlapText += fmt.Sprintf("[%d,%d]", ovrlp.Low, ovrlp.High)
	}
	return overlapText
}

func (intvls *intervals) reportGaps() string {
	gapsText := "\n • Gaps: "
	gaps := intvls.Gaps()
	for i, gap := range gaps {
		if i != 0 {
			gapsText += ", "
		}
		gapsText += fmt.Sprintf("[%d,%d]", gap.Low, gap.High)
	}
	return gapsText
}

func (intvls *intervals) reportMerged() string {
	mergedText := "\n • Merged: "
	merged := intvls.Merge()
	for i, merge := range merged {
		if i != 0 {
			mergedText += ", "
		}
		mergedText += fmt.Sprintf("[%d,%d]", merge.Low, merge.High)
	}
	return mergedText
}

func (intvls *intervals) buildAxisLegend() string {
	axisLegend := " "
	for i := intvls.MinLow; i < intvls.MaxHigh/blockSize; i++ {
		mark := fmt.Sprintf("%d", i*blockSize)
		axisLegend += mark
		limit := (blockSize - len(mark)) + 1
		for j := 0; j < limit; j++ {
			axisLegend += " "
		}
	}
	axisLegend += fmt.Sprintf("%v", intvls.MaxHigh)
	return axisLegend
}
