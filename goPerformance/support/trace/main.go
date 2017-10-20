package main

import (
	"net/http"
	_ "net/http/pprof"
	"time"

	_ "github.com/mkevac/debugcharts"
)

func CPUHogger() {
	var acc uint64
	t := time.Tick(2 * time.Second)
	for {
		select {
		case <-t:
			time.Sleep(50 * time.Millisecond)
		default:
			acc++
		}
	}
}

func main() {
	go CPUHogger()
	go CPUHogger()
	http.ListenAndServe("0.0.0.0:8181", nil)
}