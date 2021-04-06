package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/reactivex/rxgo/v2"
)

func main() {

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
