package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	lru "github.com/hashicorp/golang-lru"
)

func main() {
	fmt.Println("hello")

	lrueg()

	// url := `http://file.enfamily.cn/nce/book3/NCE3001.mp3`
	// resp, err := http.Get(url)
	// if err != nil {
	// 	panic(err)
	// }
	// // ioutil.ReadAll()
	// data, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }

	// err = ioutil.WriteFile("test.mp3", data, 0755)
	// if err != nil {
	// 	panic(err)
	// }

	// prefix := "NCE30"
	// max := 60
	// for i := 1; i <= max; i++ {
	// 	num := fmt.Sprintf("%d", i)
	// 	if len(num) == 1 {
	// 		num = "0" + num
	// 	}
	// 	num = prefix + num
	// 	fmt.Println(num)
	// 	curl(num)
	// }
	// renameFile()
}

func curl(name string) {
	fname := name + ".mp3"
	url := "http://file.enfamily.cn/nce/book3/" + fname

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(fname, data, 0755)
	if err != nil {
		panic(err)
	}
}

func renameFile() {
	os.Rename("misc/test.txt", "misc/text.json")
	filepath.Walk("./nce3_en", func(name string, info os.FileInfo, err error) error {
		//	fname := filepath.Base(name)
		// fmt.Println(name)

		to := strings.ReplaceAll(name, "NCE", "NewConceptEnglish")
		os.Rename(name, to)
		return nil
	})
}

func lrueg() {
	c, _ := lru.New(10)

	for i := 0; i < 100; i++ {
		c.Add(i, i)
	}

	fmt.Println(c.Contains(1))
	fmt.Println(c.Contains(99))
}

func fb(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	return fb(n-2) + fb(n-1)
}

func fb2(n int) int {
	return dofb(n, make(map[int]int))
}

func dofb(n int, m map[int]int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	if res, ok := m[n]; ok {
		return res
	}
	res := fb(n-2) + fb(n-1)
	m[n] = res
	return res
}
