package timertask

import "time"

type TimerTask struct {
	task     Task
	duration time.Duration
}

type Task struct {
	run  func(interface{})
	data interface{}
}

// Creates a new Task with argument which needs to be passed to a function to be scheduled
func NewTaskWithArgument(f func(interface{}), arg interface{}) *Task {
	return &Task{
		run:  f,
		data: arg,
	}
}

// Creates a new Task without argument which needs to be passed to a function to be scheduled
func NewTask(f func(interface{})) *Task {
	run := func(interface{}) {
		f()
	}

	return &Task{
		run: run,
	}
}

func newTimerTask(t Task, d time.Duration) *TimerTask {
	return &TimerTask{
		task:     t,
		duration: d,
	}
}

// Schedules a function `f` to run at a `d` Duration
func Schedule(t Task, d time.Duration) {
	timerTask := newTimerTask(t, d)
	ticker := time.NewTicker(timerTask.duration)
	taskInvoker(ticker, t)
}

func taskInvoker(ticker *time.Ticker, t Task) {
	go func() {
		for {
			select {
			case <-ticker.C:
				t.run(t.data)
			}
		}
	}()
}
