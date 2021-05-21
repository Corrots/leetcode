package main

import (
	"net/http"
	"runtime"
	"runtime/pprof"
	"strconv"
)

func write(w http.ResponseWriter, data []byte) {
	_, _ = w.Write(data)
}

func count(w http.ResponseWriter, r *http.Request) {
	write(w, []byte(strconv.Itoa(runtime.NumGoroutine())))
}

func main() {
	http.HandleFunc("/_count", count)
	http.HandleFunc("/query", query)
	http.HandleFunc("/_goroutine", goroutineStack)
	http.ListenAndServe(":6080", nil)
}

func goroutineStack(w http.ResponseWriter, r *http.Request) {
	_ = pprof.Lookup("goroutine").WriteTo(w, 1)
}

func query(w http.ResponseWriter, r *http.Request) {
	c := make(chan byte)

	go func() {
		c <- 0x31
	}()

	go func() {
		c <- 0x32
	}()

	go func() {
		c <- 0x33
	}()

	rs := make([]byte, 0)
	for i := 0; i < 2; i++ {
		rs = append(rs, <-c)
	}

	write(w, rs)
}
