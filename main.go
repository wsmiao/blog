package main

import (
	"github.com/wsmiao/blog/controller"
	"github.com/wsmiao/blog/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func init(){
	service.ConnectDB()
}

func main() {
	println("server run")

	router := controller.Router()
	server := &http.Server{
		Addr:    "0.0.0.0:80",
		Handler: router,
	}

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	go func() {
		s := <-c
		log.Print("server exiting %s", s)
		if err := server.Close(); nil != err {
			log.Fatalf("server close failed: " + err.Error())
		}
		service.DisconnectDB()
		os.Exit(0)
	}()

	if err := server.ListenAndServe(); nil != err {
		log.Fatalf("listen and serve failed: " + err.Error())
	}
}