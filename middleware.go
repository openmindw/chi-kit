package httpd

import (
	"net/http"
	"time"

	"github.com/opentdp/go-helper/logman"
)

// LoggerMiddleware is a custom logger middleware for Chi
func LoggerMiddleware(debug bool) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// 创建一个包装的ResponseWriter来捕获状态码
			ww := &responseWriter{
				ResponseWriter: w,
				statusCode:     http.StatusOK,
			}

			// 处理请求
			next.ServeHTTP(ww, r)

			// 记录日志
			duration := time.Since(start)
			
			logFields := []interface{}{
				"method", r.Method,
				"path", r.URL.Path,
				"status", ww.statusCode,
				"duration", duration.String(),
				"ip", r.RemoteAddr,
			}

			if debug {
				logFields = append(logFields, "query", r.URL.RawQuery)
			}

			if ww.statusCode >= 500 {
				logman.Error("http request", logFields...)
			} else if ww.statusCode >= 400 {
				logman.Warn("http request", logFields...)
			} else {
				logman.Info("http request", logFields...)
			}
		})
	}
}

// responseWriter wraps http.ResponseWriter to capture status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
