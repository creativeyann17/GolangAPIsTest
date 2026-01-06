package hertz

import (
	"context"
	"log"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type Server struct {
	hertz *server.Hertz
}

func Start(port string) *Server {
	h := server.New(server.WithHostPorts(port))

	h.GET("/hello", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, map[string]string{
			"message":   "Hello, World!",
			"framework": "Hertz",
		})
	})

	server := &Server{hertz: h}

	go func() {
		h.Spin()
	}()

	return server
}

func (s *Server) Stop() error {
	return s.hertz.Shutdown(context.Background())
}

func Main() {
	log.Println("Hertz server starting on :8080")
	server := Start(":8080")
	defer server.Stop()

	select {}
}
