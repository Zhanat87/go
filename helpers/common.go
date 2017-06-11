package helpers

import (
	"fmt"
	"os/user"
	"os"
	"runtime"
	"time"
	"log"
)

func CurrentUser() *user.User {
	usr, err := user.Current()
	if err != nil {
		SendErrorEmail("golang balu", fmt.Sprintf("failed get current user: %s", err))
		panic(err)
	}
	return usr
}

func IsDocker() bool {
	return os.Getenv("HOME") == "/root"
}

func MonitorRuntime() {
	log.Println("Number of CPUs:", runtime.NumCPU())
	m := &runtime.MemStats{}
	for {
		r := runtime.NumGoroutine()
		log.Println("Number of goroutines", r)
		runtime.ReadMemStats(m)
		log.Println("Allocated memory", m.Alloc)
		time.Sleep(10 * time.Second)
	}
}