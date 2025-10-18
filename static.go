package chikit

import (
	"embed"
	"io/fs"
	"net/http"
	"path"
	"strings"
)

func Static(prefix, root string) {

	if router == nil {
		Engine(false)
	}

	hfs := http.Dir(root)
	router.Use(StaticServe(prefix, hfs))

}

func StaticIndex(prefix, root string) {

	if router == nil {
		Engine(false)
	}

	hfs := http.Dir(root)
	router.Use(StaticServeWithIndex(prefix, hfs))

}

func StaticEmbed(prefix, sub string, efs *embed.FS) {

	if router == nil {
		Engine(false)
	}

	var hfs http.FileSystem

	if sub == "" {
		hfs = http.FS(efs)
	} else {
		subFS, _ := fs.Sub(efs, sub)
		hfs = http.FS(subFS)
	}

	router.Use(StaticServe(prefix, hfs))

}

func StaticServe(prefix string, hfs http.FileSystem) func(next http.Handler) http.Handler {

	fileServer := http.FileServer(hfs)
	if prefix != "" {
		fileServer = http.StripPrefix(prefix, fileServer)
	}

	isExists := func(p, s string) bool {
		if p := strings.TrimPrefix(s, p); len(p) < len(s) {
			if f, err := hfs.Open(path.Join("/", p)); err == nil {
				defer f.Close()
				return true
			}
		}
		return false
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if isExists(prefix, r.URL.Path) {
				fileServer.ServeHTTP(w, r)
				return
			}
			next.ServeHTTP(w, r)
		})
	}

}

func StaticServeWithIndex(prefix string, hfs http.FileSystem) func(next http.Handler) http.Handler {

	fileServer := http.FileServer(hfs)
	if prefix != "" {
		fileServer = http.StripPrefix(prefix, fileServer)
	}

	isExists := func(p, s string) bool {
		if p := strings.TrimPrefix(s, p); len(p) < len(s) {
			if f, err := hfs.Open(path.Join("/", p)); err == nil {
				defer f.Close()
				return true
			}
		}
		return false
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if isExists(prefix, r.URL.Path) {
				fileServer.ServeHTTP(w, r)
				return
			}
			next.ServeHTTP(w, r)
		})
	}

}
