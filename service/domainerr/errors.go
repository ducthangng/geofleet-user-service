package domainerr

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	// --- Authentication Errors (401, 403) ---
	ErrInvalidCredentials = errors.New("invalid phone number or password")
	ErrTokenExpired       = errors.New("authentication token has expired")
	ErrTokenInvalid       = errors.New("authentication token is invalid")
	ErrMissingMetadata    = errors.New("missing metadata in request")
	ErrPermissionDenied   = errors.New("you do not have permission to perform this action")
	ErrAccountLocked      = errors.New("account is temporarily locked due to many failed attempts")
	ErrAccountInactive    = errors.New("account is deactivated, please contact support")

	// --- Validation Errors (400) ---
	ErrInvalidPhoneNumber = errors.New("invalid phone number format")
	ErrInvalidEmail       = errors.New("invalid email address format")
	ErrPasswordTooWeak    = errors.New("password must be at least 8 characters with numbers and symbols")
	ErrInvalidFullName    = errors.New("full name cannot be empty or contain special characters")
	ErrInvalidRole        = errors.New("invalid user role provided")
	ErrInvalidDateFormat  = errors.New("invalid date format for birthday (use YYYY-MM-DD)")

	// --- Resource Errors (404, 409) ---
	ErrUserNotFound          = errors.New("user not found in the system")
	ErrPhoneAlreadyExists    = errors.New("phone number is already registered")
	ErrEmailAlreadyExists    = errors.New("email address is already registered")
	ErrProfileAlreadyCreated = errors.New("user profile has already been initialized")

	// --- System & Internal Errors (500) ---
	ErrDatabaseOpFailed  = errors.New("unexpected error during database operation")
	ErrInternalServer    = errors.New("an internal server error occurred")
	ErrThirdPartyService = errors.New("failed to communicate with external identity provider")
	ErrDataConsistency   = errors.New("data inconsistency detected in user records")
)

/*
Gateway are not allowed to know private error in user service.
*/
func MapToGRPCError(err error) error {
	switch {
	case errors.Is(err, ErrUserNotFound):
		return status.Error(codes.NotFound, err.Error())

	case errors.Is(err, ErrPhoneAlreadyExists), errors.Is(err, ErrEmailAlreadyExists):
		return status.Error(codes.AlreadyExists, err.Error())

	case errors.Is(err, ErrInvalidCredentials), errors.Is(err, ErrTokenExpired):
		return status.Error(codes.Unauthenticated, err.Error())

	case errors.Is(err, ErrInvalidPhoneNumber), errors.Is(err, ErrPasswordTooWeak):
		return status.Error(codes.InvalidArgument, err.Error())

	case errors.Is(err, ErrPermissionDenied):
		return status.Error(codes.PermissionDenied, err.Error())

	default:
		return status.Error(codes.Internal, "Internal Server Error")
	}
}
