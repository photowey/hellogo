package web

import (
	"github.com/hellogo/pkg/logger"
	"github.com/valyala/fasthttp"
)

type Handler struct {
	Host string
}

func (h *Handler) Run(ch chan error) {
	go func() {
		logger.Info("Init successful and the server listen on host:[{}]", "192.168.1.101")
		ch <- fasthttp.ListenAndServe(h.Host, h.ServerOnInt)
	}()
}

func (h *Handler) ServerOnInt(ctx *fasthttp.RequestCtx) {
}

func (h *Handler) ServerOnExt(ctx *fasthttp.RequestCtx) {
}

func route(ctx *fasthttp.RequestCtx, req *Request) {
}
