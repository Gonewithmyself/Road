package handler

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func Home(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html; charset=utf-8")

	var list []string
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || filepath.Ext(path) != ".html" {
			return nil
		}
		fname := info.Name()
		list = append(list, fname)
		return nil
	})
	ctx.String(200, render("home", list))
}

func Content(ctx *gin.Context) {
	pd := ctx.Param("file")
	id := ctx.Query("id")
	_ = id

	f, err := os.OpenFile(pd, os.O_RDONLY, 0666)
	if err != nil {
		ctx.String(400, "%v not found", pd)
		return
	}

	data, _ := ioutil.ReadAll(f)
	ctx.Header("Content-Type", "text/html; charset=utf-8")
	ctx.String(200, string(data))
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
