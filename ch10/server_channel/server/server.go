package server

import "log"

type Work struct {
	Op    func(int, int) int
	A, B  int
	Reply chan int // Server sends result on this channel.
}

// New creates a new server that accepts Work requests
// through the req channel.
func New() (req chan<- *Work) {
	wc := make(chan *Work)
	go serve(wc)
	return wc
}

func serve(wc <-chan *Work) {
	for w := range wc {
		go safelyDo(w)
	}
}

func safelyDo(w *Work) {
	// Regain control of a panicking goroutine to avoid
	// killing the other executing goroutines.
	defer func() {
		if err := recover(); err != nil {
			log.Println("work failed:", err)
		}
	}()
	do(w)
}

func do(w *Work) {
	w.Reply <- w.Op(w.A, w.B)
}
