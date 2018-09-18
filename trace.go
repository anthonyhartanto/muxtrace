package muxtrace

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/opentracing/opentracing-go"
)

// TraceAndServe will apply tracing to the given http.Handler using the passed tracer under the given service and resource.
func TraceAndServe(h http.Handler, w http.ResponseWriter, r *http.Request, service, resource string, spanopts ...opentracing.StartSpanOption) {
	opts := append([]opentracing.StartSpanOption{
		opentracing.Tag{
			Key:   "http.method",
			Value: r.Method,
		},
		opentracing.Tag{
			Key:   "http.url",
			Value: r.URL.Path,
		},
		opentracing.Tag{
			Key:   "resource.name",
			Value: resource,
		},
	}, spanopts...)

	span, ctx := opentracing.StartSpanFromContext(r.Context(), service, opts...)
	defer span.Finish()

	w = wrapResponseWriter(w, span)

	h.ServeHTTP(w, r.WithContext(ctx))
}

// responseWriter is a small wrapper around an http response writer that will
// intercept and store the status of a request.
type responseWriter struct {
	http.ResponseWriter
	span   opentracing.Span
	status int
}

func newResponseWriter(w http.ResponseWriter, span opentracing.Span) *responseWriter {
	return &responseWriter{w, span, 0}
}

// Write writes the data to the connection as part of an HTTP reply.
// We explicitly call WriteHeader with the 200 status code
// in order to get it reported into the span.
func (w *responseWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.WriteHeader(http.StatusOK)
	}
	return w.ResponseWriter.Write(b)
}

// WriteHeader sends an HTTP response header with status code.
// It also sets the status code to the span.
func (w *responseWriter) WriteHeader(status int) {
	w.ResponseWriter.WriteHeader(status)
	w.status = status
	w.span.SetTag("http.status_code", strconv.Itoa(status))
	if status >= 500 && status < 600 {
		w.span.SetTag("error", fmt.Errorf("%d: %s", status, http.StatusText(status)))
	}
}
