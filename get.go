package interval

func (intvls *intervals) Get() []*Interval {
	intvls.Sort()
	return intvls.Intervals
}
