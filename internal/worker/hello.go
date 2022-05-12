package worker

import (
	"time"
)

func HelloWork(worker *Worker) {
	worker.h.Infof("hello work !!!!")
	time.Sleep(10 * time.Second)
}
