package controllers

import (
	"TaskManager/structs"
	"context"
	"fmt"
	"time"
)

type TasksData struct {
	tasks     map[string]*structs.Task
	cancelMap map[string]context.CancelFunc
}

type ImplementsTask interface {
	AddTask()
	ManageTask(*structs.Task, context.Context)
	DeleteTask()
	ListTasks()
}

func GetNewTasks() *TasksData {
	return &TasksData{
		tasks:     make(map[string]*structs.Task),
		cancelMap: make(map[string]context.CancelFunc),
	}
}

func (t *TasksData) AddTask() {
	var title string
	var time uint16
	fmt.Print("Ingrese el titulo de la tarea: ")
	fmt.Scanf("%s", &title)
	fmt.Print("Ingrese el tiempo de la tarea: ")
	fmt.Scanf("%d", &time)

	ctx, cancel := context.WithCancel(context.Background())

	task := structs.Task{
		Title: title,
		Time:  time,
		State: "R",
	}
	t.cancelMap[task.Title] = cancel
	t.tasks[task.Title] = &task
	go t.ManageTask(&task, ctx)

}

func (t *TasksData) ManageTask(task *structs.Task, ctx context.Context) {
	select {
	case <-time.After(time.Duration(task.Time) * time.Second):
		task.State = "C"
		fmt.Printf("\nTask %s completed", task.Title)
	case <-ctx.Done():
		fmt.Println("Task state: " + task.State)
		if task.State == "E" {
			fmt.Printf("\nTask %s canceled", task.Title)
		} else if task.State == "R" {
			fmt.Printf("\nTask %s completed", task.Title)
		}
	}
}

func (t *TasksData) DeleteTask() {
	var title string
	fmt.Print("\nIngrese el nombre de la tarea que desea eliminar: ")
	fmt.Scanf("%s", &title)
	task, ok := t.tasks[title]
	if ok {
		if task.State == "R" {
			task.State = "E"
			cancelFn := t.cancelMap[title]
			cancelFn()
		}
	} else {
		fmt.Println("Task not found")
	}
}

func (t *TasksData) ListTasks() {
	for _, task := range t.tasks {
		fmt.Printf("Task %s is %s \n", task.Title, task.State)
	}
}
