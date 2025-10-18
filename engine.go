package httpd

import (
	"mime"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/opentdp/go-helper/logman"
)

var router *chi.Mux

func Engine(debug bool) *chi.Mux {

	// Chi没有内置的debug/release模式，我们通过日志级别来控制
	if debug {
		logman.Info("chi engine starting in debug mode")
	} else {
		logman.Info("chi engine starting in release mode")
	}

	mime.AddExtensionType(".css", "text/css; charset=utf-8")
	mime.AddExtensionType(".js", "text/javascript; charset=utf-8")

	router = chi.NewRouter()

	// 添加Chi的默认中间件（类似gin.Default()）
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(LoggerMiddleware(debug)) // 自定义日志中间件
	router.Use(middleware.Recoverer)

	return router

}

func Group(pattern string) chi.Router {

	if router == nil {
		Engine(false)
	}

	return router.Route(pattern, func(r chi.Router) {})

}

func Route(pattern string, fn func(r chi.Router)) chi.Router {

	if router == nil {
		Engine(false)
	}

	return router.Route(pattern, fn)

}

func Use(middlewares ...func(next http.Handler) http.Handler) {

	if router == nil {
		Engine(false)
	}

	router.Use(middlewares...)

}

func Get(pattern string, handlerFn http.HandlerFunc) {

	if router == nil {
		Engine(false)
	}

	router.Get(pattern, handlerFn)

}

func Post(pattern string, handlerFn http.HandlerFunc) {

	if router == nil {
		Engine(false)
	}

	router.Post(pattern, handlerFn)

}

func Put(pattern string, handlerFn http.HandlerFunc) {

	if router == nil {
		Engine(false)
	}

	router.Put(pattern, handlerFn)

}

func Delete(pattern string, handlerFn http.HandlerFunc) {

	if router == nil {
		Engine(false)
	}

	router.Delete(pattern, handlerFn)

}

func Patch(pattern string, handlerFn http.HandlerFunc) {

	if router == nil {
		Engine(false)
	}

	router.Patch(pattern, handlerFn)

}
