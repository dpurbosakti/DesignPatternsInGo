package main

import (
	"fmt"
	"sync"
)

// commonly used for DB instance, log, configuration

type config struct {
	DBName string
}

var (
	mutex = &sync.Mutex{}
	cfg   *config
)

func getConfig() *config {
	if cfg == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if cfg == nil {
			fmt.Println("creating config...")
			cfg = &config{
				DBName: "test",
			}
		} else {
			fmt.Println("config already exist... code 1")
		}
	} else {
		fmt.Println("config already exist... code 1")
	}
	return cfg
}

func main() {
	for i := 0; i < 100; i++ {
		go getConfig()
	}
	fmt.Scanln()
}
