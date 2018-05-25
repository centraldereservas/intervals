package interval

import "sort"

func (intvls *intervals) Sort() {
	if !intvls.Sorted {
		sort.Sort(ByLow(intvls.Intervals))
	}
	intvls.Sorted = true
}
