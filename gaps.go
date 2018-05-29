package interval

import "fmt"

func (intvls *intervals) HasGaps() bool {
	intvls.Gaps()
	if intvls.GapsList != nil && len(intvls.GapsList) > 0 {
		return true
	}
	return false
}

func (intvls *intervals) Gaps() []*Interval {
	if intvls.GapsList == nil {
		intvls.GapsList = intvls.calculateGaps()
	}
	return intvls.GapsList
}

func (intvls *intervals) calculateGaps() []*Interval {
	list := []*Interval{}
	if len(intvls.Intervals) == 0 {
		return list
	}

	// sort intervals (if necessary)
	intvls.Sort()

	gapThreshold := intvls.MinLow
	for _, intvl := range intvls.Intervals {
		// convert if necessary exclusive low/high values into inclusive ones
		low, high, err := intvls.getInclusives(intvl.Low, intvl.High)
		if err != nil {
			fmt.Printf("calculateGaps - unable to get inclusives: %v", err)
			continue
		}

		// if the current Low is higher than the last maximal High, means that there is a gap so we add this gap to the list
		if low > gapThreshold {
			list = append(list, &Interval{Low: gapThreshold, High: low - 1})
		}

		// update if necessary the threshold for the next gap
		if high >= gapThreshold {
			gapThreshold = high + 1
		}
	}

	// if intvls.Intervals haven't covered all the range until the end, we need to fill the rest until the end as a gap
	if gapThreshold < intvls.MaxHigh {
		list = append(list, &Interval{Low: gapThreshold, High: intvls.MaxHigh})
	}
	return list
}
