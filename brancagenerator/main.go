package main

import "projects/applets/brancagenerator/service"

func main() {

	server := service.NewServer()
	// server.Run(":" + os.Getenv("PORT"))
	server.Run(":3001")
}
