package dates

import "time"

func New(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func Day(in time.Time) time.Time {
	return in.UTC().Truncate(24 * time.Hour)
}

func EndOfMonth(d time.Time) time.Time {
	return Day(NextMonth(d).Add(-1 * time.Hour))
}

func NextMonth(d time.Time) time.Time {
	if d.Month() == 12 {
		return New(d.Year()+1, 1, 1)
	} else {
		return New(d.Year(), int(d.Month()+1), 1)
	}
}

func Range(start, end time.Time) (list Times) {
	c := Day(start)
	end = Day(end)
	for !c.After(end) {
		list = append(list, c)
		c = c.Add(24 * time.Hour)
	}
	return list
}

type Times []time.Time

func (list Times) Len() int {
	return len(list)
}

func (list Times) Swap(a, b int) {
	list[a], list[b] = list[b], list[a]
}

func (list Times) Less(a, b int) bool {
	return list[a].Before(list[b])
}