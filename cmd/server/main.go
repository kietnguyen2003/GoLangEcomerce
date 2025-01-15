package main

import "src/internal/initialize"

func main() {
	// r := routers.NewRouter()
	// // InitMySQL()
	// // InitRedis()
	// // InitKafka()
	// // InitRabbitMQ()
	// r.Run()

	initialize.Run()
}
