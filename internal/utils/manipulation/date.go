package manipulation

import (
	"fmt"
	"time"
)

func getLocation(tz string) *time.Location {
	if tz == "" {
		return time.UTC
	}

	loc, err := time.LoadLocation(tz)
	if err != nil {
		fmt.Printf("invalid timezone '%s', defaulting to UTC: %v\n", tz, err)
		return time.UTC
	}
	return loc
}

func Now(tz string) time.Time {
	loc := getLocation(tz)
	return time.Now().In(loc)
}

func NowUTC() time.Time {
	return time.Now().UTC()
}

func IsToday(t time.Time) bool {
	nowUTC := NowUTC()
	tUTC := t.UTC()

	return tUTC.Year() == nowUTC.Year() && tUTC.Month() == nowUTC.Month() && tUTC.Day() == nowUTC.Day()
}

func IsYesterday(t time.Time) bool {
	nowUTC := NowUTC()
	yesterdayUTC := nowUTC.AddDate(0, 0, -1)
	tUTC := t.UTC()

	return tUTC.Year() == yesterdayUTC.Year() && tUTC.Month() == yesterdayUTC.Month() && tUTC.Day() == yesterdayUTC.Day()
}

func StartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func EndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, int(time.Nanosecond*999999999), t.Location())
}

func GetStartOfDayWithClientDate(date string) (*time.Time, error) {
	clientTime, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return nil, fmt.Errorf("invalid client date: %w", err)
	}

	s := StartOfDay(clientTime)
	return &s, nil
}

func GetEndOfDayWithClientDate(date string) (*time.Time, error) {
	clientTime, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return nil, fmt.Errorf("invalid client date: %w", err)
	}

	s := EndOfDay(clientTime)
	return &s, nil
}
