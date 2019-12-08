package api

import "net/http"

var (
	errInvalidJSON = newAPIError(http.StatusBadRequest, false, "Invalid json")
	// FIXME: This shouldn't shouldn't exist
	errInvalidBody = newAPIError(http.StatusBadRequest, false, "Invalid body")
)

type apiError struct {
	status   int
	message  string
	err      error
	expected bool
}

func (apiErr apiError) Error() string {
	if apiErr.err != nil {
		return apiErr.message + ": " + apiErr.err.Error()
	}
	return apiErr.message
}

func (apiErr apiError) Unwrap() error {
	return apiErr.err
}

func unexpectedAPIError(err error) apiError {
	return apiError{
		status:   http.StatusInternalServerError,
		message:  "Unexpected error",
		err:      err,
		expected: false,
	}
}

func newSimpleAPIError(status int, expected bool, message string) apiError {
	return apiError{
		status:   status,
		message:  message,
		expected: expected,
	}
}

func newAPIError(status int, expected bool, message string) func(error) apiError {
	return func(err error) apiError {
		return apiError{
			status:   status,
			message:  message,
			expected: expected,
			err:      err,
		}
	}
}
