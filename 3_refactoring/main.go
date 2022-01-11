package main

import (
	"fmt"
	"time"
)

type Activity interface {
	Title() string
	Do()
}

type Work struct{}

func (w Work) Title() string {
	return "work"
}

func (w Work) Do() {}

type Leisure struct{}

func (l Leisure) Title() string {
	return "leisure"
}

func (l Leisure) Do() {}

func main() {
	activity := FindSomethingToDo()
	fmt.Printf("doing: %s\n", activity.Title())
	activity.Do()
}

func FindSomethingToDo() Activity {
	if time.Now().Weekday() == time.Sunday {
		return Leisure{}
	}
	return Work{}
}
