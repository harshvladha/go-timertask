package gotimertask

import "time"

type TimerTask struct {
	task     *Task
	duration time.Duration
	ticker   *time.Ticker
	exit     chan bool
	stopped  bool
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
func NewTask(f func()) *Task {
	run := func(interface{}) {
		f()
	}

	return &Task{
		run: run,
	}
}

// Stops the given timer task from further execution
// Once stopped it won't run again
// and again invoking Stop() doesn't do anything
func (t *TimerTask) Stop() {
	if !t.stopped {
		t.exit <- true
		t.stopped = true
		close(t.exit)
	}
}

func newTimerTask(t *Task, d time.Duration) *TimerTask {
	return &TimerTask{
		task:     t,
		duration: d,
		ticker:   time.NewTicker(d),
		exit:     make(chan bool),
		stopped:  false,
	}
}

// Schedules a function `f` to run at a `d` Duration
func Schedule(t *Task, d time.Duration) *TimerTask {
	timerTask := newTimerTask(t, d)
	taskInvoker(timerTask)
	return timerTask
}

func taskInvoker(t *TimerTask) {
	go func() {
		for {
			select {
			case <-t.ticker.C:
				t.task.run(t.task.data)
			case <-t.exit:
				t.ticker.Stop()
				return
			}
		}
	}()
}
