package httpd

/*
使用示例 (Example Usage)

package main

import (
	"net/http"
	"github.com/opentdp/go-helper/httpd"
)

func main() {
	// 1. 初始化引擎（调试模式）
	httpd.Engine(true)

	// 2. 注册全局中间件
	// httpd.Use(yourMiddleware)

	// 3. 基础路由
	httpd.Get("/", func(w http.ResponseWriter, r *http.Request) {
		httpd.SendSuccess(w, map[string]string{
			"message": "Hello Chi!",
		})
	})

	httpd.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		httpd.SendSuccess(w, "pong")
	})

	httpd.Post("/users", func(w http.ResponseWriter, r *http.Request) {
		httpd.SendSuccess(w, map[string]interface{}{
			"id": 1,
			"name": "张三",
		})
	})

	// 4. 路由组
	httpd.Route("/api/v1", func(r chi.Router) {
		r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
			httpd.SendSuccess(w, []map[string]interface{}{
				{"id": 1, "name": "张三"},
				{"id": 2, "name": "李四"},
			})
		})

		r.Get("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
			id := chi.URLParam(r, "id")
			httpd.SendSuccess(w, map[string]string{
				"id": id,
				"name": "用户" + id,
			})
		})

		r.Post("/users", func(w http.ResponseWriter, r *http.Request) {
			// 业务逻辑...
			httpd.SendError(w, http.StatusBadRequest, "参数错误")
		})
	})

	// 5. 静态文件服务
	// httpd.Static("/static", "./public")
	// httpd.StaticIndex("/files", "./uploads")

	// 6. 嵌入式静态文件
	// //go:embed public/*
	// var embedFS embed.FS
	// httpd.StaticEmbed("/assets", "public", &embedFS)

	// 7. 自定义JSON响应
	httpd.Get("/custom", func(w http.ResponseWriter, r *http.Request) {
		httpd.SendJSON(w, http.StatusCreated, map[string]interface{}{
			"code": 201,
			"data": "资源已创建",
		})
	})

	// 8. 启动服务器
	httpd.Server(":8080", true) // true表示debug模式
}

*/
