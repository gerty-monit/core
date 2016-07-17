package gerty

import (
	"sync"
	"time"
)

var interval = 30 * time.Second

type Monitoreable interface {
	GetGroups() []Group
	Failed(Monitor)
	Restored(Monitor)
}

func Ping(subject Monitoreable) chan interface{} {
	ticker := time.NewTicker(interval)
	quit := make(chan interface{})
	go func() {
		for {
			select {
			case <-ticker.C:
				refreshGroups(subject)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	refreshGroups(subject)
	return quit
}

func refreshGroups(subject Monitoreable) {
	groups := subject.GetGroups()
	for i := range groups {
		refresh(groups[i].Monitors, subject)
	}
}

func check(m Monitor, wg *sync.WaitGroup) {
	defer wg.Done()
	m.Check()
}

func refresh(monitors []Monitor, subject Monitoreable) {
	var wg sync.WaitGroup
	wg.Add(len(monitors))

	for i := range monitors {
		go func(i int) {
			monitor := monitors[i]
			check(monitor, &wg)

			if AllFailed(monitor) && !monitor.IsTripped() {
				monitor.Trip()
				go subject.Failed(monitor)
			}

			if AllOk(monitor) && monitor.IsTripped() {
				monitor.Untrip()
				go subject.Restored(monitor)
			}
		}(i)
	}

	wg.Wait()
}
