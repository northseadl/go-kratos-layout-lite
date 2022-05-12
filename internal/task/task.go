package task

import (
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go-kratos-layout-lite/internal/biz"
)

// ProviderSet is task providers.
var ProviderSet = wire.NewSet(NewTasker)

type Tasker struct {
	ucg     *biz.GreeterUC
	h       *log.Helper
	taskMap map[string]ITaskFunc
}

type ITaskFunc func(params []string, tasker *Tasker) error

func NewTasker(logger log.Logger, ucg *biz.GreeterUC) *Tasker {
	tasker := Tasker{
		ucg: ucg,
		h:   log.NewHelper(logger),
	}
	tasker.registerTask("hello", helloTask)
	return &tasker
}

func (t *Tasker) registerTask(taskName string, task ITaskFunc) {
	if t.taskMap == nil {
		t.taskMap = make(map[string]ITaskFunc, 0)
	}
	t.taskMap[taskName] = task
}

// Run 根据flag执行指令, 错误需要暴露出来
func (t *Tasker) Run(taskName string, params []string) {
	if taskFunc, ok := t.taskMap[taskName]; ok {
		t.h.Infof("ready to execute the command: %s, params %p", taskName, params)
		if err := taskFunc(params, t); err != nil {
			panic(err)
		}
	} else {
		panic(errors.New("not found command"))
	}
	t.h.Infof("task %s execution completed", taskName)
}
