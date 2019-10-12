package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type play struct {
	PlayID   string
	Audience int64
}

type invoice struct {
	Customer    string
	Performaces []*play
}

func readPlays() (m map[string]map[string]string) {
	src, err := ioutil.ReadFile("plays.json")
	if nil != err {
		log.Fatal(err, "read file")
	}

	json.Unmarshal(src, &m)
	return
}

func readInvoice() (m []*invoice) {
	src, err := ioutil.ReadFile("invoices.json")
	if nil != err {
		log.Fatal(err, "read file")
	}

	json.Unmarshal(src, &m)
	return
}

var (
	plays    map[string]map[string]string
	invoices []*invoice
)

func init() {
	plays = readPlays()
	invoices = readInvoice()

	fmt.Print(plays)
}
