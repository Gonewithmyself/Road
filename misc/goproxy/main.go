package main

import (
	"net/http"
	"os"

	"github.com/goproxy/goproxy"
	"github.com/goproxy/goproxy/cacher"
)

func main() {
	g := goproxy.New()
	g.GoBinEnv = append(
		os.Environ(),
		"GOPROXY=https://goproxy.cn,direct", // 使用 goproxy.cn 作为上游代理
		"GOPRIVATE=git.example.com",         // 解决私有模块的拉取问题（比如可以配置成公司内部的代码源）
	)
	g.Cacher = &cacher.Disk{Root: "/data/gocache"}
	// g.ProxiedSUMDBs = []string{"sum.golang.org https://goproxy.cn/sumdb/sum.golang.org"} // 代理默认的校验和数据库
	http.ListenAndServe("localhost:8081", g)
	// // go env -w GOPROXY=https://goproxy.cn,direct
	// // go env -w GOPROXY=http://127.0.0.1:8080,direct
}
