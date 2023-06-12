package web

import (
	"net"
	"net/http"
)

type HandleFunc func(ctx Context)

// 确保结构体一定实现了 Server 接口
var _ Server = &HTTPServer{}

type Server interface {
	http.Handler
	Start() error

	AddRoute(method string, path string, handleFunc HandleFunc) // 路由注册
}

type HTTPServer struct {
	Addr string
}

// ServeHTTP 处理请求入口
func (h *HTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// 构建 Context
	ctx := &Context{
		Req:  request,
		Resp: writer,
	}
	h.Serve(ctx)
}

func (h *HTTPServer) Serve(ctx *Context) {
	// 查找路由，并命中逻辑
}

func (h *HTTPServer) AddRoute(method string, path string, handleFunc HandleFunc) {
	panic("implement me")
}

func (h *HTTPServer) Get(path string, handleFunc HandleFunc) {

}

func NewHTTPServer(addr string) *HTTPServer {
	return &HTTPServer{
		Addr: addr,
	}
}

func (h *HTTPServer) Start() error {
	listen, err := net.Listen("tcp", h.Addr)
	if err != nil {
		return err
	}
	return http.Serve(listen, h)
}