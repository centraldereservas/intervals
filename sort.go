package interval

import "sort"

func (intvls *intervals) Sort() {
	// sort the intervals slice just if necessary
	if !intvls.Sorted {
		sort.Sort(ByLow(intvls.Intervals))
	}
	intvls.Sorted = true
}
