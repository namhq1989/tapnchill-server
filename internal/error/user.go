package apperrors

import "errors"

var User = struct {
	InvalidUserID        error
	InvalidClientID      error
	UserNotFound         error
	FreePlanLimitReached error
	ProPlanLimitReached  error
}{
	InvalidUserID:        errors.New("user_invalid_id"),
	InvalidClientID:      errors.New("user_invalid_client_id"),
	UserNotFound:         errors.New("user_not_found"),
	FreePlanLimitReached: errors.New("user_free_plan_limit_reached"),
	ProPlanLimitReached:  errors.New("user_pro_plan_limit_reached"),
}
