package main

import (
	service "github.com/liankui/prenatal-server/question/cmd/service"
	"github.com/liankui/prenatal-server/question/pkg/db"
)

func main() {
	err := db.Init()
	if err != nil {
		panic(err)
	}
	service.Run()
}
