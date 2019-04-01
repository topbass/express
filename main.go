package main

import (
	"log"

	"github.com/valyala/fasthttp"
)

func main() {
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		c := &fasthttp.HostClient{
			Addr: "localhost:9000",
		}
		if err := c.Do(&ctx.Request, &ctx.Response); err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		}
	}
	s := &fasthttp.Server{
		Handler: requestHandler,
	}
	if err := s.ListenAndServe("127.0.0.1:8080"); err != nil {
		log.Fatalf("error in ListenAndServe: %s", err)
	}
}
