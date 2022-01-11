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

	if a, ok := activity.(Work); ok {
		fmt.Println("doing: work")
		a.Work()
	} else if a, ok := activity.(Leisure); ok {
		fmt.Println("doing: leisure")
		a.Chill()
	} else {
		log.Fatalln("doing: well...")
	}
}

func FindSomethingToDo() interface{} {
	if time.Now().Weekday() == time.Sunday {
		return Leisure{}
	}
	return Work{}
}
