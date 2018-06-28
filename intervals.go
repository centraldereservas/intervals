package interval

import (
	"math"
)

const (
	defaultMinLow            = 0
	defaultMaxHigh           = math.MaxInt64
	defaultLowInclusive      = true
	defaultHighInclusive     = true
	defaultSelfAdjustMinLow  = false
	defaultSelfAdjustMaxHigh = true
)

// Intervals is an interface to handle Interval structures discovering the existence of gaps or overlays
type Intervals interface {
	// Add creates an interval from the input parameters and call AddInterval
	Add(low, high int, obj interface{}) error

	// AddInterval appends a new interval to the list. If the interval range (low, high) is invalid, it returns an error
	AddInterval(itvl *Interval) error

	// HasGaps returns true if exists gaps for the introduced intervals between MinLow and MaxHigh
	HasGaps() bool

	// Get returns the interval list
	GetIntervals() []*Interval

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

	// IsLowInclusive indicates if the Low part of the interval is included, e. g. (3,5) --> the 3 is included as part of the interval
	IsLowInclusive() bool

	// IsHighInclusive indicates if the High part of the interval is included, e. g. (3,5) --> the 5 is included as part of the interval
	IsHighInclusive() bool

	// GetMinLow returns the minimal Low, either the one configured in the constructor, or the self-adjusted calculated if SelfAdjustMinLow=true
	GetMinLow() int

	// GetMaxHigh returns the maximal High, either the one configured in the constructor, or the self-adjusted calculated if SelfAdjustMaxHigh=true
	GetMaxHigh() int
}

// intervals implements Intervals interface
type intervals struct {
	Intervals         []*Interval
	GapsList          []*Interval
	OverlappedList    []*Interval
	MergeList         []*Interval
	MinLow            int
	MaxHigh           int
	Sorted            bool
	LowInclusive      bool
	HighInclusive     bool
	SelfAdjustMinLow  bool // set the minLow to the minimal Low value passed in Add or AddInterval methods
	SelfAdjustMaxHigh bool // set the maxHigh to the maximal High value passed in Add or AddInterval methods
}

// String implements Stringer.Interface Interval
func (itvls *intervals) String() string {
	return itvls.report()
}

// NewIntervalsDefault is a constructor that returns an instance of the Intervals interface with default values
func NewIntervalsDefault() Intervals {
	return NewIntervals(defaultMinLow, defaultMaxHigh, defaultLowInclusive, defaultHighInclusive, defaultSelfAdjustMinLow, defaultSelfAdjustMaxHigh)
}

// NewIntervals is a constructor that returns an instance of the Intervals interface
func NewIntervals(minLow, maxHigh int, lowInclusive, highInclusive, selfAdjustMinLow, selfAdjustMaxHigh bool) Intervals {
	return &intervals{
		MinLow:            minLow,
		MaxHigh:           maxHigh,
		Intervals:         []*Interval{},
		Sorted:            false,
		LowInclusive:      lowInclusive,
		HighInclusive:     highInclusive,
		SelfAdjustMinLow:  selfAdjustMinLow,
		SelfAdjustMaxHigh: selfAdjustMaxHigh,
	}
}

func (itvls *intervals) IsLowInclusive() bool {
	return itvls.LowInclusive
}

func (itvls *intervals) IsHighInclusive() bool {
	return itvls.HighInclusive
}

func (itvls *intervals) GetMinLow() int {
	return itvls.MinLow
}

func (itvls *intervals) GetMaxHigh() int {
	return itvls.MaxHigh
}
