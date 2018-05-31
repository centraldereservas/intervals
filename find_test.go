package interval_test

import (
	"fmt"
	"testing"

	"bitbucket.org/differenttravel/interval"
)

func initIntervalsForDemo001() interval.Intervals {
	// initialize Intervals
	minLow := 0
	maxHigh := 100
	lowInclusive := true
	highInclusive := true
	itvls := interval.NewIntervals(minLow, maxHigh, lowInclusive, highInclusive)

	// add new intervals
	if err := itvls.Add(&interval.Interval{Low: 5, High: 7}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	if err := itvls.Add(&interval.Interval{Low: 2, High: 4}); err != nil {
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
	if err := itvls.Add(&interval.Interval{Low: 30, High: 32}); err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	return itvls
}

//  matches for value=2
func buildFindDemo001() demo {
	itvls := initIntervalsForDemo001()
	matches := []interval.Interval{}
	matches = append(matches, interval.Interval{Low: 2, High: 4})
	return demo{Intervals: itvls, ExpectedFindMatches: matches, ValueToFind: 2}
}

//  matches for value=4
func buildFindDemo002() demo {
	itvls := initIntervalsForDemo001()
	matches := []interval.Interval{}
	matches = append(matches, interval.Interval{Low: 2, High: 4})
	matches = append(matches, interval.Interval{Low: 3, High: 6})
	return demo{Intervals: itvls, ExpectedFindMatches: matches, ValueToFind: 4}
}

func TestFindIntervalsForValue(t *testing.T) {
	demo001 := buildFindDemo001()
	demo002 := buildFindDemo002()

	tt := []struct {
		name            string
		valueToFind     int
		intvls          interval.Intervals
		expectedMatches []interval.Interval
	}{
		{name: "demo001", valueToFind: demo001.ValueToFind, intvls: demo001.Intervals, expectedMatches: demo001.ExpectedFindMatches},
		{name: "demo002", valueToFind: demo002.ValueToFind, intvls: demo002.Intervals, expectedMatches: demo002.ExpectedFindMatches},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			matches := tc.intvls.FindIntervalsForValue(tc.valueToFind)
			unexpectedLength := false
			if len(matches) != len(tc.expectedMatches) {
				t.Errorf("unexpected gaps length: expected %d; got %d", len(tc.expectedMatches), len(matches))
				unexpectedLength = true
			}
			if !unexpectedLength {
				for i, m := range matches {
					if m.Low != tc.expectedMatches[i].Low || m.High != tc.expectedMatches[i].High {
						t.Errorf("unexpected gap[%d]: expected %v; got %v", i, tc.expectedMatches[i], m)
					}
				}
			}
		})
	}
}