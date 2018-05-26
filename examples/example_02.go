package main

import (
	"github.com/harshvladha/timertask"
	"time"
	"fmt"
)

func main() {
	timertask.Schedule(timertask.NewTask(func() {
		fmt.Println(time.Now().UTC())
	}), 2 * time.Second)

	timertask.Schedule(timertask.NewTask(testFunction), 1 * time.Second)

	time.Sleep(11 * time.Second)
}

func testFunction()  {
	fmt.Println("=====")
}
