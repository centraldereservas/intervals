package interval_test

import (
	"testing"

	"bitbucket.org/differenttravel/interval"
)

func TestInsert(t *testing.T) {
	itvls := interval.NewIntervals(0, 100)

	itvls.Add(&interval.Interval{Low: 5, High: 7})
	itvls.Add(&interval.Interval{Low: 2, High: 4})
	itvls.Add(&interval.Interval{Low: 3, High: 6})
	itvls.Add(&interval.Interval{Low: 18, High: 20})
	itvls.Add(&interval.Interval{Low: 20, High: 30})
	itvls.Add(&interval.Interval{Low: 25, High: 28})
	itvls.Add(&interval.Interval{Low: 30, High: 32})

	tt := []struct {
		name  string
		itvls interval.Intervals
	}{
		{name: "normal case", itvls: itvls},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.itvls.Print())
		})
	}
}
