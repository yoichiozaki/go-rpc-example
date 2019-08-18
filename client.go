package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
)

type ToDo struct {
	Title  string
	Status string
}

type UpdateToDo struct {
	Title     string
	NewTitle  string
	NewStatus string
}

var scanner = bufio.NewScanner(os.Stdin)

func main() {

	// Create TCP connection to localhost on port 1234
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatalf("error: can not establish HTTP connection\n")
	}

	for {
		var reply ToDo
		fmt.Print(">> ")
		scanner.Scan()
		inputs := strings.Split(scanner.Text(), ":")
		if len(inputs) > 2 {
			log.Println("error: invalid input")
			continue
		}
		method, title := inputs[0], inputs[1]
		switch method {
		case "get":
			err = client.Call("Task.GetToDoWithTitle", title, &reply)
			if err != nil {
				log.Fatalf("error: RPC `GetToDoWithTitle` failed\n")
			}
			fmt.Printf("-> Get: %+v\n\n", reply)
		case "delete":
			err = client.Call("Task.DeleteToDoWithTitle", title, &reply)
			if err != nil {
				log.Fatalf("error: RPC `DeleteToDoWithTitle` failed\n")
			}
			fmt.Printf("-> Delete: %+v\n\n", reply)
		case "update":
			err = client.Call("Task.GetToDoWithTitle", title, &reply)
			if err != nil {
				log.Fatalf("error: RPC `GetToDoWithTitle` failed\n")
			}
			fmt.Println("Please input updated title and status like <new title>:<new status>")
			fmt.Printf("[Current] Title: `%s`, Status: `%s`\n", reply.Title, reply.Status)
			fmt.Print("(update)>> ")
			scanner.Scan()
			args := strings.Split(scanner.Text(), ":")
			if len(args) > 2 {
				log.Println("error: invalid input")
				continue
			}
			update := UpdateToDo{Title: title, NewTitle: args[0], NewStatus: args[1]}
			err = client.Call("Task.UpdateToDo", update, &reply)
			if err != nil {
				log.Fatalf("error: RPC `UpdateToDo` failed\n")
			}
			fmt.Printf("-> Update: %+v\n\n", reply)
		case "create":
			err = client.Call("Task.MakeToDoWithTitle", title, &reply)
			if err != nil {
				log.Fatalf("error: RPC `MakeToDoWithTitle` failed\n")
			}
			fmt.Printf("-> Create: %+v\n\n", reply)
		case "help":
			fmt.Println("Please input what you want to do w.r.t todo items. Exported methods are below.")
			fmt.Println("- `get:<title>`	=> Get todo item with given title and print it.")
			fmt.Println("- `update:<title>`	=> Update todo item's title and status.")
			fmt.Println("- `delete:<title>`	=> Delete todo item with given title.")
			fmt.Println("- `create:<title>`	=> Create new todo item with given title.")
			fmt.Println("- `help`			=> Show this message.")
		default:
			log.Println("error: invalid input")
		}
	}
}
