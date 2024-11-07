package apperrors

import "errors"

var Common = struct {
	Success             error
	BadRequest          error
	NotFound            error
	Unauthorized        error
	Forbidden           error
	AlreadyExisted      error
	EmailAlreadyExisted error
	InvalidID           error
	InvalidName         error
	InvalidEmail        error
	InvalidCode         error
	InvalidStatus       error
	InvalidFeedback     error
	InvalidQuote        error
}{
	Success:             errors.New("success"),
	BadRequest:          errors.New("bad_request"),
	NotFound:            errors.New("not_found"),
	Unauthorized:        errors.New("unauthorized"),
	Forbidden:           errors.New("forbidden"),
	AlreadyExisted:      errors.New("already_existed"),
	EmailAlreadyExisted: errors.New("email_already_existed"),
	InvalidID:           errors.New("invalid_id"),
	InvalidName:         errors.New("invalid_name"),
	InvalidEmail:        errors.New("invalid_email"),
	InvalidCode:         errors.New("invalid_code"),
	InvalidStatus:       errors.New("invalid_status"),
	InvalidFeedback:     errors.New("invalid_feedback"),
	InvalidQuote:        errors.New("invalid_quote"),
}
