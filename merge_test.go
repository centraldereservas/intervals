package interval_test

import (
	"testing"

	"bitbucket.org/differenttravel/interval"
)

func TestMerge(t *testing.T) {
	// tests for low/high inclusive
	demo001 := buildIntervalsDemo001()
	demo002 := buildIntervalsDemo002()
	demo003 := buildIntervalsDemo003()
	demo004 := buildIntervalsDemo004()
	demo005 := buildIntervalsDemo005()
	demo006 := buildIntervalsDemo006()
	demo007 := buildIntervalsDemo007()
	demo008 := buildIntervalsDemo008()
	demo009 := buildIntervalsDemo009()
	demo010 := buildIntervalsDemo010()
	demo011 := buildIntervalsDemo011()
	demo012 := buildIntervalsDemo012()
	demo013 := buildIntervalsDemo013()
	demo014 := buildIntervalsDemo014()
	demo015 := buildIntervalsDemo015()

	// tests for low/high exclusive
	demo101 := buildIntervalsDemo101()
	demo102 := buildIntervalsDemo102()
	demo103 := buildIntervalsDemo103()
	demo104 := buildIntervalsDemo104()
	demo105 := buildIntervalsDemo105()
	demo106 := buildIntervalsDemo106()
	demo107 := buildIntervalsDemo107()
	demo108 := buildIntervalsDemo108()
	demo109 := buildIntervalsDemo109()
	demo110 := buildIntervalsDemo110()
	demo111 := buildIntervalsDemo111()
	demo112 := buildIntervalsDemo112()

	// tests for low inclusive and high exclusive
	demo201 := buildIntervalsDemo201()
	demo202 := buildIntervalsDemo202()
	demo203 := buildIntervalsDemo203()
	demo204 := buildIntervalsDemo204()
	demo205 := buildIntervalsDemo205()
	demo206 := buildIntervalsDemo206()
	demo207 := buildIntervalsDemo207()
	demo208 := buildIntervalsDemo208()
	demo209 := buildIntervalsDemo209()
	demo210 := buildIntervalsDemo210()
	demo211 := buildIntervalsDemo211()
	demo212 := buildIntervalsDemo212()

	// tests for low exclusive and high inclusive
	demo301 := buildIntervalsDemo301()
	demo302 := buildIntervalsDemo302()
	demo303 := buildIntervalsDemo303()
	demo304 := buildIntervalsDemo304()
	demo305 := buildIntervalsDemo305()
	demo306 := buildIntervalsDemo306()
	demo307 := buildIntervalsDemo307()
	demo308 := buildIntervalsDemo308()
	demo309 := buildIntervalsDemo309()
	demo310 := buildIntervalsDemo310()
	demo311 := buildIntervalsDemo311()
	demo312 := buildIntervalsDemo312()
	tt := []struct {
		name           string
		intvls         interval.Intervals
		expectedMerges []interval.Interval
	}{
		{name: "demo001", intvls: demo001.Intervals, expectedMerges: demo001.ExpectedMerges},
		{name: "demo002", intvls: demo002.Intervals, expectedMerges: demo002.ExpectedMerges},
		{name: "demo003", intvls: demo003.Intervals, expectedMerges: demo003.ExpectedMerges},
		{name: "demo004", intvls: demo004.Intervals, expectedMerges: demo004.ExpectedMerges},
		{name: "demo005", intvls: demo005.Intervals, expectedMerges: demo005.ExpectedMerges},
		{name: "demo006", intvls: demo006.Intervals, expectedMerges: demo006.ExpectedMerges},
		{name: "demo007", intvls: demo007.Intervals, expectedMerges: demo007.ExpectedMerges},
		{name: "demo008", intvls: demo008.Intervals, expectedMerges: demo008.ExpectedMerges},
		{name: "demo009", intvls: demo009.Intervals, expectedMerges: demo009.ExpectedMerges},
		{name: "demo010", intvls: demo010.Intervals, expectedMerges: demo010.ExpectedMerges},
		{name: "demo011", intvls: demo011.Intervals, expectedMerges: demo011.ExpectedMerges},
		{name: "demo012", intvls: demo012.Intervals, expectedMerges: demo012.ExpectedMerges},
		{name: "demo013", intvls: demo013.Intervals, expectedMerges: demo013.ExpectedMerges},
		{name: "demo014", intvls: demo014.Intervals, expectedMerges: demo014.ExpectedMerges},
		{name: "demo015", intvls: demo015.Intervals, expectedMerges: demo015.ExpectedMerges},

		{name: "demo101", intvls: demo101.Intervals, expectedMerges: demo101.ExpectedMerges},
		{name: "demo102", intvls: demo102.Intervals, expectedMerges: demo102.ExpectedMerges},
		{name: "demo103", intvls: demo103.Intervals, expectedMerges: demo103.ExpectedMerges},
		{name: "demo104", intvls: demo104.Intervals, expectedMerges: demo104.ExpectedMerges},
		{name: "demo105", intvls: demo105.Intervals, expectedMerges: demo105.ExpectedMerges},
		{name: "demo106", intvls: demo106.Intervals, expectedMerges: demo106.ExpectedMerges},
		{name: "demo107", intvls: demo107.Intervals, expectedMerges: demo107.ExpectedMerges},
		{name: "demo108", intvls: demo108.Intervals, expectedMerges: demo108.ExpectedMerges},
		{name: "demo109", intvls: demo109.Intervals, expectedMerges: demo109.ExpectedMerges},
		{name: "demo110", intvls: demo110.Intervals, expectedMerges: demo110.ExpectedMerges},
		{name: "demo111", intvls: demo111.Intervals, expectedMerges: demo111.ExpectedMerges},
		{name: "demo112", intvls: demo112.Intervals, expectedMerges: demo112.ExpectedMerges},

		{name: "demo201", intvls: demo201.Intervals, expectedMerges: demo201.ExpectedMerges},
		{name: "demo202", intvls: demo202.Intervals, expectedMerges: demo202.ExpectedMerges},
		{name: "demo203", intvls: demo203.Intervals, expectedMerges: demo203.ExpectedMerges},
		{name: "demo204", intvls: demo204.Intervals, expectedMerges: demo204.ExpectedMerges},
		{name: "demo205", intvls: demo205.Intervals, expectedMerges: demo205.ExpectedMerges},
		{name: "demo206", intvls: demo206.Intervals, expectedMerges: demo206.ExpectedMerges},
		{name: "demo207", intvls: demo207.Intervals, expectedMerges: demo207.ExpectedMerges},
		{name: "demo208", intvls: demo208.Intervals, expectedMerges: demo208.ExpectedMerges},
		{name: "demo209", intvls: demo209.Intervals, expectedMerges: demo209.ExpectedMerges},
		{name: "demo210", intvls: demo210.Intervals, expectedMerges: demo210.ExpectedMerges},
		{name: "demo211", intvls: demo211.Intervals, expectedMerges: demo211.ExpectedMerges},
		{name: "demo212", intvls: demo212.Intervals, expectedMerges: demo212.ExpectedMerges},

		{name: "demo301", intvls: demo301.Intervals, expectedMerges: demo301.ExpectedMerges},
		{name: "demo302", intvls: demo302.Intervals, expectedMerges: demo302.ExpectedMerges},
		{name: "demo303", intvls: demo303.Intervals, expectedMerges: demo303.ExpectedMerges},
		{name: "demo304", intvls: demo304.Intervals, expectedMerges: demo304.ExpectedMerges},
		{name: "demo305", intvls: demo305.Intervals, expectedMerges: demo305.ExpectedMerges},
		{name: "demo306", intvls: demo306.Intervals, expectedMerges: demo306.ExpectedMerges},
		{name: "demo307", intvls: demo307.Intervals, expectedMerges: demo307.ExpectedMerges},
		{name: "demo308", intvls: demo308.Intervals, expectedMerges: demo308.ExpectedMerges},
		{name: "demo309", intvls: demo309.Intervals, expectedMerges: demo309.ExpectedMerges},
		{name: "demo310", intvls: demo310.Intervals, expectedMerges: demo310.ExpectedMerges},
		{name: "demo311", intvls: demo311.Intervals, expectedMerges: demo311.ExpectedMerges},
		{name: "demo312", intvls: demo312.Intervals, expectedMerges: demo312.ExpectedMerges},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// if tc.name == "demo105" {
			// 	t.Log("break here")
			// }
			merges := tc.intvls.Merge()
			unexpectedLength := false
			if len(merges) != len(tc.expectedMerges) {
				t.Errorf("unexpected merges length: expected %d; got %d", len(tc.expectedMerges), len(merges))
				unexpectedLength = true
			}
			if !unexpectedLength {
				for i, merge := range merges {
					if merge.Low != tc.expectedMerges[i].Low || merge.High != tc.expectedMerges[i].High {
						t.Errorf("unexpected merge[%d]: expected %v; got %v", i, tc.expectedMerges[i], merge)
					}
				}
			}
		})
	}
}
