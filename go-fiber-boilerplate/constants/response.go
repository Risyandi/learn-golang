package constants

const (
	ResponseSuccessGet    = "success_get"
	ResponseSuccessCreate = "success_create"
	ResponseSuccessUpdate = "success_update"
	ResponseSuccessDelete = "success_delete"

	ResponseErrorMessage  = "error_response"
	ResponseErrorNotFound = "error_not_found"

	// Authorization
	ErrInvalidOrEmptyToken       = "error_invalid_token"
	ErrorHttpInvalidServiceToken = "error_invalid_service_token"
	ErrTokenIsExpired            = "error_token_expired"
	ErrInvalidSignature          = "error_invalid_signature"
	ErrAccountSuspended          = "error_account_suspended"

	ValidationErrors          = "error_validation"
	ValidationErrorRequired   = "error_validation_required"
	ValidationErrorEmail      = "error_validation_email"
	ValidationErrorUnique     = "error_validation_unique"
	ValidationErrorNumber     = "error_validation_numeric"
	ValidationErrorGte        = "error_validation_gte"
	ValidationErrorGt         = "error_validation_gt"
	ValidationErrorLte        = "error_validation_lte"
	ValidationErrorLt         = "error_validation_lt"
	ValidationErrorMin        = "error_validation_min_length"
	ValidationErrorMax        = "error_validation_max_length"
	ValidationErrorStartswith = "error_validation_startwith"
	ValidationErrorLen        = "error_validation_length"
	ValidationErrorOneof      = "error_validation_oneof"
	ValidationErrorUUID       = "error_validation_uuid"
	ValidationErrorCleanEmoji = "error_validation_clean_emoji"

	ErrLimitReached = "error_limit_reached"
)
