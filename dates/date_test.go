package dates

import (
	"testing"
	"time"
)

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
