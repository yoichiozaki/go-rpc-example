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
	var reply ToDo

	// Create TCP connection to localhost on port 1234
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatalf("error: can not establish HTTP connection\n")
	}

	for {
		fmt.Println("Please input what you want to do w.r.t todo items. Exported methods are below.")
		fmt.Println("\t- `Get <title>`		=> Get todo item with given title and print it.")
		fmt.Println("\t- `Update <title>`	=> Update todo item's title and status.")
		fmt.Println("\t- `Delete <title>`	=> Delete todo item with given title.")
		fmt.Println("\t- `Make <title>`		=> Make new todo item with given title.")
		fmt.Print(">> ")
		scanner.Scan()
		inputs := strings.Split(scanner.Text(), " ")
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
			log.Println("get: #v", reply)
		case "delete":
			err = client.Call("Task.DeleteToDoWithTitle", title, &reply)
			if err != nil {
				log.Fatalf("error: RPC `DeleteToDoWithTitle` failed\n")
			}
			log.Println("delete: #v", reply)
		case "make":
			err = client.Call("Task.GetToDoWithTitle", title, &reply)
			if err != nil {
				log.Fatalf("error: RPC `GetToDoWithTitle` failed\n")
			}
			fmt.Printf("[Current] Title: `%s`, Status: `%s`\n", reply.Title, reply.Status)
			fmt.Println("Please input updated title and status like <new title>:<new status>")
			fmt.Print(">> ")
			scanner.Scan()
			args := strings.Split(scanner.Text(), ":")
			if len(args) > 2 {
				log.Println("error: invalid input")
				continue
			}
			update := UpdateToDo{Title: title, NewTitle: args[0], NewStatus: args[1]}
			err = client.Call("Task.MakeToDoWithTitle", update, &reply)
			if err != nil {
				log.Fatalf("error: RPC `MakeToDoWithTitle` failed\n")
			}
			log.Println("make: #v", reply)
		case "update":
			err = client.Call("Task.UpdateToDo", title, &reply)
			if err != nil {
				log.Fatalf("error: RPC `UpdateToDo` failed\n")
			}
			log.Println("reply: #v", reply)
		default:
			log.Println("error: invalid input")
		}
	}
}
