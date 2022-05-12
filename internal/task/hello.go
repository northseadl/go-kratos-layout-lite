package task

func helloTask(params []string, tasker *Tasker) error {
	tasker.h.Infof("hello %v", params)
	return nil
}
