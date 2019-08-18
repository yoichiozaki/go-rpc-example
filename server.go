package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type ToDo struct {
	Title string
	Status string
}

type UpdateToDo struct {
	Title string
	NewTitle string
	NewStatus string
}

type Task int

var ToDoList []ToDo

func (t *Task) GetToDoWithTitle(title string, reply *ToDo) error {
	var found ToDo
	for _, t := range ToDoList {
		if t.Title == title {
			found = t
			break
		}
	}
	*reply = found
	return nil
}

func (t *Task) MakeToDoWithTitle(title string, reply *ToDo) error {
	newToDo := ToDo{Title: title, Status: "Just launched"}
	ToDoList = append(ToDoList, newToDo)
	*reply = newToDo
	return nil
}

func (t *Task) UpdateToDo(todo UpdateToDo, reply *ToDo) error {
	var updated ToDo
	for i, t := range ToDoList {
		if t.Title == todo.Title {
			ToDoList[i].Title = todo.NewTitle
			ToDoList[i].Status = todo.NewStatus
			updated = ToDoList[i]
			break
		}
	}
	*reply = updated
	return nil
}

func (t *Task) DeleteToDo(todo ToDo, reply *ToDo) error {
	var deleted ToDo
	for i, t := range ToDoList {
		if t.Title == todo.Title && t.Status == todo.Status {
			ToDoList = append(ToDoList[:i], ToDoList[i+1:]...)
			deleted = todo
			break
		}
	}
	*reply = deleted
	return nil
}

func main() {
	task := new(Task)

	// Publish the receivers methods
	err:= rpc.Register(task)
	if err != nil {
		log.Fatalf("error: format of service `Task` does not meet `net/rpc` criteria.\n")
	}

	// Register a HTTP handler
	rpc.HandleHTTP()

	// Listen to TCP connections on port 1234
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("error: something wrong with `net.Listen`\n")
	}
	log.Printf("Servering RPC server on port `:%d`\n", 1234)

	// Start accepting incoming HTTP connections
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatalf("error: something wrong with `http.Server`\n")
	}
}