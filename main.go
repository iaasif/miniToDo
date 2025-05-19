package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	todoList := TodoList{}
	scanner := bufio.NewScanner(os.Stdin) 

	for{
	fmt.Println("====================todolist====================")
	fmt.Println("1. Add a task")
	fmt.Println("2. View tasks")
	fmt.Println("3. mark a task as done")
	fmt.Println("4. Exit")
	fmt.Println("================================================")
	fmt.Println("Please enter your choice (1-4):")

	
	scanner.Scan()
	choice := scanner.Text();
	// fmt.Println("scaned text:", choice)

	switch choice {
	case "1":
		fmt.Println("Enter A Task name :")
		scanner.Scan()
		
		taskName := scanner.Text()
		todoList.AddTask(taskName)
	case "2":
		// fmt.Println("Viewing tasks...")
		todoList.ViewTasks()
	case "3":

		fmt.Println("Enter a task Id :")
		scanner.Scan()
		taskIdstr := scanner.Text()
		taskId,err := strconv.Atoi(taskIdstr)
		if err != nil {
			fmt.Println("Invalid task ID")
			continue
		}
		todoList.MarkTaskAsDone(taskId) 
	case "4":
		fmt.Println("Exiting the program...")
		return
	default:
		fmt.Println("Invalid choice. Please try again.")
	}
	
}
}