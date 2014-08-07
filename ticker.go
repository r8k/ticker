package ticker

import "time"

func TickEvery(fn func(args []interface{}), d time.Duration, args []interface{}) chan bool {
	// to handle states, we create a channel
	// and use that to communicate about the
	// state of a ticker, what it should be
	// in, at the given moment: `run` vs `stop`
	stop := make(chan bool, 1)
	// create a go routine to spawn off ticker
	go func() {
		// create a ticker, instead of a tick
		// to ensure that we can stop and start
		// @see: http://golang.org/src/pkg/time/tick.go#L49
		ticker := time.NewTicker(d)

		// create a loop, to handle ticks
		// and run the given function.
		for {
			select {
			// for lifetime of ticker channel,
			case <-ticker.C:
				// execute the given function
				fn(args)

			// listen on a channel, and if true,
			// stop the ticker, and close the channel
			case <-stop:
				ticker.Stop()
				close(stop)
				return
			}
		}
	}()

	return stop
}
