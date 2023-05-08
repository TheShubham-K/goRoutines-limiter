package main

import (
	"multithreading2/pkg/events"
	"time"
)

func main() {

	manager := events.CreateManager(time.Second / 300)
	processor := events.CreateProcessor()

	limiter := make(chan int, 10)
	for event := range manager.Stream() {
		limiter <- 1
		go func(e events.Event) {
			processor.ProcessEvent(e)
			<-limiter
		}(event)
	}
}
