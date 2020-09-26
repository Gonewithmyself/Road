package handler

import (
	"log"
	"os"
	"path/filepath"
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
	ctx.String(200, pd+"__"+id)
}

type file struct {
	Name  string
	Class string
}

var m = map[string][]*file{}

func listDir() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		f := &file{
			Name: info.Name(),
		}

		if info.IsDir() {
			f.Class = "dir"
		}

		ff, _ := m[""]
		ff = append(ff, f)
		m[""] = ff
		return nil
	})
}

func render(tpl string, data interface{}) string {
	var b strings.Builder
	if err := tpls[tpl].Execute(&b, data); err != nil {
		log.Println("render, ", err)
		return ""
	}

	return b.String()
}
