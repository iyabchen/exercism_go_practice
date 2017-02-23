package meetup

import "time"

const testVersion = 3

type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Last
	Teenth // 13(thirteen)-19(nineteen) in that month
)

const DayDur time.Duration = time.Hour * 24

// Given the weekschedul, weekday, month, year, eg. first monday in November 2013
// Calculate the date of meetups.
func Day(ws WeekSchedule, wd time.Weekday, month time.Month, year int) int {
	day := 1
	t := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	if t.Weekday() < wd {
		t = t.Add(DayDur * time.Duration(wd-t.Weekday()))
	} else if t.Weekday() > wd {
		t = t.AddDate(0, 0, int(7+wd-t.Weekday())) // another API of add days
	}
	day = t.Day()

	switch ws {
	case First:
		return day
	case Second:
		return day + 7
	case Third:
		return day + 14
	case Fourth:
		return day + 21
	case Teenth:
		if day+7 >= 13 {
			return day + 7
		} else {
			return day + 14
		}
	case Last:
		if t.Add(DayDur*28).Month() == month {
			return day + 28
		} else {
			return day + 21
		}
	default:
		return -1
	}
}
