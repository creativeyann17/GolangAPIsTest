package fasthttp

import (
	"log"

	"github.com/valyala/fasthttp"
)

type Server struct {
	srv *fasthttp.Server
}

func Start(port string) *Server {
	handler := func(ctx *fasthttp.RequestCtx) {
		if string(ctx.Path()) == "/hello" && string(ctx.Method()) == "GET" {
			ctx.SetContentType("application/json")
			ctx.SetStatusCode(fasthttp.StatusOK)
			ctx.WriteString(`{"message":"Hello, World!","framework":"FastHTTP"}`)
		} else {
			ctx.SetStatusCode(fasthttp.StatusNotFound)
		}
	}

	srv := &fasthttp.Server{
		Handler: handler,
	}

	server := &Server{srv: srv}

	go func() {
		if err := srv.ListenAndServe(port); err != nil {
			log.Printf("FastHTTP server error: %v", err)
		}
	}()

	return server
}

func (s *Server) Stop() error {
	return s.srv.Shutdown()
}

func Main() {
	log.Println("FastHTTP server starting on :8080")
	server := Start(":8080")
	defer server.Stop()

	select {}
}
