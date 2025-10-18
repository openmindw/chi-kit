# Chi HTTP 封装库

基于 `github.com/go-chi/chi/v5` 的HTTP服务封装库，提供简洁易用的API接口。

## 📦 功能特性

- ✅ 简洁的路由注册（Get, Post, Put, Delete, Patch）
- ✅ 路由组和中间件支持
- ✅ 静态文件服务（支持embed.FS）
- ✅ 统一的JSON响应格式
- ✅ 优雅关闭机制
- ✅ 集成日志系统
- ✅ 请求日志中间件

## 🚀 快速开始

### 基础使用

```go
package main

import (
    "net/http"
    "github.com/opentdp/go-helper/httpd"
)

func main() {
    // 初始化引擎
    httpd.Engine(true) // true = debug模式

    // 注册路由
    httpd.Get("/", func(w http.ResponseWriter, r *http.Request) {
        httpd.SendSuccess(w, "Hello World!")
    })

    // 启动服务
    httpd.Server(":8080")
}
```

### 路由组

```go
httpd.Route("/api/v1", func(r chi.Router) {
    r.Get("/users", listUsers)
    r.Post("/users", createUser)
    r.Get("/users/{id}", getUser)
})
```

### 响应方法

```go
// 成功响应
httpd.SendSuccess(w, data)
// 输出: {"success": true, "data": ...}

// 错误响应
httpd.SendError(w, http.StatusBadRequest, "参数错误")
// 输出: {"success": false, "message": "参数错误"}

// 自定义JSON响应
httpd.SendJSON(w, http.StatusCreated, customData)
```

### 静态文件

```go
// 普通静态文件
httpd.Static("/static", "./public")

// 带目录索引
httpd.StaticIndex("/files", "./uploads")

// 嵌入式文件系统
//go:embed public/*
var embedFS embed.FS
httpd.StaticEmbed("/assets", "public", &embedFS)
```

## 📝 API对比

| Gin版本 | Chi版本 | 说明 |
|---------|---------|------|
| `gin.Context` | `http.ResponseWriter, *http.Request` | 上下文对象 |
| `c.JSON()` | `httpd.SendJSON()` | JSON响应 |
| `engine.Group()` | `httpd.Route()` | 路由组 |
| `gin.HandlerFunc` | `http.HandlerFunc` | 处理函数 |

## 🔧 主要组件

### engine.go
- `Engine(debug bool)` - 初始化路由引擎
- `Route(pattern, fn)` - 创建路由组
- `Use(middlewares...)` - 注册中间件
- `Get/Post/Put/Delete/Patch` - HTTP方法快捷注册

### server.go
- `Server(addr, options...)` - 启动HTTP服务器
- 自动集成优雅关闭（9秒超时）

### static.go
- `Static(prefix, root)` - 普通静态文件服务
- `StaticIndex(prefix, root)` - 带索引的静态文件
- `StaticEmbed(prefix, sub, efs)` - 嵌入式文件系统

### response.go
- `SendSuccess(w, data)` - 成功响应
- `SendError(w, statusCode, message)` - 错误响应
- `SendJSON(w, statusCode, data)` - 通用JSON响应

### middleware.go
- `LoggerMiddleware(debug)` - 请求日志中间件

## 🎯 设计理念

1. **保持API一致性** - 与Gin版本保持相似的使用体验
2. **标准化响应** - 统一的JSON响应格式
3. **简化开发** - 封装复杂性，提供简洁API
4. **生产就绪** - 内置日志、优雅关闭等生产特性

## 📄 License

与主项目保持一致
