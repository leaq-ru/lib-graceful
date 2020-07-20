package graceful

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func HandleSignals(stopFunc ...func()) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)

	<-signals
	wg := sync.WaitGroup{}
	wg.Add(len(stopFunc))
	for _, fun := range stopFunc {
		go func(f func()) {
			defer wg.Done()
			f()
		}(fun)
	}
	wg.Wait()
}
