package mutexz

import (
	`log`
	`time`
)

func Walk() {
	monitor := GetInstance()
	monitor.Submit()
wait:
	for {
		select {
		case <-monitor.resultQueue:
			log.Printf("update cache successfully")
			break wait
		case <-time.After(3 * time.Second):
			log.Printf("time out after 3s.")
		}
	}
}
