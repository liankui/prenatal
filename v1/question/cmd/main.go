package main

import (
	service "liankui/prenatal-server/v1/question/cmd/service"
	"liankui/prenatal-server/v1/question/pkg/db"
)

func main() {
	err := db.Init()
	if err != nil {
		panic(err)
	}
	service.Run()
}
