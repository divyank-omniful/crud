package main

import (
	"taskOne/mongodatabase"
	"taskOne/router"
)

func main() {
	mongodatabase.Init()
	router.Routes()
}
