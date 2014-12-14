package data

import (
	"errors"
	"net/http"
	"testing"
)

type MockResponseWriter struct {
	header http.Header
}

func NewMockResponseWriter() http.ResponseWriter {
	var header http.Header = make(map[string][]string)
	return &MockResponseWriter{header}
}

func (m *MockResponseWriter) Header() http.Header {
	return m.header
}

func (m *MockResponseWriter) Write(b []byte) (int, error) {
	return 500, errors.New("not implemented")
}

func (m *MockResponseWriter) WriteHeader(c int) {}

func TestSecureCookie(t *testing.T) {
	w := NewMockResponseWriter()
	d := new(Data)

	uuid := d.CreateUserId(w)

	if !uuid.Valid() {
		t.Error("invalid uuid")
	}

	if w.Header().Get("Set-Cookie") == "" {
		t.Error("nothing in set in the cookie")
	}

}
