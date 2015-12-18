package monitors

import (
	"sync"
	"time"
)

var interval = 30 * time.Second

type Monitoreable interface {
	GetGroups() []Group
	Failed(Monitor)
}

func Ping(subject Monitoreable) chan interface{} {
	ticker := time.NewTicker(interval)
	quit := make(chan interface{})
	go func() {
		for {
			select {
			case <-ticker.C:
				for _, g := range subject.GetGroups() {
					refresh(g.Monitors, subject)
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	for _, g := range subject.GetGroups() {
		refresh(g.Monitors, subject)
	}
	return quit
}

func check(m Monitor, wg *sync.WaitGroup) {
	defer wg.Done()
	m.Check()
}

func refresh(monitors []Monitor, subject Monitoreable) {
	ns := len(monitors)
	var wg sync.WaitGroup
	wg.Add(ns)
	for _, m := range monitors {
		check(m, &wg)
		if AllFailed(m) {
			go subject.Failed(m)
		}
	}
	wg.Wait()
}
