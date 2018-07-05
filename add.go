package intervals

import "math"

func (intvls *intervals) Add(low, high int, obj interface{}) error {
	itvl := &Interval{
		Low:    low,
		High:   high,
		Object: obj,
	}
	return intvls.AddInterval(itvl)
}

func (intvls *intervals) AddInterval(itvl *Interval) error {

	// the first time an interval is added, we check if a self adjustment is programmed to update some control variables
	if len(intvls.Intervals) == 0 {
		if intvls.SelfAdjustMinLow {
			intvls.MinLow = math.MaxInt64
		}
		if intvls.SelfAdjustMaxHigh {
			intvls.MaxHigh = math.MinInt64
		}
	}

	low := intvls.getInclusiveLow(itvl.Low)
	high := intvls.getInclusiveHigh(itvl.High)

	err := intvls.checkValidInterval(low, high)
	if err != nil {
		return err
	}

	if intvls.SelfAdjustMaxHigh && high > intvls.MaxHigh {
		intvls.MaxHigh = high
	}
	if intvls.SelfAdjustMinLow && low < intvls.MinLow {
		intvls.MinLow = low
	}

	intvls.Intervals = append(intvls.Intervals, itvl)
	intvls.reset()
	return nil
}

func (intvls *intervals) reset() {
	intvls.Sorted = false
	intvls.GapsList = nil
	intvls.OverlappedList = nil
	intvls.MergeList = nil
}
