package main

import (
	"fmt"
)

type ToDo struct {
	Title string
	Status Status
}

type Status string
const (
	DONE Status = "Done!"
	WIP Status = "Work in progress..."
	LAUNCHED Status = "Just launched, so let's start!"
)

var todoSlice []ToDo

func ReadToDoWithTitle(title string) (*ToDo, error) {
	for _, t := range todoSlice {
		if t.Title == title {
			return &t, nil
		}
	}
	return nil, fmt.Errorf(
		"error: no ToDo item with such title: %s", title)
}

func CreateToDoWithTitle(title string) (*ToDo, error) {
	todo := ToDo{Title:title, Status:LAUNCHED}
	todoSlice = append(todoSlice, todo)
	return &todo, nil
}

func UpdateToDoWithTitle(before string, after string) (*ToDo, error) {
	var updated ToDo
	for i, t := range todoSlice {
		if t.Title == before {
			todoSlice[i].Title = after
			updated = todoSlice[i]
			return &updated, nil
		}
	}
	return nil, fmt.Errorf("error: no ToDo item with such title: %s", before)
}

func UpdateToDoWithStatus(title string, status Status) (*ToDo, error) {
	var updated ToDo
	for i, t := range todoSlice {
		if t.Title == title {
			todoSlice[i].Status = status
			updated = todoSlice[i]
			return &updated, nil
		}
	}
	return nil, fmt.Errorf("error: no ToDo item with such title: %s", title)
}

func DeleteToDoWithTitle(title string) (*ToDo, error) {
	var deleted ToDo
	for i, t := range todoSlice {
		if t.Title == title {
			switch t.Status {
			case DONE:
				todoSlice = append(todoSlice[:i], todoSlice[i+1:]...)
				deleted = t
				return &deleted, nil
			case WIP, LAUNCHED:
				return nil, fmt.Errorf("error: ToDo item with title `%s` is not completed yet", title)
			default:
				return nil, fmt.Errorf("error: invalid status found: %+v", t)
				}
		}
	}
	return nil, fmt.Errorf("error: no ToDo item with such title: %s", title)
}

func main() {

}