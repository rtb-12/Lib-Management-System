package main

import (
	"fmt"

	"github.com/ashpect/go-mvc/pkg/api"
)

func main() {
	fmt.Println("Started the API server")
	api.Start()
}
