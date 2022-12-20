package main

import (
	"fmt"
	"github.com/godemo/coredemo/app/provider/demoProvider"
	"github.com/godemo/coredemo/framework/gin"
	"log"
	"net/http"
)

func main() {
	//core := gin.NewCore()
	//registerRouter(core)
	//server := &http.Server{
	//	Handler: core,
	//	Addr:    ":8081",
	//}
	//err := server.ListenAndServe()
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	fmt.Println("hello")
	core := gin.New()
	core.Bind(&demoProvider.DemoProvider{})

	registerRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8081",
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		return
	}
}
