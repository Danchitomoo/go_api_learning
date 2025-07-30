package middlewares

import (
	"log"
	"net/http"
)

/*
Response Writer is an interface and the required methods are below.
- Header
- Write
- WriteHeader
Header method and Write method are delegated to the http.ResponseWriter.
WriteHeader method is overrided to store the status code in the resLogging Writer struct.
*/

type resLoggingWriter struct {
	http.ResponseWriter
	code int
}

func NewResLoggingWriter(w http.ResponseWriter) *resLoggingWriter {
	return &resLoggingWriter{ResponseWriter: w, code: http.StatusOK}
}

// 　override WriteHeader method
func (rlw *resLoggingWriter) WriteHeader(code int) {
	rlw.code = code
	rlw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		traceID := newTraceID()

		log.Printf("[%d]%s %s \n", traceID, req.RequestURI, req.Method)

		ctx := setTraceID(req.Context(), traceID)
		req = req.WithContext(ctx)
		rlw := NewResLoggingWriter(w)
		next.ServeHTTP(rlw, req)

		log.Printf("[%d]res: [%d]", traceID, rlw.code)
	})
}
