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

	// var list []string
	// filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
	// 	if info.IsDir() {
	// 		return nil
	// 	}
	// 	fname := info.Name()
	// 	list = append(list, fname)
	// 	log.Println(filepath.Dir(path), fname)
	// 	return nil
	// })
	ctx.String(200, render("home", m["."]))
}

func Content(ctx *gin.Context) {
	pd := ctx.Param("file")
	class := ctx.Query("id")
	// fpath := ctx.Query("fpath")
	// log.Println(class, " --- ", fpath)

	if class == "dir" {
		ctx.Header("Content-Type", "text/html; charset=utf-8")
		ctx.String(200, render("home", m[pd]))
		return
	}

	f, err := os.OpenFile(fm[pd], os.O_RDONLY, 0666)
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
	Path  string
}

var m = map[string][]*file{}
var fm = map[string]string{}

func listDir() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		f := &file{
			Name: info.Name(),
			Path: filepath.Dir(path),
		}

		if info.IsDir() {
			if f.Path != "." || f.Name == "." {
				return nil
			}
			f.Class = "dir"
		} else if filepath.Ext(f.Name) != ".html" {
			return nil
		}

		ff, _ := m[f.Path]
		ff = append(ff, f)
		m[f.Path] = ff
		fm[f.Name] = f.Path + "/" + f.Name
		return nil
	})
}

func init() {
	listDir()
}

func render(tpl string, data interface{}) string {
	var b strings.Builder
	if err := tpls[tpl].Execute(&b, data); err != nil {
		log.Println("render, ", err)
		return ""
	}

	return b.String()
}
