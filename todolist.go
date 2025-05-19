package main

import "fmt"
type TodoList struct {
	tasks []Task	
}

func (td *TodoList) AddTask(taskName string) {
	taskId := len(td.tasks) + 1
	task := NewTask(taskId, taskName) 
    td.tasks = append(td.tasks, task)

	fmt.Println(taskName," Task added successfully!")
}


 func (td *TodoList) ViewTasks() {
	
	fmt.Println("==============Tasks are================")

	for _, taks := range td.tasks{
		doneStatus := " "

		if taks.Done {
			doneStatus = "x"
		} 
		fmt.Printf("[%s] Task: #%d: %s\n", doneStatus, taks.ID, taks.Name)
	}
	fmt.Println("=======================================")

}

func (td *TodoList) MarkTaskAsDone(taskId int) {

	if taskId < 1 || taskId > len(td.tasks) {
		fmt.Println("Invalid task ID")
		return
	}
	td.tasks[taskId-1].MarkAsDone()

	
}
