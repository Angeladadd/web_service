package main

import (
	"net/http"
	"io/ioutil"
	"log"
	"time"
	"fmt"
)

func echo(wr http.ResponseWriter, r *http.Request) {
	msg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		wr.Write([]byte("echo error"))
		return
	}
	writeLen, err := wr.Write([]byte(fmt.Sprintf("At %s, we got %s", time.Now().String(), msg)))
	if err != nil || writeLen != len(msg) {
		log.Println(err, "write len:", writeLen)
	}
}

func main() {
	http.HandleFunc("/", echo)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}