package main

// import "C"
// import (
// 	// "radiumme.com/dragon/core/db"
// 	"fmt"
// 	"io/ioutil"
// 	"net"

// 	"misc/db"
// )

// func main() {

// //export gotest
// func gotest(s *C.char) {
// 	defer func() {
// 		fmt.Println("i am in defer")
// 	}()
// 	fmt.Println("this is go ", C.GoString(s))
// }

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"misc/tool"
	"net"
)

func main() {
	l, err := net.Listen("tcp4", "localhost:45000")
	if nil != err {
		panic(err)
	}

	fmt.Println("listening...")
	for {
		conn, err := l.Accept()
		if nil != err {
			panic(err)
		}
		fmt.Println("serv")
		Server(conn)
		conn.Close()
	}
}

func Server(conn net.Conn) {
	buf, err := ioutil.ReadAll(conn)
	if nil != err {
		return
	}

	p := &tool.Packet{}
	// fmt.Println(string(buf))
	err = json.Unmarshal(buf, p)
	if nil != err {
		fmt.Println(err, string(buf))
		return
	}
	fn := tool.Find
	if p.M == "fs" {
		fn = tool.ReadFs
	}
	res := fn(p)

	_, err = conn.Write(res)
	if nil != err {
		fmt.Println(err)
	}
}
