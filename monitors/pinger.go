package monitors

import (
	"sync"
	"time"
)

var interval = 30 * time.Second

func Ping(monitors []Monitor) chan interface{} {
	ticker := time.NewTicker(interval)
	quit := make(chan interface{})
	go func() {
		for {
			select {
			case <-ticker.C:
				refresh(monitors)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	refresh(monitors)
	return quit
}

func check(m Monitor, wg *sync.WaitGroup) {
	defer wg.Done()
	m.Check()
}

func refresh(monitors []Monitor) {
	ns := len(monitors)
	var wg sync.WaitGroup
	wg.Add(ns)
	for _, m := range monitors {
		check(m, &wg)
	}
	wg.Wait()
}
