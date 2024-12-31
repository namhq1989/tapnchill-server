package manipulation

import (
	"fmt"
	"time"
)

func resolveTimezone(tz string) *time.Location {
	if tz == "" {
		return time.UTC
	}

	// Handle fixed offsets like "+07:00"
	if len(tz) == 6 && (tz[0] == '+' || tz[0] == '-') {
		var sign int
		if tz[0] == '+' {
			sign = 1
		} else {
			sign = -1
		}

		var hours, minutes int
		_, err := fmt.Sscanf(tz[1:], "%02d:%02d", &hours, &minutes)
		if err != nil {
			fmt.Printf("invalid offset format '%s', defaulting to UTC: %v\n", tz, err)
			return time.UTC
		}

		offsetSeconds := sign * (hours*3600 + minutes*60)
		return time.FixedZone(tz, offsetSeconds)
	}

	// Handle named time zones
	loc, err := time.LoadLocation(tz)
	if err != nil {
		fmt.Printf("invalid timezone '%s', defaulting to UTC: %v\n", tz, err)
		return time.UTC
	}
	return loc
}

// func GetTimezoneIdentifier(t time.Time) string {
// 	location := t.Location()
// 	return location.String()
// }

func GetUTCOffset(t time.Time) string {
	_, offset := t.Zone()
	return fmt.Sprintf("%+03d:%02d", offset/3600, (offset%3600)/60)
}

func Now(tz string) time.Time {
	loc := resolveTimezone(tz)
	return time.Now().In(loc)
}

func NowUTC() time.Time {
	return time.Now().UTC()
}

func IsToday(t time.Time, tz string) bool {
	now := Now(tz)
	tInLoc := t.In(resolveTimezone(tz))

	return tInLoc.Year() == now.Year() && tInLoc.Month() == now.Month() && tInLoc.Day() == now.Day()
}

func IsYesterday(t time.Time, tz string) bool {
	now := Now(tz)
	yesterday := now.AddDate(0, 0, -1)
	tInLoc := t.In(resolveTimezone(tz))

	return tInLoc.Year() == yesterday.Year() && tInLoc.Month() == yesterday.Month() && tInLoc.Day() == yesterday.Day()
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

func PreviousDay(date time.Time, tz string) time.Time {
	loc := resolveTimezone(tz)
	yesterday := date.AddDate(0, 0, -1)
	return yesterday.In(loc)
}

func IsSameDay(date1 time.Time, date2 time.Time, tz string) bool {
	loc := resolveTimezone(tz)
	date1InLoc := date1.In(loc)
	date2InLoc := date2.In(loc)
	return date1InLoc.Year() == date2InLoc.Year() && date1InLoc.Month() == date2InLoc.Month() && date1InLoc.Day() == date2InLoc.Day()
}
