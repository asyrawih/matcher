package matcher

import (
	"errors"
	"net/http"
	"sync"
)

var (
	// ErrorPatternNotFound 🚀
	ErrorPatternNotFound = errors.New("pattern not found")

	// ErrInvalidStatusCode 🚀
	ErrInvalidStatusCode = errors.New("invalid status code")
)

// Matcher struct 🚀
type Matcher struct {
	pattern map[error]int
	mutex   *sync.Mutex
}

// NewMatcher function 🚀
func NewMatcher() *Matcher {
	m := new(sync.Mutex)
	return &Matcher{
		pattern: make(map[error]int),
		mutex:   m,
	}
}

// AddPattern method 🚀
func (m *Matcher) AddPattern(err error, statusCode int) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.pattern[err] = statusCode
}

func (m *Matcher) Match(err error) (int, error) {
	statusCode, ok := m.pattern[err]
	if !ok {
		return 0, ErrorPatternNotFound
	}

	if statusCode < http.StatusOK || statusCode > http.StatusNetworkAuthenticationRequired {
		return 0, ErrInvalidStatusCode
	}
	return statusCode, nil
}

// GetPattern method 🚀
func (m *Matcher) GetPattern() map[error]int {
	return m.pattern
}
