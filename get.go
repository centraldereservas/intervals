package interval

func (intvls *intervals) Get() []*Interval {
	// sort intervals (if necessary)
	intvls.Sort()

	// return the intervals sorted
	return intvls.Intervals
}
