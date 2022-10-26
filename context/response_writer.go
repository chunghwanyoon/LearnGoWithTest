package context

import "net/http"

type SpyResponseWriter struct {
	written bool
}

/**
SpyResponseWriter implements http.ResponseWriter interface
*/
func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, nil
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}
