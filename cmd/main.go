package main

import (
	"fmt"
	"log"

	"github.com/jestebanrods/dummy-golang/internal"
)

func main() {
	list := []internal.Todo{}

	task1, err := internal.NewTodo("HOLA MUNDO", "jossie", 11)
	if err != nil {
		log.Fatal(err)
	}

	list = append(list, task1)

	fmt.Println(task1.Message())
	fmt.Printf("%#v\n", list)
}
