package main

import (
	"example/server"
)

func main() {
	srv := server.DefaultServer
	srv.ListenAndServe()
}
