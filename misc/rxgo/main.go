package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/reactivex/rxgo/v2"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {})

	go http.ListenAndServe(":8080", nil)

	var ntf chan<- rxgo.Item
	observable := rxgo.Create([]rxgo.Producer{func(ctx context.Context, next chan<- rxgo.Item) {
		next <- rxgo.Of(1)
		next <- rxgo.Of(2)
		next <- rxgo.Of(3)
		next <- rxgo.Error(errors.New("unknown"))
		next <- rxgo.Of(4)
		next <- rxgo.Of(5)
		ntf = next
	}}, rxgo.WithBufferedChannel(10))

	go func() {
		fmt.Println("???")
		ntf <- rxgo.Of(9)
		fmt.Println("???")
	}()

	ch := observable.Observe()
	for item := range ch {
		if item.Error() {
			fmt.Println("error:", item.E)
		} else {
			fmt.Println(item.V)
		}
	}

	time.Sleep(time.Second * 5)
	fmt.Println("exit")
}

type myint int64

func testarray(arr []interface{}, a int64) bool {
	testFn := func(item interface{}) bool {
		switch num := item.(type) {
		case int64:
			return true
		case int32:
			return num == int32(a)
		case float64:
			return int64(num) == a
		case int:
			return int64(num) == a
		default:
			tp := reflect.TypeOf(num)
			v := reflect.ValueOf(num)
			switch tp.Kind() {
			case reflect.Float32, reflect.Float64:
				return int64(v.Float()) == a

			case reflect.Int64, reflect.Int8:
				return v.Int() == a

			case reflect.Uint64, reflect.Uint16:
				return int64(v.Uint()) == a
			}
		}
		return false
	}
	for i := range arr {
		if testFn(arr[i]) {
			return true
		}
	}
	return false
}
