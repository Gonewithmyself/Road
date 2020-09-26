package router

import (
	"context"
	"io/ioutil"
	"log"
	"misc/tool/htmlviewer/handler"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	addr string
	s    *http.Server
}

func NewHttpServer(addr string) *HttpServer {
	return &HttpServer{
		addr: addr,
	}
}

func (server *HttpServer) Stop() {
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*30))
	log.Print("stopping rest server")
	err := server.s.Shutdown(ctx)
	if err != nil {
		log.Print("stop timeout")
	}
}

func (server *HttpServer) Start() {
	// gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	r := gin.Default()
	r.Use(setCROSOptions)

	r.GET("/", handler.Home)
	r.GET("/content/:file", handler.Content)
	// r.GET("/api/products/:productId/trades", GetProductTrades)
	// r.GET("/api/products/:productId/book", GetProductOrderBook)
	// r.GET("/api/products/:productId/candles", GetProductCandles)

	server.s = &http.Server{
		Addr:    server.addr,
		Handler: r,
	}

	log.Println("server start ", server.addr)
	err := server.s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func setCROSOptions(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
	c.Header("Content-Type", "application/json")
}
