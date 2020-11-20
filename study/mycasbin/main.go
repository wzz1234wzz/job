package main

import (
	"mycasbin/db"
	"mycasbin/handle"
)

func main() {
	db.Init()
	handle.InitHandlers()
}




