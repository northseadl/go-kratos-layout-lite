package worker

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go-kratos-layout-lite/internal/biz"
)

// ProviderSet is work providers.
var ProviderSet = wire.NewSet(NewWorker)

var works []Work

type Worker struct {
	h  *log.Helper
	uc *biz.GreeterUC
}

type WorkFunc func(worker *Worker)

func (w *Worker) registerWork(name string, iworkFunc WorkFunc) {
	works = append(works, Work{
		Name:     name,
		WorkFunc: iworkFunc,
	})
}

func (w *Worker) Run(done chan struct{}) {
	wc := len(works)
	cChan := make(chan struct{}, wc)
	for _, work := range works {
		go func(wer *Worker, w Work) {
			w.WorkFunc(wer)
			cChan <- struct{}{}
		}(w, work)
	}
	for i := 0; i < wc; i++ {
		<-cChan
	}
	done <- struct{}{}
}

type Work struct {
	Name     string
	WorkFunc WorkFunc
}

func NewWorker(logger log.Logger, uc *biz.GreeterUC) *Worker {
	worker := Worker{
		h:  log.NewHelper(logger),
		uc: uc,
	}
	worker.registerWork("hello", HelloWork)
	return &worker
}
