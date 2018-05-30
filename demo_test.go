package interval_test

import (
	"fmt"

	"bitbucket.org/differenttravel/interval"
)

type demo struct {
	Intervals        interval.Intervals
	ExpectedGaps     []interval.Interval
	ExpectedOverlaps []interval.Interval
}

///////////////////////////////////////////////////////////
///   Tests from 000 to 099: low/high are inclusive     ///
///////////////////////////////////////////////////////////

//  no intervals (low/high inclusive) --> all is a gap
func buildIntervalsDemo001() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := true
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

// one interval at the beginning (low/high inclusive)
func buildIntervalsDemo002() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := true
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: minLow, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: 5, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

//  one interval at the end (low/high inclusive)
func buildIntervalsDemo003() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := true
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: 8, High: maxHigh}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 7})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

//  one interval in the middle (low/high inclusive)
func buildIntervalsDemo004() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := true
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: 5, High: 8}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 4})
	gaps = append(gaps, interval.Interval{Low: 9, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

//  two intervals in the middle, one inside the other (low/high inclusive)
func buildIntervalsDemo005() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := true
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: 2, High: 8}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 4, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 1})
	gaps = append(gaps, interval.Interval{Low: 9, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	overlaps = append(overlaps, interval.Interval{Low: 4, High: 6})
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

// two intervals in the middle, not overlapping (low/high inclusive)
func buildIntervalsDemo006() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := true
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: 2, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 6, High: 8}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 1})
	gaps = append(gaps, interval.Interval{Low: 5, High: 5})
	gaps = append(gaps, interval.Interval{Low: 9, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

//  two intervals in the middle, consecutives (not overlapping) (low/high inclusive)
func buildIntervalsDemo007() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := true
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: 2, High: 3}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 4, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 1})
	gaps = append(gaps, interval.Interval{Low: 7, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

// two intervals in the middle, overlapping by 1 position (low/high inclusive)
func buildIntervalsDemo008() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := true
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: 2, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 6, High: 7}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 1})
	gaps = append(gaps, interval.Interval{Low: 8, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	overlaps = append(overlaps, interval.Interval{Low: 6, High: 6})
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

//  two intervals in the middle, overlapping by 3 positions (low/high inclusive)
func buildIntervalsDemo009() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := true
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: 2, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 4, High: 8}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 1})
	gaps = append(gaps, interval.Interval{Low: 9, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	overlaps = append(overlaps, interval.Interval{Low: 4, High: 6})
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

// three intervals (leading, middle and trailing), not overlapping (low/high inclusive)
func buildIntervalsDemo010() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := true
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: minLow, High: 2}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 4, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 8, High: maxHigh}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: 3, High: 3})
	gaps = append(gaps, interval.Interval{Low: 5, High: 7})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

//  three intervals (in the middle), overlapping (low/high inclusive)
func buildIntervalsDemo011() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := true
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: 2, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 3, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 5, High: 7}); err != nil {
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
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

// complex case with a lot of intervals and overlapping (low/high inclusive)
func buildIntervalsDemo012() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 100
	lowInclusive := true
	highInclusive := true
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: 5, High: 7}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 2, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 35, High: 35}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 3, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 18, High: 20}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 20, High: 30}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 25, High: 28}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: minLow, High: 1}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 30, High: 32}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 10, High: 12}); err != nil {
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
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

///////////////////////////////////////////////////////////
///   Tests from 100 to 199: low/high are exclusive     ///
///////////////////////////////////////////////////////////

//  no intervals (low/high are exclusive)
func buildIntervalsDemo101() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

// one interval at the beginning (low/high exclusive)
func buildIntervalsDemo102() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: minLow, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: minLow})
	gaps = append(gaps, interval.Interval{Low: 4, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

//  one interval at the end (low/high exclusive)
func buildIntervalsDemo103() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: 8, High: maxHigh}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 8})
	gaps = append(gaps, interval.Interval{Low: maxHigh, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

//  one interval in the middle (low/high exclusive)
func buildIntervalsDemo104() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: 5, High: 8}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 5})
	gaps = append(gaps, interval.Interval{Low: 8, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

//  two intervals in the middle, one inside the other (low/high exclusive)
func buildIntervalsDemo105() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: 2, High: 8}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 4, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 2})
	gaps = append(gaps, interval.Interval{Low: 8, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	overlaps = append(overlaps, interval.Interval{Low: 5, High: 5})
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

// two intervals in the middle (low/high exclusive)
func buildIntervalsDemo106() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: 2, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 6, High: 8}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 2})
	gaps = append(gaps, interval.Interval{Low: 4, High: 6})
	gaps = append(gaps, interval.Interval{Low: 8, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

//  two intervals in the middle, consecutives  (low/high exclusive)
func buildIntervalsDemo107() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: 2, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 4, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 2})
	gaps = append(gaps, interval.Interval{Low: 4, High: 4})
	gaps = append(gaps, interval.Interval{Low: 6, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

// two intervals in the middle (low/high exclusive)
func buildIntervalsDemo108() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: 2, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 6, High: 7}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 2})
	gaps = append(gaps, interval.Interval{Low: 6, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

//  two intervals in the middle (low/high exclusive)
func buildIntervalsDemo109() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: 2, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 4, High: 8}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 2})
	gaps = append(gaps, interval.Interval{Low: 8, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	overlaps = append(overlaps, interval.Interval{Low: 5, High: 5})
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

// three intervals (leading, middle and trailing) (low/high exclusive)
func buildIntervalsDemo110() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: minLow, High: 2}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 4, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 8, High: maxHigh}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: minLow})
	gaps = append(gaps, interval.Interval{Low: 2, High: 8})
	gaps = append(gaps, interval.Interval{Low: maxHigh, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

//  three intervals (in the middle) (low/high exclusive)
func buildIntervalsDemo111() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: 2, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 3, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 5, High: 7}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 2})
	gaps = append(gaps, interval.Interval{Low: 7, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

// complex case with a lot of intervals (low/high exclusive)
func buildIntervalsDemo112() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 100
	lowInclusive := false
	highInclusive := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: 5, High: 7}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 2, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 35, High: 35}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 3, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 18, High: 20}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 20, High: 30}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 25, High: 28}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: minLow, High: 1}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 30, High: 32}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 10, High: 12}); err != nil {
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
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

/////////////////////////////////////////////////////////////////
///  Tests from 200 to 299: low inclusive and high exclusive  ///
/////////////////////////////////////////////////////////////////

//  no intervals (low inclusive and high exclusive) --> all is a gap
func buildIntervalsDemo201() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

// one interval at the beginning (low inclusive and high exclusive)
func buildIntervalsDemo202() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: minLow, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: 4, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

//  one interval at the end (low inclusive and high exclusive)
func buildIntervalsDemo203() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: 8, High: maxHigh}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 7})
	gaps = append(gaps, interval.Interval{Low: maxHigh, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

//  one interval in the middle (low inclusive and high exclusive)
func buildIntervalsDemo204() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: 5, High: 8}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 4})
	gaps = append(gaps, interval.Interval{Low: 8, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

//  two intervals in the middle, one inside the other (low inclusive and high exclusive)
func buildIntervalsDemo205() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: 2, High: 8}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 4, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 1})
	gaps = append(gaps, interval.Interval{Low: 8, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	overlaps = append(overlaps, interval.Interval{Low: 4, High: 5})
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

// two intervals in the middle (low inclusive and high exclusive)
func buildIntervalsDemo206() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: 2, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 6, High: 8}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 1})
	gaps = append(gaps, interval.Interval{Low: 4, High: 5})
	gaps = append(gaps, interval.Interval{Low: 8, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

//  two intervals in the middle, consecutives (low inclusive and high exclusive)
func buildIntervalsDemo207() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: 2, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 4, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 1})
	gaps = append(gaps, interval.Interval{Low: 6, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

// two intervals in the middle (low inclusive and high exclusive)
func buildIntervalsDemo208() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: 2, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 6, High: 7}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 1})
	gaps = append(gaps, interval.Interval{Low: 7, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

//  two intervals in the middle (low inclusive and high exclusive)
func buildIntervalsDemo209() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: 2, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 4, High: 8}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 1})
	gaps = append(gaps, interval.Interval{Low: 8, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	overlaps = append(overlaps, interval.Interval{Low: 4, High: 5})
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

// three intervals (leading, middle and trailing) (low inclusive and high exclusive)
func buildIntervalsDemo210() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: minLow, High: 2}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 4, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 8, High: maxHigh}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: 2, High: 7})
	gaps = append(gaps, interval.Interval{Low: maxHigh, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

//  three intervals (in the middle) (low inclusive and high exclusive)
func buildIntervalsDemo211() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: 2, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 3, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 5, High: 7}); err != nil {
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
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

// complex case with a lot of intervals (low inclusive and high exclusive)
func buildIntervalsDemo212() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 100
	lowInclusive := true
	highInclusive := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: 5, High: 7}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 2, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 35, High: 35}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 3, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 18, High: 20}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 20, High: 30}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 25, High: 28}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: minLow, High: 1}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 30, High: 32}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 10, High: 12}); err != nil {
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
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

/////////////////////////////////////////////////////////////////
///  Tests from 300 to 399: low exclusive and high inclusive  ///
/////////////////////////////////////////////////////////////////

//  no intervals (low exclusive and high inclusive) --> all is a gap
func buildIntervalsDemo301() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := true
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

// one interval at the beginning (low exclusive and high inclusive)
func buildIntervalsDemo302() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := true
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: minLow, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: minLow})
	gaps = append(gaps, interval.Interval{Low: 5, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

//  one interval at the end (low exclusive and high inclusive)
func buildIntervalsDemo303() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := true
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: 8, High: maxHigh}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 8})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

//  one interval in the middle (low exclusive and high inclusive)
func buildIntervalsDemo304() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := true
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: 5, High: 8}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 5})
	gaps = append(gaps, interval.Interval{Low: 9, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

//  two intervals in the middle, one inside the other (low exclusive and high inclusive)
func buildIntervalsDemo305() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := true
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: 2, High: 8}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 4, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 2})
	gaps = append(gaps, interval.Interval{Low: 9, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	overlaps = append(overlaps, interval.Interval{Low: 5, High: 6})
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

// two intervals in the middle (low exclusive and high inclusive)
func buildIntervalsDemo306() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := true
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: 2, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 6, High: 8}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 2})
	gaps = append(gaps, interval.Interval{Low: 5, High: 6})
	gaps = append(gaps, interval.Interval{Low: 9, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

//  two intervals in the middle, consecutives (low exclusive and high inclusive)
func buildIntervalsDemo307() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := true
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: 2, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 4, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 2})
	gaps = append(gaps, interval.Interval{Low: 7, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

// two intervals in the middle (low exclusive and high inclusive)
func buildIntervalsDemo308() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := true
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: 2, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 6, High: 7}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 2})
	gaps = append(gaps, interval.Interval{Low: 8, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

//  two intervals in the middle (low exclusive and high inclusive)
func buildIntervalsDemo309() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := true
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: 2, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 4, High: 8}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 2})
	gaps = append(gaps, interval.Interval{Low: 9, High: maxHigh})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	overlaps = append(overlaps, interval.Interval{Low: 5, High: 6})
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

// three intervals (leading, middle and trailing) (low exclusive and high inclusive)
func buildIntervalsDemo310() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := true
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: minLow, High: 2}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 4, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 8, High: maxHigh}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: minLow})
	gaps = append(gaps, interval.Interval{Low: 3, High: 8})

	// calculate expected overlaps
	overlaps := []interval.Interval{}
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

//  three intervals (in the middle) (low exclusive and high inclusive)
func buildIntervalsDemo311() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := true
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: 2, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 3, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 5, High: 7}); err != nil {
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
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}

// complex case with a lot of intervals (low exclusive and high inclusive)
func buildIntervalsDemo312() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 100
	lowInclusive := false
	highInclusive := true
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	if err := itvls.Add(&interval.Interval{Low: 5, High: 7}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 2, High: 4}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 35, High: 35}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 3, High: 6}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 18, High: 20}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 20, High: 30}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 25, High: 28}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: minLow, High: 1}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 30, High: 32}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 10, High: 12}); err != nil {
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
	return demo{Intervals: itvls, ExpectedGaps: gaps, ExpectedOverlaps: overlaps}
}
