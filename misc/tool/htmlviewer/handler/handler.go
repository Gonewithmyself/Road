package handler

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func Home(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html; charset=utf-8")
	ctx.String(200, render("home", []string{"test", "go"}))
}

func Content(ctx *gin.Context) {
	pd := ctx.Param("file")
	id := ctx.Query("id")
	ctx.String(300, pd+"__"+id)
}

func render(tpl string, data interface{}) string {
	var b strings.Builder
	if err := tpls[tpl].Execute(&b, data); err != nil {
		log.Println("render, ", err)
		return ""
	}

	return b.String()
}
