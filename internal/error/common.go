package apperrors

import "errors"

var Common = struct {
	Success               error
	BadRequest            error
	NotFound              error
	Unauthorized          error
	Forbidden             error
	SomethingWentWrong    error
	AlreadyExisted        error
	EmailAlreadyExisted   error
	InvalidID             error
	InvalidName           error
	InvalidDescription    error
	InvalidType           error
	InvalidContent        error
	InvalidEmail          error
	InvalidCode           error
	InvalidDate           error
	InvalidGoal           error
	InvalidDaysOfWeek     error
	InvalidIcon           error
	InvalidStatus         error
	InvalidFeedback       error
	InvalidQuote          error
	InvalidNote           error
	InvalidQRCode         error
	InvalidQRCodeSettings error
}{
	Success:               errors.New("success"),
	BadRequest:            errors.New("bad_request"),
	NotFound:              errors.New("not_found"),
	Unauthorized:          errors.New("unauthorized"),
	Forbidden:             errors.New("forbidden"),
	SomethingWentWrong:    errors.New("something_went_wrong"),
	AlreadyExisted:        errors.New("already_existed"),
	EmailAlreadyExisted:   errors.New("email_already_existed"),
	InvalidID:             errors.New("invalid_id"),
	InvalidName:           errors.New("invalid_name"),
	InvalidDescription:    errors.New("invalid_description"),
	InvalidType:           errors.New("invalid_type"),
	InvalidContent:        errors.New("invalid_content"),
	InvalidEmail:          errors.New("invalid_email"),
	InvalidCode:           errors.New("invalid_code"),
	InvalidDate:           errors.New("invalid_date"),
	InvalidGoal:           errors.New("invalid_goal"),
	InvalidDaysOfWeek:     errors.New("invalid_days_of_week"),
	InvalidIcon:           errors.New("invalid_icon"),
	InvalidStatus:         errors.New("invalid_status"),
	InvalidFeedback:       errors.New("invalid_feedback"),
	InvalidQuote:          errors.New("invalid_quote"),
	InvalidNote:           errors.New("invalid_note"),
	InvalidQRCode:         errors.New("invalid_qr_code"),
	InvalidQRCodeSettings: errors.New("invalid_qr_code_settings"),
}
