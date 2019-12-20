package middleware

import (
	"fmt"
	"github.com/joshia/automated-api-test-service/testapp/lib/logger"
	"net/http"
)

type ResponseLogger struct {
	rw				http.ResponseWriter
	body			string
	status			int
}

func (w *ResponseLogger) Header() http.Header {
	return w.rw.Header()
}

func (w *ResponseLogger) WriteHeader(status int) {
	w.status = status
	w.rw.WriteHeader(status)
}

func (w *ResponseLogger) Write(b []byte) (int, error) {
	w.body = string(b)
	if w.status != 200 {
		b = RewriteErrorResponse(b)
	}
	return w.rw.Write(b)
}



func NewHTTPHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Logging...")
		nw := &ResponseLogger{rw:w}
		r = InjectRequestId(r)
		h.ServeHTTP(nw, r) // call original
		log := fmt.Sprintf("URL:%+v Method:%+v RequestId:%+v Body:%+v RequestHeaders:%+v ResHttpCode:%+v Response:%s",
			r.RequestURI, r.Method, r.Context().Value("requestId"), nw.body, w.Header(), nw.status, nw.body)
		if nw.status != 200 {
			logger.Error(log)
			return
		}
		logger.Info(log)
	})
}
