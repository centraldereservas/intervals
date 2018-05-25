package interval_test

import (
	"fmt"
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
			t.Log(tc.itvls.Report())
		})
	}
}

func TestFindIntervalsForValue(t *testing.T) {
	itvls := interval.NewIntervals(0, 100)

	itvls.Add(&interval.Interval{Low: 5, High: 7})
	itvls.Add(&interval.Interval{Low: 2, High: 4})
	itvls.Add(&interval.Interval{Low: 3, High: 6})
	itvls.Add(&interval.Interval{Low: 18, High: 20})
	itvls.Add(&interval.Interval{Low: 20, High: 30})
	itvls.Add(&interval.Interval{Low: 25, High: 28})
	itvls.Add(&interval.Interval{Low: 30, High: 32})

	valueToFind1 := 2
	matches1 := itvls.FindIntervalsForValue(valueToFind1)
	matches1Txt := fmt.Sprintf("\nFind(value=%d)={", valueToFind1)
	for _, m := range matches1 {
		matches1Txt += fmt.Sprintf("%v,", m)
	}
	matches1Txt += "}"
	t.Logf(matches1Txt)

	valueToFind2 := 4
	matches2 := itvls.FindIntervalsForValue(4)
	matches2Txt := fmt.Sprintf("\nFind(value=%d)={", valueToFind2)
	for _, m := range matches2 {
		matches2Txt += fmt.Sprintf("%v,", m)
	}
	matches2Txt += "}"
	t.Logf(matches2Txt)
}
