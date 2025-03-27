package main

import (
	"fmt"
	"task-management/routes"
)

func main(){
	fmt.Println("Starting server on port 8080")
	routes.RegisterRoutes();
}