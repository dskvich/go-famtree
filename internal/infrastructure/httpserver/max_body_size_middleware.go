package httpserver

import (
	"errors"
	"io"
	"net/http"
)

type middlewareBodyLimit struct {
	limit int64
	next  http.Handler
}

func bodyLimitMiddleware(limit int64) Middleware {
	return func(next http.Handler) http.Handler {
		return &middlewareBodyLimit{
			limit: limit,
			next:  next,
		}
	}
}

func (m *middlewareBodyLimit) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// ContentLength of -1 is unknown
	if r.ContentLength > m.limit {
		c := NewContext(w, r, m.s.log.Logger)
		c.WriteError(ErrWithRes(nil, http.StatusRequestEntityTooLarge, "", "Request too large"))
		return
	}
	r.Body = &maxBytesBodyReader{
		body: http.MaxBytesReader(w, r.Body, m.limit),
	}
	m.next.ServeHTTP(w, r)
}

func MaxBodySizeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		const (
			megabyte = 1 << 20
			maxSize  = 1 * megabyte
		)
		r2 := *r // shallow copy
		r2.Body = http.MaxBytesReader(w, r.Body, maxSize)
		next.ServeHTTP(w, &r2)
	})
}

type maxBytesBodyReader struct {
	body io.ReadCloser
}

func (r *maxBytesBodyReader) Read(p []byte) (int, error) {
	n, err := r.body.Read(p)
	if err != nil && !errors.Is(err, io.EOF) {
		var rerr *http.MaxBytesError
		if errors.As(err, &rerr) {
			return n, ErrWithRes(err, http.StatusRequestEntityTooLarge, "", "Request too large")
		}
		return n, ErrWithRes(err, http.StatusBadRequest, "", "Failed reading request body")
	}
	return n, err
}

// Close implements [io.Closer] on top of a [http.MaxBytesReader]
func (r *maxBytesBodyReader) Close() error {
	return r.body.Close()
}
