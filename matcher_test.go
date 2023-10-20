package matcher

import (
	"errors"
	"log"
	"net/http"
	"testing"
)

func TestMatcher_AddPattern(t *testing.T) {

	Notfund := errors.New("error")
	UnAuthorize := errors.New("Un UnAuthorize")

	m := NewMatcher()
	m.AddPattern(Notfund, http.StatusNotFound)
	m.AddPattern(UnAuthorize, http.StatusUnauthorized)

	i, err := m.Match(Notfund)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(i)

	i, err = m.Match(UnAuthorize)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(i)

}

func AddPattern() {
	Notfund := errors.New("error")
	UnAuthorize := errors.New("unauthorize")

	m := NewMatcher()
	m.AddPattern(Notfund, http.StatusNotFound)
	m.AddPattern(UnAuthorize, http.StatusUnauthorized)

	i, err := m.Match(Notfund)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(i)
}
