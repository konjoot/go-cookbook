package http_client

import (
	"io"
	"log"
	"net/http"
	"runtime"
	"time"
)

func Srv() chan struct{} {
	ch := make(chan struct{})

	go func() {
		go func() { time.Sleep(500 * time.Millisecond); close(ch) }()

		http.HandleFunc("/hello", hello)
		err := http.ListenAndServeTLS("localhost:10443", "server.pem", "server.key", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()

	runtime.Gosched()

	return ch
}

func hello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}
