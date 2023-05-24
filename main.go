package main

import (
	"fmt"
	"web-service/pkg/api"
	"web-service/pkg/db"
)

func main() {
	fmt.Println("start initialization web-service")
	db.InitRepo()
	fmt.Println("db connection initialized")
	api.InitApi()
}
