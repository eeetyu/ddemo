package main

import (
	api "gotest/kitex_gen/api/gotest"
	"log"
)

func main() {

	DbClient.InitDb()
	svr := api.NewServer(new(GotestImpl))
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
