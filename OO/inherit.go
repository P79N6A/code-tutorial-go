package main

import "fmt"

type ITask interface {
	Inject() string
}
type Task struct {
	name string
}
func (t *Task)Inject() string {
	return "Task"
}
func (t *Task)Run(i ITask) {
	x := i.Inject()
	println("in Task Run  " + x)
}
type MyTask struct {
	*Task
}
func (t* MyTask)Inject() string {
	return "Mytask" + t.name
}

func main() {
	t := &MyTask{
		Task:&Task{"xxx"},
	}

	fmt.Println(t.Inject())
}
