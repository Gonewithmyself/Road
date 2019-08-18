package db

import "C"
import (
	"encoding/json"
	"fmt"
	"runtime/debug"
	"sync"
	"time"

	"github.com/globalsign/mgo"
)

/*
db.getProfilingStatus()
db.system.profile.find()
db.stats(1024)
*/

// -buildmode=c-shared
// -buildmode=c-archive
var gMgoSession *mgo.Session

var dbName string
var o sync.Once

//export InitSession
func InitSession(dbname, host, username, password *C.char) {
	o.Do(func() {
		doInitSession(dbname, host, username, password)
	})
}

func doInitSession(x, y, z, w *C.char) {
	dbname, host, username, password :=
		C.GoString(x), C.GoString(y), C.GoString(z), C.GoString(w)
	dbName = dbname
	dialInfo := mgo.DialInfo{
		Addrs:    []string{host},
		Source:   dbName,
		Database: dbName,
		Username: username,
		Password: password,
		Timeout:  time.Duration(5) * time.Second,
	}

	var err error
	gMgoSession, err = mgo.DialWithInfo(&dialInfo)
	if nil != err {
		fmt.Println("reset mongo session", "error", err, dbname, host, username, password)
		return
	}

	gMgoSession.SetMode(mgo.Strong, true)
}

// put data to gridfs
// 1. remove origin gridfs file
// 2. write data to gridfs
//export WriteGridFS
func WriteGridFS(sname, sdata *C.char) {
	defer func() {
		r := recover()
		if nil != r {
			debug.PrintStack()
			// panic(r)
		}
	}()
	name, data := C.GoString(sname), []byte(C.GoString(sdata))
	// copy mgo.v2 session
	sess := CopySession()
	defer sess.Close()

	// open gridfs
	fs := sess.DB(dbName).GridFS("fs")
	originFile, err := doesExistFS(fs, name)
	if originFile != nil {
		defer originFile.Close()
	}

	if err != nil {
		fmt.Println("open mongofs", "file", name, "error", err)
		return
	}

	// create new file
	file, err := fs.Create(name)
	if nil != err {
		fmt.Println("create new mongofiles", "file", name, "error", err)
		return
	}
	defer file.Close()

	// write data to db
	_, err = file.Write(data)
	if nil != err {
		fmt.Println("gridfs write", "file", name, "error", err)
		return
	}

	if nil == originFile {
		return
	}

	// remove origin grid file
	err = fs.RemoveId(originFile.Id())
	if nil != err {
		fmt.Println("gridfs revome origin file", "file", name, "error", err)
	}

	return
}

// does gridfs exist
func doesExistFS(fs *mgo.GridFS, name string) (*mgo.GridFile, error) {
	file, err := fs.Open(name)

	if nil == err || "not found" == err.Error() {
		return file, nil
	}

	return nil, err
}

func CopySession() *mgo.Session {
	return gMgoSession.Copy()
}

type user struct {
	Id   int32
	Name string
}

func Find(c, query string) {

	dd := []byte(query)

	var dst map[string]interface{}
	json.Unmarshal(dd, &dst)

	m := 0 // dst.Data.(map[string]interface{})
	fmt.Println(dst, m)

	var res []interface{}
	gMgoSession.DB("s1001").C(c).Find(dst).All(&res)
	data, _ := json.Marshal(res)
	fmt.Println(string(data))

}

func init() {
	dbname, host, username, psw :=
		C.CString("s1001"), C.CString("192.168.15.93:27017"), C.CString("test"), C.CString("123456")
	InitSession(dbname, host, username, psw)
}

// gMgoSession.DB("ss").C("").Find
