package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync/atomic"
)

var counter int64

func main() {
	port := os.Getenv("APP_PORT")
	if len(port) == 0 {
		port = "8089"
	}

	http.HandleFunc("/count", func(writer http.ResponseWriter, request *http.Request) {
		val := atomic.AddInt64(&counter, 1)
		writer.Write([]byte(fmt.Sprintf("counter:%v", val)))
	})
	log.Printf("start app at port:%v\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		panic(fmt.Sprintf("fail to run http server:%s", err))
	}
}
