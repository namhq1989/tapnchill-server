package domain

type HabitStatus string

const (
	HabitStatusUnknown  HabitStatus = ""
	HabitStatusActive   HabitStatus = "active"
	HabitStatusInactive HabitStatus = "inactive"
)

func (s HabitStatus) String() string {
	return string(s)
}

func (s HabitStatus) IsValid() bool {
	return s != HabitStatusUnknown
}

func ToHabitStatus(value string) HabitStatus {
	switch value {
	case HabitStatusActive.String():
		return HabitStatusActive
	case HabitStatusInactive.String():
		return HabitStatusInactive
	default:
		return HabitStatusUnknown
	}
}
