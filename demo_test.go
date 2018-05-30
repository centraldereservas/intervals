package interval_test

import "bitbucket.org/differenttravel/interval"

type demo struct {
	Intervals    interval.Intervals
	ExpectedGaps []interval.Interval
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
	return demo{Intervals: itvls, ExpectedGaps: gaps}
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
	itvls.Add(&interval.Interval{Low: minLow, High: 4})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: 5, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
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
	itvls.Add(&interval.Interval{Low: 8, High: maxHigh})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 7})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
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
	itvls.Add(&interval.Interval{Low: 5, High: 8})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 4})
	gaps = append(gaps, interval.Interval{Low: 9, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
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
	itvls.Add(&interval.Interval{Low: 2, High: 8})
	itvls.Add(&interval.Interval{Low: 4, High: 6})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 1})
	gaps = append(gaps, interval.Interval{Low: 9, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
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
	itvls.Add(&interval.Interval{Low: 2, High: 4})
	itvls.Add(&interval.Interval{Low: 6, High: 8})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 1})
	gaps = append(gaps, interval.Interval{Low: 5, High: 5})
	gaps = append(gaps, interval.Interval{Low: 9, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
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
	itvls.Add(&interval.Interval{Low: 2, High: 4})
	itvls.Add(&interval.Interval{Low: 4, High: 6})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 1})
	gaps = append(gaps, interval.Interval{Low: 7, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
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
	itvls.Add(&interval.Interval{Low: 2, High: 6})
	itvls.Add(&interval.Interval{Low: 6, High: 7})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 1})
	gaps = append(gaps, interval.Interval{Low: 8, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
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
	itvls.Add(&interval.Interval{Low: 2, High: 6})
	itvls.Add(&interval.Interval{Low: 4, High: 8})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 1})
	gaps = append(gaps, interval.Interval{Low: 9, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
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
	itvls.Add(&interval.Interval{Low: minLow, High: 2})
	itvls.Add(&interval.Interval{Low: 4, High: 4})
	itvls.Add(&interval.Interval{Low: 8, High: maxHigh})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: 3, High: 3})
	gaps = append(gaps, interval.Interval{Low: 5, High: 7})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
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
	itvls.Add(&interval.Interval{Low: 2, High: 4})
	itvls.Add(&interval.Interval{Low: 3, High: 6})
	itvls.Add(&interval.Interval{Low: 5, High: 7})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 1})
	gaps = append(gaps, interval.Interval{Low: 8, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
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
	itvls.Add(&interval.Interval{Low: 5, High: 7})
	itvls.Add(&interval.Interval{Low: 2, High: 4})
	itvls.Add(&interval.Interval{Low: 35, High: 35})
	itvls.Add(&interval.Interval{Low: 3, High: 6})
	itvls.Add(&interval.Interval{Low: 18, High: 20})
	itvls.Add(&interval.Interval{Low: 20, High: 30})
	itvls.Add(&interval.Interval{Low: 25, High: 28})
	itvls.Add(&interval.Interval{Low: minLow, High: 1})
	itvls.Add(&interval.Interval{Low: 30, High: 32})
	itvls.Add(&interval.Interval{Low: 10, High: 12})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: 8, High: 9})
	gaps = append(gaps, interval.Interval{Low: 13, High: 17})
	gaps = append(gaps, interval.Interval{Low: 33, High: 34})
	gaps = append(gaps, interval.Interval{Low: 36, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
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
	return demo{Intervals: itvls, ExpectedGaps: gaps}
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
	itvls.Add(&interval.Interval{Low: minLow, High: 4})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: minLow})
	gaps = append(gaps, interval.Interval{Low: 4, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
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
	itvls.Add(&interval.Interval{Low: 8, High: maxHigh})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 8})
	gaps = append(gaps, interval.Interval{Low: maxHigh, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
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
	itvls.Add(&interval.Interval{Low: 5, High: 8})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 5})
	gaps = append(gaps, interval.Interval{Low: 8, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
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
	itvls.Add(&interval.Interval{Low: 2, High: 8})
	itvls.Add(&interval.Interval{Low: 4, High: 6})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 2})
	gaps = append(gaps, interval.Interval{Low: 8, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
}

// two intervals in the middle, not overlapping (low/high exclusive)
func buildIntervalsDemo106() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	itvls.Add(&interval.Interval{Low: 2, High: 4})
	itvls.Add(&interval.Interval{Low: 6, High: 8})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 2})
	gaps = append(gaps, interval.Interval{Low: 4, High: 6})
	gaps = append(gaps, interval.Interval{Low: 8, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
}

//  two intervals in the middle, consecutives (not overlapping) (low/high exclusive)
func buildIntervalsDemo107() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	itvls.Add(&interval.Interval{Low: 2, High: 4})
	itvls.Add(&interval.Interval{Low: 4, High: 6})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 2})
	gaps = append(gaps, interval.Interval{Low: 4, High: 4})
	gaps = append(gaps, interval.Interval{Low: 6, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
}

// two intervals in the middle, overlapping by 1 position (low/high exclusive)
func buildIntervalsDemo108() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	itvls.Add(&interval.Interval{Low: 2, High: 6})
	itvls.Add(&interval.Interval{Low: 6, High: 7})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 2})
	gaps = append(gaps, interval.Interval{Low: 6, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
}

//  two intervals in the middle, overlapping by 3 positions (low/high exclusive)
func buildIntervalsDemo109() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	itvls.Add(&interval.Interval{Low: 2, High: 6})
	itvls.Add(&interval.Interval{Low: 4, High: 8})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 2})
	gaps = append(gaps, interval.Interval{Low: 8, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
}

// three intervals (leading, middle and trailing), not overlapping (low/high exclusive)
func buildIntervalsDemo110() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	itvls.Add(&interval.Interval{Low: minLow, High: 2})
	itvls.Add(&interval.Interval{Low: 4, High: 4})
	itvls.Add(&interval.Interval{Low: 8, High: maxHigh})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: minLow})
	gaps = append(gaps, interval.Interval{Low: 2, High: 8})
	gaps = append(gaps, interval.Interval{Low: maxHigh, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
}

//  three intervals (in the middle), overlapping (low/high exclusive)
func buildIntervalsDemo111() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	itvls.Add(&interval.Interval{Low: 2, High: 4})
	itvls.Add(&interval.Interval{Low: 3, High: 6})
	itvls.Add(&interval.Interval{Low: 5, High: 7})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 2})
	gaps = append(gaps, interval.Interval{Low: 7, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
}

// complex case with a lot of intervals and overlapping (low/high exclusive)
func buildIntervalsDemo112() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 100
	lowInclusive := false
	highInclusive := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	itvls.Add(&interval.Interval{Low: 5, High: 7})
	itvls.Add(&interval.Interval{Low: 2, High: 4})
	itvls.Add(&interval.Interval{Low: 35, High: 35})
	itvls.Add(&interval.Interval{Low: 3, High: 6})
	itvls.Add(&interval.Interval{Low: 18, High: 20})
	itvls.Add(&interval.Interval{Low: 20, High: 30})
	itvls.Add(&interval.Interval{Low: 25, High: 28})
	itvls.Add(&interval.Interval{Low: minLow, High: 1})
	itvls.Add(&interval.Interval{Low: 30, High: 32})
	itvls.Add(&interval.Interval{Low: 10, High: 12})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 2})
	gaps = append(gaps, interval.Interval{Low: 7, High: 10})
	gaps = append(gaps, interval.Interval{Low: 12, High: 18})
	gaps = append(gaps, interval.Interval{Low: 20, High: 20})
	gaps = append(gaps, interval.Interval{Low: 30, High: 30})
	gaps = append(gaps, interval.Interval{Low: 32, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
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
	return demo{Intervals: itvls, ExpectedGaps: gaps}
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
	itvls.Add(&interval.Interval{Low: minLow, High: 4})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: 4, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
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
	itvls.Add(&interval.Interval{Low: 8, High: maxHigh})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 7})
	gaps = append(gaps, interval.Interval{Low: maxHigh, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
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
	itvls.Add(&interval.Interval{Low: 5, High: 8})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 4})
	gaps = append(gaps, interval.Interval{Low: 8, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
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
	itvls.Add(&interval.Interval{Low: 2, High: 8})
	itvls.Add(&interval.Interval{Low: 4, High: 6})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 1})
	gaps = append(gaps, interval.Interval{Low: 8, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
}

// two intervals in the middle, not overlapping (low inclusive and high exclusive)
func buildIntervalsDemo206() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	itvls.Add(&interval.Interval{Low: 2, High: 4})
	itvls.Add(&interval.Interval{Low: 6, High: 8})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 1})
	gaps = append(gaps, interval.Interval{Low: 4, High: 5})
	gaps = append(gaps, interval.Interval{Low: 8, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
}

//  two intervals in the middle, consecutives (not overlapping) (low inclusive and high exclusive)
func buildIntervalsDemo207() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	itvls.Add(&interval.Interval{Low: 2, High: 4})
	itvls.Add(&interval.Interval{Low: 4, High: 6})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 1})
	gaps = append(gaps, interval.Interval{Low: 6, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
}

// two intervals in the middle, overlapping by 1 position (low inclusive and high exclusive)
func buildIntervalsDemo208() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	itvls.Add(&interval.Interval{Low: 2, High: 6})
	itvls.Add(&interval.Interval{Low: 6, High: 7})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 1})
	gaps = append(gaps, interval.Interval{Low: 7, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
}

//  two intervals in the middle, overlapping by 3 positions (low inclusive and high exclusive)
func buildIntervalsDemo209() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	itvls.Add(&interval.Interval{Low: 2, High: 6})
	itvls.Add(&interval.Interval{Low: 4, High: 8})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 1})
	gaps = append(gaps, interval.Interval{Low: 8, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
}

// three intervals (leading, middle and trailing), not overlapping (low inclusive and high exclusive)
func buildIntervalsDemo210() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	itvls.Add(&interval.Interval{Low: minLow, High: 2})
	itvls.Add(&interval.Interval{Low: 4, High: 4})
	itvls.Add(&interval.Interval{Low: 8, High: maxHigh})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: 2, High: 7})
	gaps = append(gaps, interval.Interval{Low: maxHigh, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
}

//  three intervals (in the middle), overlapping (low inclusive and high exclusive)
func buildIntervalsDemo211() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := true
	highInclusive := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	itvls.Add(&interval.Interval{Low: 2, High: 4})
	itvls.Add(&interval.Interval{Low: 3, High: 6})
	itvls.Add(&interval.Interval{Low: 5, High: 7})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 1})
	gaps = append(gaps, interval.Interval{Low: 7, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
}

// complex case with a lot of intervals and overlapping (low inclusive and high exclusive)
func buildIntervalsDemo212() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 100
	lowInclusive := true
	highInclusive := false
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	itvls.Add(&interval.Interval{Low: 5, High: 7})
	itvls.Add(&interval.Interval{Low: 2, High: 4})
	itvls.Add(&interval.Interval{Low: 35, High: 35})
	itvls.Add(&interval.Interval{Low: 3, High: 6})
	itvls.Add(&interval.Interval{Low: 18, High: 20})
	itvls.Add(&interval.Interval{Low: 20, High: 30})
	itvls.Add(&interval.Interval{Low: 25, High: 28})
	itvls.Add(&interval.Interval{Low: minLow, High: 1})
	itvls.Add(&interval.Interval{Low: 30, High: 32})
	itvls.Add(&interval.Interval{Low: 10, High: 12})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: 1, High: 1})
	gaps = append(gaps, interval.Interval{Low: 7, High: 9})
	gaps = append(gaps, interval.Interval{Low: 12, High: 17})
	gaps = append(gaps, interval.Interval{Low: 32, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
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
	return demo{Intervals: itvls, ExpectedGaps: gaps}
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
	itvls.Add(&interval.Interval{Low: minLow, High: 4})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: minLow})
	gaps = append(gaps, interval.Interval{Low: 5, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
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
	itvls.Add(&interval.Interval{Low: 8, High: maxHigh})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 8})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
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
	itvls.Add(&interval.Interval{Low: 5, High: 8})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 5})
	gaps = append(gaps, interval.Interval{Low: 9, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
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
	itvls.Add(&interval.Interval{Low: 2, High: 8})
	itvls.Add(&interval.Interval{Low: 4, High: 6})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 2})
	gaps = append(gaps, interval.Interval{Low: 9, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
}

// two intervals in the middle, not overlapping (low exclusive and high inclusive)
func buildIntervalsDemo306() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := true
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	itvls.Add(&interval.Interval{Low: 2, High: 4})
	itvls.Add(&interval.Interval{Low: 6, High: 8})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 2})
	gaps = append(gaps, interval.Interval{Low: 5, High: 6})
	gaps = append(gaps, interval.Interval{Low: 9, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
}

//  two intervals in the middle, consecutives (not overlapping) (low exclusive and high inclusive)
func buildIntervalsDemo307() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := true
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	itvls.Add(&interval.Interval{Low: 2, High: 4})
	itvls.Add(&interval.Interval{Low: 4, High: 6})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 2})
	gaps = append(gaps, interval.Interval{Low: 7, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
}

// two intervals in the middle, overlapping by 1 position (low exclusive and high inclusive)
func buildIntervalsDemo308() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := true
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	itvls.Add(&interval.Interval{Low: 2, High: 6})
	itvls.Add(&interval.Interval{Low: 6, High: 7})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 2})
	gaps = append(gaps, interval.Interval{Low: 8, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
}

//  two intervals in the middle, overlapping by 3 positions (low exclusive and high inclusive)
func buildIntervalsDemo309() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := true
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	itvls.Add(&interval.Interval{Low: 2, High: 6})
	itvls.Add(&interval.Interval{Low: 4, High: 8})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 2})
	gaps = append(gaps, interval.Interval{Low: 9, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
}

// three intervals (leading, middle and trailing), not overlapping (low exclusive and high inclusive)
func buildIntervalsDemo310() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := true
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	itvls.Add(&interval.Interval{Low: minLow, High: 2})
	itvls.Add(&interval.Interval{Low: 4, High: 4})
	itvls.Add(&interval.Interval{Low: 8, High: maxHigh})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: minLow})
	gaps = append(gaps, interval.Interval{Low: 3, High: 8})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
}

//  three intervals (in the middle), overlapping (low exclusive and high inclusive)
func buildIntervalsDemo311() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 10
	lowInclusive := false
	highInclusive := true
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	itvls.Add(&interval.Interval{Low: 2, High: 4})
	itvls.Add(&interval.Interval{Low: 3, High: 6})
	itvls.Add(&interval.Interval{Low: 5, High: 7})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: 2})
	gaps = append(gaps, interval.Interval{Low: 8, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
}

// complex case with a lot of intervals and overlapping (low exclusive and high inclusive)
func buildIntervalsDemo312() demo {
	// initialize Intervals
	minLow := 0
	maxHigh := 100
	lowInclusive := false
	highInclusive := true
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add intervals
	itvls.Add(&interval.Interval{Low: 5, High: 7})
	itvls.Add(&interval.Interval{Low: 2, High: 4})
	itvls.Add(&interval.Interval{Low: 35, High: 35})
	itvls.Add(&interval.Interval{Low: 3, High: 6})
	itvls.Add(&interval.Interval{Low: 18, High: 20})
	itvls.Add(&interval.Interval{Low: 20, High: 30})
	itvls.Add(&interval.Interval{Low: 25, High: 28})
	itvls.Add(&interval.Interval{Low: minLow, High: 1})
	itvls.Add(&interval.Interval{Low: 30, High: 32})
	itvls.Add(&interval.Interval{Low: 10, High: 12})

	// calculate expected gaps
	gaps := []interval.Interval{}
	gaps = append(gaps, interval.Interval{Low: minLow, High: minLow})
	gaps = append(gaps, interval.Interval{Low: 2, High: 2})
	gaps = append(gaps, interval.Interval{Low: 8, High: 10})
	gaps = append(gaps, interval.Interval{Low: 13, High: 18})
	gaps = append(gaps, interval.Interval{Low: 33, High: maxHigh})
	return demo{Intervals: itvls, ExpectedGaps: gaps}
}
