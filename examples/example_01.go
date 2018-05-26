package main

import (
	"fmt"
	timertask "github.com/harshvladha/gotimertask"
	"time"
)

func main() {
	task := timertask.NewTask(func() {
		fmt.Println("Saying Hello, Go?")
	})
	scheduler := timertask.Schedule(task, 1*time.Second)
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Killing timertask")
	scheduler.Stop()
	time.Sleep(5 * time.Second)
}
