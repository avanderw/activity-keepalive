package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

const timeFmt = "2006-01-02 15:04:05"

var timeout = 2 * time.Minute
var lastActivity = time.Now()

func main() {
	if len(os.Args) > 1 {
		timeout, _ = time.ParseDuration(os.Args[1])
	}
	fmt.Println(time.Now().Format(timeFmt), "timeout set to", timeout)
	go monitorInput()
	monitorIdle()
}

func monitorInput() {
	chanHook := hook.Start()
	defer hook.End()

	for _ = range chanHook {
		lastActivity = time.Now()
	}
}

func monitorIdle() {
	for {
		if time.Since(lastActivity) > timeout {
			fmt.Println(time.Now().Format(timeFmt), "idle event detected")
			moveMouse()
		}
		time.Sleep(1 * time.Second)
	}
}

func moveMouse() {
	mx, my := robotgo.Location()
	x := rand.Intn(3) - 2
	y := rand.Intn(3) - 2
	robotgo.Move(mx+x, my+y)
}
