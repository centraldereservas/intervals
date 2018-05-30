package interval

import (
	"math"
)

const (
	defaultMinLow        = 0
	defaultMaxHigh       = math.MaxInt64
	defaultLowInclusive  = true
	defaultHighInclusive = true
)

// Intervals is an interface to handle Interval structures discovering the existence of gaps or overlays
type Intervals interface {
	Add(itvl *Interval)

	// Sort sorts the intervals list by the Low property (ascending)
	Sort()

	// HasGaps returns true if exists gaps for the introduced intervals between MinLow and MaxHigh
	HasGaps() bool

	// Get returns the interval list
	Get() []*Interval

	// Gaps returns the interval gaps
	Gaps() []*Interval

	// Merge fussion together overlapped intervals
	Merge() []*Interval

	// HasOverlapped returns true if exists overlapping for the introduced intervals
	HasOverlapped() bool

	// Overlapped returns the overlapped intervals
	Overlapped() []*Interval

	// FindIntervalsForValue returns all the intervals which contains the passed value
	FindIntervalsForValue(value int) []*Interval

	// Report creates a report of the interval sequence
	Report() string

	// IsLowInclusive indicates if the Low part of the interval is included, e. g. (3,5) --> the 3 is included as part of the interval
	IsLowInclusive() bool

	// IsHighInclusive indicates if the High part of the interval is included, e. g. (3,5) --> the 5 is included as part of the interval
	IsHighInclusive() bool
}

// intervals implements Intervals interface
type intervals struct {
	Intervals      []*Interval
	GapsList       []*Interval
	OverlappedList []*Interval
	MergeList      []*Interval
	MinLow         int
	MaxHigh        int
	Sorted         bool
	LowInclusive   bool
	HighInclusive  bool
}

// NewIntervalsDefault is a constructor that returns an instance of the Intervals interface with default values
func NewIntervalsDefault() Intervals {
	return NewIntervals(defaultMinLow, defaultMaxHigh, defaultLowInclusive, defaultHighInclusive)
}

// NewIntervals is a constructor that returns an instance of the Intervals interface
func NewIntervals(minLow int, maxHigh int, lowInclusive bool, highInclusive bool) Intervals {
	return &intervals{
		MinLow:        minLow,
		MaxHigh:       maxHigh,
		Intervals:     []*Interval{},
		Sorted:        false,
		LowInclusive:  lowInclusive,
		HighInclusive: highInclusive,
	}
}

func (intvls *intervals) IsLowInclusive() bool {
	return intvls.LowInclusive
}

func (intvls *intervals) IsHighInclusive() bool {
	return intvls.HighInclusive
}
