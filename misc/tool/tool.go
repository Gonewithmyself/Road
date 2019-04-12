package tool

import (
	"encoding/json"
	"fmt"
	"runtime/debug"
	"strings"
	"sync"
	"time"

	"github.com/globalsign/mgo"
)

var gMgoSession *mgo.Session

var dbName string
var o sync.Once

type Auth struct {
	Db     string
	User   string
	Psw    string
	Host   string
	Source string
	Rs     string
}

type Packet struct {
	C string
	A *Auth
	Q map[string]interface{}
	S map[string]interface{}
	M string
}

//export InitSession
func InitSession(dbname, host, username, password string) {
	o.Do(func() {
		doInitSession(dbname, host, username, password, dbname, "")
	})
}

func doInitSession(dbname, host, username, password, rs, source string) {
	dbName = dbname
	hosts := strings.Split(host, ",")

	fmt.Println(rs, hosts)
	dialInfo := mgo.DialInfo{
		Addrs:          hosts,
		Source:         source,
		Database:       dbName,
		Username:       username,
		Password:       password,
		Timeout:        time.Duration(5) * time.Second,
		ReplicaSetName: rs,
	}

	var err error
	gMgoSession, err = mgo.DialWithInfo(&dialInfo)
	if nil != err {
		fmt.Println("reset mongo session", "error", err, dbname, host, username, password)
		return
	}

	gMgoSession.SetMode(mgo.Strong, true)
}

func CopySession(auth *Auth) *mgo.Session {
	if dbName != auth.Db {
		doInitSession(auth.Db, auth.Host,
			auth.User, auth.Psw, auth.Rs, auth.Source)
		// fmt.Println(gMgoSession.DatabaseNames())
	}
	return gMgoSession.Clone()
}

func Find(p *Packet) (res []byte) {
	defer func() {
		r := recover()
		if nil != r {
			fmt.Println(r)
			debug.PrintStack()
			fmt.Println(p)
		}
	}()
	var (
		dst = p.Q
		err error
	)

	var list []interface{}
	sess := CopySession(p.A)
	defer sess.Close()

	// fmt.Println(p)
	sess.DB(dbName).C(p.C).Find(dst).Select(p.S).All(&list)
	res, err = json.Marshal(list)
	if nil != err {
		fmt.Println(err, list)
	}
	return
}

func ReadFs(p *Packet) []byte {
	defer func() {
		r := recover()
		if nil != r {
			fmt.Println(r)
			debug.PrintStack()
			fmt.Println(p)
		}
	}()

	sess := CopySession(p.A)
	defer sess.Close()
	// open gridfs
	fs := sess.DB(dbName).GridFS("fs")
	file, err := fs.Open(p.C)
	if nil != err {
		fmt.Println("open gridfs error", "name", p, "error", err)
		return nil
	}
	defer file.Close()

	// read
	data := make([]byte, file.Size())
	_, err = file.Read(data)
	if nil != err {
		fmt.Println("read gridfs error", "name", p, "error", err)
		return nil
	}

	return data
}
