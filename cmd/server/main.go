package main

import "src/internal/routers"

func main() {
	r := routers.NewRouter()

	r.Run()
}
