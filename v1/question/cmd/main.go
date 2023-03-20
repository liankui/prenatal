package main

import (
	service "github.com/liankui/prenatal-server/v1/question/cmd/service"
	"github.com/liankui/prenatal-server/v1/question/pkg/db"
)

func main() {
	err := db.Init()
	if err != nil {
		panic(err)
	}
	service.Run()
}
