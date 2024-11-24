package apperrors

import "errors"

var User = struct {
	InvalidUserID        error
	InvalidClientID      error
	UserNotFound         error
	ResourceLimitReached error
}{
	InvalidUserID:        errors.New("user_invalid_id"),
	InvalidClientID:      errors.New("user_invalid_client_id"),
	UserNotFound:         errors.New("user_not_found"),
	ResourceLimitReached: errors.New("user_resource_limit_reached"),
}
