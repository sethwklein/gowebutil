package webutil

import (
	"fmt"
	"net/http"
)

// HTTPError represents an undesired HTTP response. It is usually used by
// functions that want to return an error on non-200 responses.
type HTTPError struct {
	Status     string
	Code int
	URL        string
}

// NewHTTPError constructs an HTTPError from the provided http.Response.
func NewHTTPError(resp *http.Response) *HTTPError {
	return &HTTPError{
		Status:     resp.Status,
		Code: resp.StatusCode,
		URL:        resp.Request.URL.String(),
	}
}

// Error returns a human readable version.
func (e *HTTPError) Error() string {
	return fmt.Sprintf("server returned %v for %v", e.Status, e.URL)
}
