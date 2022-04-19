package main

import "http_req/router"

func main() {
	var PORT = ":8080"

	router.StartServer().Run(PORT)
}
