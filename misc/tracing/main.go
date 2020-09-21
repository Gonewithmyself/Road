package main

import (
	"bytes"
	"io"
	"log"
	"math/rand"
	"misc/tracing/tracer"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/opentracing/opentracing-go"
)

func main() {
	var (
		err error
		io  io.Closer
	)
	//创建tracer对象
	tracer.Tracer, io, err = tracer.NewTracer("usersvc", "127.0.0.1:6831")
	if err != nil {
		log.Fatalf("tracer.NewTracer error(%v)", err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(tracer.Tracer)

	//server
	http.HandleFunc("/getIP", getIP)
	http.HandleFunc("/getuser", getUser)
	http.HandleFunc("/pulluser", pullUser)

	log.Printf("Starting server on port %d", 8002)
	// err = http.ListenAndServe(
	// 	fmt.Sprintf(":%d", 8002),
	// 	// use nethttp.Middleware to enable OpenTracing for server
	// 	nethttp.Middleware(tracer.Tracer, http.DefaultServeMux))

	err = http.ListenAndServe(":8002", nil)
	if err != nil {
		log.Fatalf("Cannot start server: %s", err)
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	var rsp = "failed"
	log.Print("Received getUser request")
	span := tracer.Tracer.StartSpan("getUser")
	defer func() {
		span.Finish()
		w.Write([]byte(rsp))
	}()

	var b bytes.Buffer
	err := tracer.InjectBin(span, &b)
	if err != nil {
		log.Print("inject bin", err)
		return
	}

	n := rand.Intn(300) + 20
	time.Sleep(time.Millisecond * time.Duration(n))

	n = rand.Intn(5) + 1
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			i := i
			b := b
			wg.Add(1)
			go func() {
				doHttp(i, b)
				wg.Done()
			}()
			continue
		}
		doHttp(i, b)
	}
	wg.Wait()

	// b.Reset()
	// _, err = b.ReadFrom(resp.Body)
	// if err != nil {
	// 	log.Print("read from", err)
	// 	return
	// }

	rsp = "success"
}

func doHttp(id int, b bytes.Buffer) {
	req, _ := http.NewRequest("POST", "http://localhost:8002/pulluser", &b)
	req.Header.Set("id", strconv.Itoa(id))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print("http post", err)
		return
	}
	defer resp.Body.Close()
}

func pullUser(w http.ResponseWriter, r *http.Request) {
	var rsp = "failed"
	log.Print("Received pullUser request")
	spanCtx, err := tracer.ExtractBin(r.Body)
	if err != nil {
		log.Print("extract span", err)
		return
	}

	method := r.Header.Get("id")
	id, _ := strconv.Atoi(method)
	fn := opentracing.FollowsFrom
	if id%2 == 1 {
		fn = opentracing.ChildOf
	}

	span := tracer.Tracer.StartSpan(
		method,
		fn(spanCtx),
	)
	defer func() {
		span.Finish()
		w.Write([]byte(rsp))
	}()

	n := rand.Intn(300) + 20
	time.Sleep(time.Millisecond * time.Duration(n))

}

func getIP(w http.ResponseWriter, r *http.Request) {
	log.Print("Received getIP request")

}

func init() {
	rand.Seed(time.Now().Unix())
}

/*

UI localhost:16686

docker run -d --name jaeger \
  -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14268:14268 \
  -p 14250:14250 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.19
*/
