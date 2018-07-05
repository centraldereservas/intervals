package intervals_test

import (
	"fmt"
	"testing"

	"github.com/centraldereservas/intervals"
)

func TestReport(t *testing.T) {
	minLow := 0
	maxHigh := 100
	lowInclusive := true
	highInclusive := true
	selfAdjustMinLow := false
	selfAdjustMaxHigh := true
	itvls := intervals.New(minLow, maxHigh, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh)

	var err error
	err = itvls.AddInterval(&intervals.Interval{Low: 5, High: 7})
	if err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	err = itvls.AddInterval(&intervals.Interval{Low: 2, High: 4})
	if err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	err = itvls.AddInterval(&intervals.Interval{Low: 3, High: 6})
	if err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	err = itvls.AddInterval(&intervals.Interval{Low: 18, High: 20})
	if err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	err = itvls.AddInterval(&intervals.Interval{Low: 20, High: 30})
	if err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	err = itvls.AddInterval(&intervals.Interval{Low: 25, High: 28})
	if err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	err = itvls.AddInterval(&intervals.Interval{Low: 30, High: 32})
	if err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	tt := []struct {
		name  string
		itvls intervals.Intervals
	}{
		{name: "normal case", itvls: itvls},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// this will call report() which implements Springer interface
			t.Log(tc.itvls)
		})
	}
}
