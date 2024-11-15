package apperrors

import "errors"

var Habit = struct {
	InvalidID error
	NotFound  error
}{
	InvalidID: errors.New("habit_invalid_id"),
	NotFound:  errors.New("habit_not_found"),
}
