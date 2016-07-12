package dates

import (
	"testing"
	"time"
)

func TestBetween(t *testing.T) {
	ref := time.Date(2006, 2, 3, 15, 4, 5, 0, time.UTC)

	tests := []struct {
		Start time.Time
		End   time.Time
		Want  bool
	}{
		{New(2006, 2, 3), New(2006, 2, 4), true},
		{ref, ref, true},
		{ref.Add(-1 * time.Second), ref, true},
		{ref.Add(-1 * time.Second), ref.Add(1 * time.Second), true},
		{ref.Add(1 * time.Second), ref, false},
		{ref, ref.Add(-1 * time.Second), false},
	}
	for i, tc := range tests {
		has := Between(ref, tc.Start, tc.End)
		if has != tc.Want {
			t.Errorf("%d: expected Between for %s and %s to be %t, was %t", i+1, tc.Start, tc.End, tc.Want, has)
		}
	}
}

func TestPrevMonth(t *testing.T) {
	tests := []struct{ Input, Want time.Time }{
		{New(2006, 1, 1), New(2005, 12, 1)},
		{New(2006, 12, 4), New(2006, 11, 1)},
		{New(2006, 3, 4), New(2006, 2, 1)},
		{New(2016, 2, 29), New(2016, 1, 1)}, // leap year
		{New(2015, 2, 28), New(2015, 1, 1)}, // non leap year
	}
	for i, tc := range tests {
		has := PrevMonth(tc.Input)
		if has != tc.Want {
			t.Errorf("%d: input=%s want=%s has=%s", i+1, tc.Input.Format("2006-01-02"), tc.Want.Format("2006-01-02"), has.Format("2006-01-02"))
		}
	}
}

func TestNextMonth(t *testing.T) {
	tests := []struct{ Input, Want time.Time }{
		{New(2006, 12, 4), New(2007, 1, 1)},
		{New(2006, 3, 4), New(2006, 4, 1)},
		{New(2016, 2, 29), New(2016, 3, 1)}, // leap year
		{New(2015, 2, 28), New(2015, 3, 1)}, // non leap year
	}
	for i, tc := range tests {
		has := NextMonth(tc.Input)
		if has != tc.Want {
			t.Errorf("%d: input=%s want=%s has=%s", i+1, tc.Input.Format("2006-01-02"), tc.Want.Format("2006-01-02"), has.Format("2006-01-02"))
		}
	}
}

func TestEndOfMonth(t *testing.T) {
	tests := []struct{ Input, Want time.Time }{
		{New(2006, 12, 4), New(2006, 12, 31)},
		{New(2006, 3, 4), New(2006, 3, 31)},
		{New(2006, 4, 4), New(2006, 4, 30)},
		{New(2016, 2, 4), New(2016, 2, 29)}, // leap year
		{New(2015, 2, 4), New(2015, 2, 28)}, // non leap year
	}
	for i, tc := range tests {
		has := EndOfMonth(tc.Input)
		if has != tc.Want {
			t.Errorf("%d: input=%s want=%s has=%s", i+1, tc.Input.Format("2006-01-02"), tc.Want.Format("2006-01-02"), has.Format("2006-01-02"))
		}
	}
}
