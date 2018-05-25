package timertask

import "time"

type TimerTask struct {
	function func(interface{})
	duration time.Duration
}

func newTimerTask(f func(interface{}), d time.Duration) *TimerTask {
	return &TimerTask{
		function: f,
		duration: d,
	}
}

// Schedules a function `f` to run at a `d` Duration
func Schedule(f func(interface{}), d time.Duration) {
	timerTask := newTimerTask(f, d)
	ticker := time.NewTicker(timerTask.duration)
	taskInvoker(ticker, f)
}

func taskInvoker(ticker *time.Ticker, f func(interface{})) {
	go func() {
		for {
			select {
			case <-ticker.C:
				f()
			}
		}
	}()
}
