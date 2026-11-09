package perigon

import "fmt"

// APIError is a non-2xx response from the Perigon API.
type APIError struct {
	StatusCode int
	Body       []byte
}

func (e *APIError) Error() string {
	return fmt.Sprintf("perigon api: status %d: %s", e.StatusCode, e.Body)
}

// Failed returns nil when status is 2xx. Otherwise it returns an APIError
// containing the status code and response body.
func Failed(status int, body []byte) error {
	if status >= 200 && status < 300 {
		return nil
	}
	return &APIError{StatusCode: status, Body: body}
}

// Decoded returns value when status is 2xx and value was decoded.
// A non-2xx status returns the error from Failed. A 2xx status with a nil
// value returns an error because the body was not decoded.
func Decoded[T any](status int, body []byte, value *T) (*T, error) {
	if err := Failed(status, body); err != nil {
		return nil, err
	}
	if value == nil {
		return nil, fmt.Errorf("perigon api: status %d: response body was not decoded", status)
	}
	return value, nil
}
