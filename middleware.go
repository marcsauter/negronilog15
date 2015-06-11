package negronilog15

import (
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"gopkg.in/inconshreveable/log15.v2"
)

// Middleware is a middleware handler that logs the request as it goes in and the response as it goes out.
type Middleware struct {
	// Logger is the log.Logger instance used to log messages with the Logger middleware
	Logger log15.Logger
}

// NewMiddleware returns a new *Middleware
func NewMiddleware() *Middleware {
	l := log15.New()
	h := log15.LvlFilterHandler(log15.LvlInfo, log15.StreamHandler(os.Stdout, log15.TerminalFormat()))
	l.SetHandler(h)
	return &Middleware{Logger: l}
}

func (l *Middleware) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	start := time.Now()
	l.Logger.Info("started handling request", "request", r.RequestURI, "method", r.Method, "remote", r.RemoteAddr)
	next(rw, r)
	latency := time.Since(start)
	res := rw.(negroni.ResponseWriter)
	l.Logger.Info("completed handling request", "status", res.Status(), "statustext", http.StatusText(res.Status()), "took", latency)
}
