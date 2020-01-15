package main

import (
	"crypto/sha1"
	"encoding/hex"
	"time"

	"github.com/thanhpk/randstr"
)

func genStrings(count int) []string {
	var r []string
	for i := 0; i < count; i++ {
		token := randstr.String(16)
		r = append(r, token)
	}
	return r
}

func doJob(s string) string {
	// sleep here to simulate a slow job
	time.Sleep(100 * time.Millisecond)

	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func doJobs(items []string) []string {
	var results []string
	for _, s := range items {
		r := doJob(s)
		results = append(results, r)
	}
	return results
}

type out struct {
	ID  int
	Val string
}
type worker struct {
	ID     int
	input  chan string
	output chan out
}

func (w *worker) DoJob() {
	for v := range w.input {
		w.output <- out{ID: w.ID, Val: doJob(v)}
	}
}
func (w *worker) Send(in string) {
	w.input <- in
}
func doJobsFanout(items []string) []string {
	var results []string
	//use limits to control the amounts of worker
	limits := 10
	l := len(items)
	index := 0
	finished := 0
	reciever := make(chan out, limits)
	defer func() { close(reciever) }()
	wkList := make(map[int]*worker)
	notify := make(chan int, limits)
	//create limit worker
	for i := 0; i < l && i < limits; i++ {
		wk := &worker{ID: i, input: make(chan string), output: reciever}
		wkList[i] = wk
		go wk.DoJob()
		wk.Send(items[i])
		index++
	}
	//set notify
	go func() {
		for v := range notify {
			if index < l {
				wkList[v].Send(items[index])
				index++
			} else {
				return
			}
		}
	}()
	//keep recieved output and send input
	for {
		select {
		case v, ok := <-reciever:
			if !ok {
				break
			}
			finished++
			results = append(results, v.Val)
			if finished == l {
				return results
			}
			notify <- v.ID
		}
	}
}
