package main

import (
	"fmt"

	"github.com/rtb-12/Lib-Management-System/pkg/api"
)

func main() {
	fmt.Println("Started the API server")
	api.Start()
}
