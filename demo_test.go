package interval_test

import (
	"fmt"
	"math"

	"github.com/centraldereservas/intervals"
)

type demo struct {
	Intervals           interval.Intervals
	ExpectedGaps        []interval.Interval
	ExpectedOverlaps    []interval.Interval
	ExpectedMerges      []interval.Interval
	ExpectedFindMatches []interval.Interval
	ValueToFind         int
}

///////////////////////////////////////////////////////////
///   Tests from 000 to 099: low/high are inclusive     ///
///////////////////////////////////////////////////////////

//  no intervals (low/high inclusive) --> all is a gap  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo001() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := true
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}

	// calculate expected merges
	merges := []interval.Interval{}

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

// one interval at the beginning (low/high inclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo002() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := true
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: minLow, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: 5, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: minLow, High: 4})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

//  one interval at the end (low/high inclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo003() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := true
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: 8, High: maxHigh}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 7})

	// calculate expected overlaps
	overlaps := []interval.Interval{}

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 8, High: maxHigh})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

//  one interval in the middle (low/high inclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo004() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := true
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: 5, High: 8}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 4})
	gaps = append(gaps, interval.Interval{Low: 9, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 5, High: 8})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

//  two intervals in the middle, one inside the other (low/high inclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo005() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := true
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: 2, High: 8}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 4, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 1})
	gaps = append(gaps, interval.Interval{Low: 9, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	overlaps = append(overlaps, interval.Interval{Low: 4, High: 6})

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 2, High: 8})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

// two intervals in the middle, not overlapping (low/high inclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo006() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := true
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: 2, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 6, High: 8}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 1})
	gaps = append(gaps, interval.Interval{Low: 5, High: 5})
	gaps = append(gaps, interval.Interval{Low: 9, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 2, High: 4})
	merges = append(merges, interval.Interval{Low: 6, High: 8})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

//  two intervals in the middle, consecutives (not overlapping) (low/high inclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo007() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := true
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: 2, High: 3}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 4, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 1})
	gaps = append(gaps, interval.Interval{Low: 7, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 2, High: 6})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

// two intervals in the middle, overlapping by 1 position (low/high inclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo008() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := true
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: 2, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 6, High: 7}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 1})
	gaps = append(gaps, interval.Interval{Low: 8, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	overlaps = append(overlaps, interval.Interval{Low: 6, High: 6})

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 2, High: 7})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

//  two intervals in the middle, overlapping by 3 positions (low/high inclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo009() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := true
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: 2, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 4, High: 8}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 1})
	gaps = append(gaps, interval.Interval{Low: 9, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	overlaps = append(overlaps, interval.Interval{Low: 4, High: 6})

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 2, High: 8})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

// three intervals (leading, middle and trailing), not overlapping (low/high inclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo010() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := true
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: minLow, High: 2}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 4, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 8, High: maxHigh}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: 3, High: 3})
	gaps = append(gaps, interval.Interval{Low: 5, High: 7})

	// calculate expected overlaps
	overlaps := []interval.Interval{}

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: minLow, High: 2})
	merges = append(merges, interval.Interval{Low: 4, High: 4})
	merges = append(merges, interval.Interval{Low: 8, High: maxHigh})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

//  three intervals (in the middle), overlapping (low/high inclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo011() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := true
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: 2, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 3, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 5, High: 7}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 1})
	gaps = append(gaps, interval.Interval{Low: 8, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	overlaps = append(overlaps, interval.Interval{Low: 3, High: 4})
	overlaps = append(overlaps, interval.Interval{Low: 5, High: 6})

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 2, High: 7})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

// complex case with a lot of intervals and overlapping (low/high inclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo012() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 100
	lowInclusive := true
	highInclusive := true
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: 5, High: 7}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 2, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 35, High: 35}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 3, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 18, High: 20}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 20, High: 30}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 25, High: 28}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: minLow, High: 1}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 30, High: 32}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 10, High: 12}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: 8, High: 9})
	gaps = append(gaps, interval.Interval{Low: 13, High: 17})
	gaps = append(gaps, interval.Interval{Low: 33, High: 34})
	gaps = append(gaps, interval.Interval{Low: 36, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	overlaps = append(overlaps, interval.Interval{Low: 3, High: 4})
	overlaps = append(overlaps, interval.Interval{Low: 5, High: 6})
	overlaps = append(overlaps, interval.Interval{Low: 20, High: 20})
	overlaps = append(overlaps, interval.Interval{Low: 25, High: 28})
	overlaps = append(overlaps, interval.Interval{Low: 30, High: 30})

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: minLow, High: 7})
	merges = append(merges, interval.Interval{Low: 10, High: 12})
	merges = append(merges, interval.Interval{Low: 18, High: 32})
	merges = append(merges, interval.Interval{Low: 35, High: 35})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

// complex case with a lot of intervals and overlapping (low/high inclusive, selfAdjustMinLow=false and selfAdjustMaxHigh=true)
func buildIntervalsDemo013() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := math.MaxInt64
	lowInclusive := true
	highInclusive := true
	selfAdjustMinLow := false
	selfAdjustMaxHigh := true
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: 5, High: 7}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 2, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 35, High: 35}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 3, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 18, High: 20}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 20, High: 30}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 25, High: 28}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: minLow, High: 1}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 30, High: 32}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 10, High: 12}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: 8, High: 9})
	gaps = append(gaps, interval.Interval{Low: 13, High: 17})
	gaps = append(gaps, interval.Interval{Low: 33, High: 34})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	overlaps = append(overlaps, interval.Interval{Low: 3, High: 4})
	overlaps = append(overlaps, interval.Interval{Low: 5, High: 6})
	overlaps = append(overlaps, interval.Interval{Low: 20, High: 20})
	overlaps = append(overlaps, interval.Interval{Low: 25, High: 28})
	overlaps = append(overlaps, interval.Interval{Low: 30, High: 30})

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: minLow, High: 7})
	merges = append(merges, interval.Interval{Low: 10, High: 12})
	merges = append(merges, interval.Interval{Low: 18, High: 32})
	merges = append(merges, interval.Interval{Low: 35, High: 35})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

// complex case with a lot of intervals and overlapping (low/high inclusive, selfAdjustMinLow=true and selfAdjustMaxHigh=true)
func buildIntervalsDemo014() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := math.MaxInt64
	lowInclusive := true
	highInclusive := true
	selfAdjustMinLow := true
	selfAdjustMaxHigh := true
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: 5, High: 7}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 2, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 35, High: 35}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 3, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 18, High: 20}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 20, High: 30}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 25, High: 28}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 30, High: 32}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 10, High: 12}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: 8, High: 9})
	gaps = append(gaps, interval.Interval{Low: 13, High: 17})
	gaps = append(gaps, interval.Interval{Low: 33, High: 34})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	overlaps = append(overlaps, interval.Interval{Low: 3, High: 4})
	overlaps = append(overlaps, interval.Interval{Low: 5, High: 6})
	overlaps = append(overlaps, interval.Interval{Low: 20, High: 20})
	overlaps = append(overlaps, interval.Interval{Low: 25, High: 28})
	overlaps = append(overlaps, interval.Interval{Low: 30, High: 30})

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 2, High: 7})
	merges = append(merges, interval.Interval{Low: 10, High: 12})
	merges = append(merges, interval.Interval{Low: 18, High: 32})
	merges = append(merges, interval.Interval{Low: 35, High: 35})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

// complex case with a lot of intervals and overlapping (low/high inclusive, selfAdjustMinLow=true and selfAdjustMaxHigh=false)
func buildIntervalsDemo015() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 40
	lowInclusive := true
	highInclusive := true
	selfAdjustMinLow := true
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: 5, High: 7}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 2, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 35, High: 35}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 3, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 18, High: 20}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 20, High: 30}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 25, High: 28}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 30, High: 32}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 10, High: 12}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: 8, High: 9})
	gaps = append(gaps, interval.Interval{Low: 13, High: 17})
	gaps = append(gaps, interval.Interval{Low: 33, High: 34})
	gaps = append(gaps, interval.Interval{Low: 36, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	overlaps = append(overlaps, interval.Interval{Low: 3, High: 4})
	overlaps = append(overlaps, interval.Interval{Low: 5, High: 6})
	overlaps = append(overlaps, interval.Interval{Low: 20, High: 20})
	overlaps = append(overlaps, interval.Interval{Low: 25, High: 28})
	overlaps = append(overlaps, interval.Interval{Low: 30, High: 30})

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 2, High: 7})
	merges = append(merges, interval.Interval{Low: 10, High: 12})
	merges = append(merges, interval.Interval{Low: 18, High: 32})
	merges = append(merges, interval.Interval{Low: 35, High: 35})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

///////////////////////////////////////////////////////////
///   Tests from 100 to 199: low/high are exclusive     ///
///////////////////////////////////////////////////////////

//  no intervals (low/high are exclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo101() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := false
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}

	// calculate expected merges
	merges := []interval.Interval{}

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

// one interval at the beginning (low/high exclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo102() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := false
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: minLow, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: minLow})
	gaps = append(gaps, interval.Interval{Low: 4, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 1, High: 3})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

//  one interval at the end (low/high exclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo103() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := false
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: 8, High: maxHigh}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 8})
	gaps = append(gaps, interval.Interval{Low: maxHigh, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 9, High: 9})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

//  one interval in the middle (low/high exclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo104() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := false
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: 5, High: 8}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 5})
	gaps = append(gaps, interval.Interval{Low: 8, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 6, High: 7})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

//  two intervals in the middle, one inside the other (low/high exclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo105() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := false
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: 2, High: 8}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 4, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 2})
	gaps = append(gaps, interval.Interval{Low: 8, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	overlaps = append(overlaps, interval.Interval{Low: 5, High: 5})

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 3, High: 7})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

// two intervals in the middle (low/high exclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo106() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := false
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: 2, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 6, High: 8}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 2})
	gaps = append(gaps, interval.Interval{Low: 4, High: 6})
	gaps = append(gaps, interval.Interval{Low: 8, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 3, High: 3})
	merges = append(merges, interval.Interval{Low: 7, High: 7})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

//  two intervals in the middle, consecutives  (low/high exclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo107() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := false
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: 2, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 4, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 2})
	gaps = append(gaps, interval.Interval{Low: 4, High: 4})
	gaps = append(gaps, interval.Interval{Low: 6, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 3, High: 3})
	merges = append(merges, interval.Interval{Low: 5, High: 5})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

// two intervals in the middle (low/high exclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo108() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := false
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: 2, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 6, High: 7}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 2})
	gaps = append(gaps, interval.Interval{Low: 6, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 3, High: 5})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

//  two intervals in the middle (low/high exclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo109() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := false
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: 2, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 4, High: 8}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 2})
	gaps = append(gaps, interval.Interval{Low: 8, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	overlaps = append(overlaps, interval.Interval{Low: 5, High: 5})

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 3, High: 7})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

// three intervals (leading, middle and trailing) (low/high exclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo110() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := false
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: minLow, High: 2}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 4, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 8, High: maxHigh}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: minLow})
	gaps = append(gaps, interval.Interval{Low: 2, High: 8})
	gaps = append(gaps, interval.Interval{Low: maxHigh, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 1, High: 1})
	merges = append(merges, interval.Interval{Low: 9, High: 9})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

//  three intervals (in the middle) (low/high exclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo111() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := false
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: 2, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 3, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 5, High: 7}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 2})
	gaps = append(gaps, interval.Interval{Low: 7, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 3, High: 6})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

// complex case with a lot of intervals (low/high exclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo112() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 100
	lowInclusive := false
	highInclusive := false
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: 5, High: 7}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 2, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 35, High: 35}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 3, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 18, High: 20}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 20, High: 30}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 25, High: 28}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: minLow, High: 1}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 30, High: 32}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 10, High: 12}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 2})
	gaps = append(gaps, interval.Interval{Low: 7, High: 10})
	gaps = append(gaps, interval.Interval{Low: 12, High: 18})
	gaps = append(gaps, interval.Interval{Low: 20, High: 20})
	gaps = append(gaps, interval.Interval{Low: 30, High: 30})
	gaps = append(gaps, interval.Interval{Low: 32, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	overlaps = append(overlaps, interval.Interval{Low: 26, High: 27})

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 3, High: 6})
	merges = append(merges, interval.Interval{Low: 11, High: 11})
	merges = append(merges, interval.Interval{Low: 19, High: 19})
	merges = append(merges, interval.Interval{Low: 21, High: 29})
	merges = append(merges, interval.Interval{Low: 31, High: 31})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

/////////////////////////////////////////////////////////////////
///  Tests from 200 to 299: low inclusive and high exclusive  ///
/////////////////////////////////////////////////////////////////

//  no intervals (low inclusive and high exclusive) --> all is a gap  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo201() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := false
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}

	// calculate expected merges
	merges := []interval.Interval{}

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

// one interval at the beginning (low inclusive and high exclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo202() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := false
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: minLow, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: 4, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: minLow, High: 3})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

//  one interval at the end (low inclusive and high exclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo203() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := false
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: 8, High: maxHigh}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 7})
	gaps = append(gaps, interval.Interval{Low: maxHigh, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 8, High: maxHigh - 1})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

//  one interval in the middle (low inclusive and high exclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo204() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := false
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: 5, High: 8}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 4})
	gaps = append(gaps, interval.Interval{Low: 8, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 5, High: 7})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

//  two intervals in the middle, one inside the other (low inclusive and high exclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo205() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := false
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: 2, High: 8}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 4, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 1})
	gaps = append(gaps, interval.Interval{Low: 8, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	overlaps = append(overlaps, interval.Interval{Low: 4, High: 5})

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 2, High: 7})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

// two intervals in the middle (low inclusive and high exclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo206() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := false
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: 2, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 6, High: 8}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 1})
	gaps = append(gaps, interval.Interval{Low: 4, High: 5})
	gaps = append(gaps, interval.Interval{Low: 8, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 2, High: 3})
	merges = append(merges, interval.Interval{Low: 6, High: 7})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

//  two intervals in the middle, consecutives (low inclusive and high exclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo207() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := false
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: 2, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 4, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 1})
	gaps = append(gaps, interval.Interval{Low: 6, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 2, High: 5})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

// two intervals in the middle (low inclusive and high exclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo208() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := false
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: 2, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 6, High: 7}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 1})
	gaps = append(gaps, interval.Interval{Low: 7, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 2, High: 6})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

//  two intervals in the middle (low inclusive and high exclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo209() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := false
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: 2, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 4, High: 8}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 1})
	gaps = append(gaps, interval.Interval{Low: 8, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	overlaps = append(overlaps, interval.Interval{Low: 4, High: 5})

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 2, High: 7})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

// three intervals (leading, middle and trailing) (low inclusive and high exclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo210() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := false
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: minLow, High: 2}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 4, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 8, High: maxHigh}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: 2, High: 7})
	gaps = append(gaps, interval.Interval{Low: maxHigh, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: minLow, High: 1})
	merges = append(merges, interval.Interval{Low: 8, High: 9})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

//  three intervals (in the middle) (low inclusive and high exclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo211() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := false
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: 2, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 3, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 5, High: 7}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 1})
	gaps = append(gaps, interval.Interval{Low: 7, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	overlaps = append(overlaps, interval.Interval{Low: 3, High: 3})
	overlaps = append(overlaps, interval.Interval{Low: 5, High: 5})

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 2, High: 6})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

// complex case with a lot of intervals (low inclusive and high exclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo212() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 100
	lowInclusive := true
	highInclusive := false
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: 5, High: 7}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 2, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 35, High: 35}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 3, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 18, High: 20}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 20, High: 30}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 25, High: 28}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: minLow, High: 1}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 30, High: 32}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 10, High: 12}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: 1, High: 1})
	gaps = append(gaps, interval.Interval{Low: 7, High: 9})
	gaps = append(gaps, interval.Interval{Low: 12, High: 17})
	gaps = append(gaps, interval.Interval{Low: 32, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	overlaps = append(overlaps, interval.Interval{Low: 3, High: 3})
	overlaps = append(overlaps, interval.Interval{Low: 5, High: 5})
	overlaps = append(overlaps, interval.Interval{Low: 25, High: 27})

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: minLow, High: minLow})
	merges = append(merges, interval.Interval{Low: 2, High: 6})
	merges = append(merges, interval.Interval{Low: 10, High: 11})
	merges = append(merges, interval.Interval{Low: 18, High: 31})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

/////////////////////////////////////////////////////////////////
///  Tests from 300 to 399: low exclusive and high inclusive  ///
/////////////////////////////////////////////////////////////////

//  no intervals (low exclusive and high inclusive) --> all is a gap  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo301() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := true
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}

	// calculate expected merges
	merges := []interval.Interval{}

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

// one interval at the beginning (low exclusive and high inclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo302() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := true
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: minLow, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: minLow})
	gaps = append(gaps, interval.Interval{Low: 5, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 1, High: 4})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

//  one interval at the end (low exclusive and high inclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo303() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := true
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: 8, High: maxHigh}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 8})

	// calculate expected overlaps
	overlaps := []interval.Interval{}

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 9, High: maxHigh})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

//  one interval in the middle (low exclusive and high inclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo304() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := true
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: 5, High: 8}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 5})
	gaps = append(gaps, interval.Interval{Low: 9, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 6, High: 8})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

//  two intervals in the middle, one inside the other (low exclusive and high inclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo305() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := true
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: 2, High: 8}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 4, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 2})
	gaps = append(gaps, interval.Interval{Low: 9, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	overlaps = append(overlaps, interval.Interval{Low: 5, High: 6})

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 3, High: 8})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

// two intervals in the middle (low exclusive and high inclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo306() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := true
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: 2, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 6, High: 8}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 2})
	gaps = append(gaps, interval.Interval{Low: 5, High: 6})
	gaps = append(gaps, interval.Interval{Low: 9, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 3, High: 4})
	merges = append(merges, interval.Interval{Low: 7, High: 8})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

//  two intervals in the middle, consecutives (low exclusive and high inclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo307() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := true
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: 2, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 4, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 2})
	gaps = append(gaps, interval.Interval{Low: 7, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 3, High: 6})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

// two intervals in the middle (low exclusive and high inclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo308() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := true
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: 2, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 6, High: 7}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 2})
	gaps = append(gaps, interval.Interval{Low: 8, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 3, High: 7})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

//  two intervals in the middle (low exclusive and high inclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo309() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := true
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: 2, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 4, High: 8}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 2})
	gaps = append(gaps, interval.Interval{Low: 9, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	overlaps = append(overlaps, interval.Interval{Low: 5, High: 6})

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 3, High: 8})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

// three intervals (leading, middle and trailing) (low exclusive and high inclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo310() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := true
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: minLow, High: 2}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 4, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 8, High: maxHigh}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: minLow})
	gaps = append(gaps, interval.Interval{Low: 3, High: 8})

	// calculate expected overlaps
	overlaps := []interval.Interval{}

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 1, High: 2})
	merges = append(merges, interval.Interval{Low: 9, High: maxHigh})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

//  three intervals (in the middle) (low exclusive and high inclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo311() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := true
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: 2, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 3, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 5, High: 7}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 2})
	gaps = append(gaps, interval.Interval{Low: 8, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	overlaps = append(overlaps, interval.Interval{Low: 4, High: 4})
	overlaps = append(overlaps, interval.Interval{Low: 6, High: 6})

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 3, High: 7})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}

// complex case with a lot of intervals (low exclusive and high inclusive)  (selfAdjustMinLow/selfAdjustMaxHigh=false)
func buildIntervalsDemo312() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 100
	lowInclusive := false
	highInclusive := true
	selfAdjustMinLow := false
	selfAdjustMaxHigh := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	// add intervals
	if err := itvls.AddInterval(&interval.Interval{Low: 5, High: 7}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 2, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 35, High: 35}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 3, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 18, High: 20}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 20, High: 30}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 25, High: 28}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: minLow, High: 1}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 30, High: 32}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.AddInterval(&interval.Interval{Low: 10, High: 12}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: minLow})
	gaps = append(gaps, interval.Interval{Low: 2, High: 2})
	gaps = append(gaps, interval.Interval{Low: 8, High: 10})
	gaps = append(gaps, interval.Interval{Low: 13, High: 18})
	gaps = append(gaps, interval.Interval{Low: 33, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	overlaps = append(overlaps, interval.Interval{Low: 4, High: 4})
	overlaps = append(overlaps, interval.Interval{Low: 6, High: 6})
	overlaps = append(overlaps, interval.Interval{Low: 26, High: 28})

	// calculate expected merges
	merges := []interval.Interval{}
	merges = append(merges, interval.Interval{Low: 1, High: 1})
	merges = append(merges, interval.Interval{Low: 3, High: 7})
	merges = append(merges, interval.Interval{Low: 11, High: 12})
	merges = append(merges, interval.Interval{Low: 19, High: 32})

	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps, ExpectedMerges: merges}
}
