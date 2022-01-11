package main

import (
	"fmt"
	"log"
	"time"
)

type Work struct{}

func (w Work) Work() {}

type Leisure struct{}

func (l Leisure) Chill() {}

func main() {
	activity := FindSomethingToDo()

	switch a := activity.(type) {
	case Work:
		fmt.Println("doing: work")
		a.Work()
	case Leisure:
		fmt.Println("doing: leisure")
		a.Chill()
	default:
		log.Fatalln("doing: well...")
	}
}

func FindSomethingToDo() interface{} {
	if time.Now().Weekday() == time.Sunday {
		return Leisure{}
	}
	return Work{}
}
