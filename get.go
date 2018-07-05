package intervals

func (intvls *intervals) GetIntervals() []*Interval {
	// sort intervals (if necessary)
	intvls.Sort()

	// return the intervals sorted
	return intvls.Intervals
}
