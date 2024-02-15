package timer

import (
	"fmt"
	"time"
)

func StartPtmTimer() {
	timer := time.NewTimer(1 * time.Minute)
	defer timer.Stop()
	for {
		timer.Reset(1 * time.Minute) // 这里复用了 timer
		select {
		case <-timer.C:
			task()
		}
	}
}

func task() (int, error) {
	return fmt.Println("timer task:hardware static info .......")
}
