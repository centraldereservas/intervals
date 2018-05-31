package interval_test

import (
	"fmt"
	"testing"

	"bitbucket.org/differenttravel/interval"
)

func TestReport(t *testing.T) {
	itvls := interval.NewIntervals(0, 100, true, true)

	var err error
	err = itvls.AddInterval(&interval.Interval{Low: 5, High: 7})
	if err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	err = itvls.AddInterval(&interval.Interval{Low: 2, High: 4})
	if err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	err = itvls.AddInterval(&interval.Interval{Low: 3, High: 6})
	if err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	err = itvls.AddInterval(&interval.Interval{Low: 18, High: 20})
	if err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	err = itvls.AddInterval(&interval.Interval{Low: 20, High: 30})
	if err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	err = itvls.AddInterval(&interval.Interval{Low: 25, High: 28})
	if err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}
	err = itvls.AddInterval(&interval.Interval{Low: 30, High: 32})
	if err != nil {
		fmt.Printf("invalid interval discarded: %v\n", err)
	}

	tt := []struct {
		name  string
		itvls interval.Intervals
	}{
		{name: "normal case", itvls: itvls},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.itvls.Report())
		})
	}
}
