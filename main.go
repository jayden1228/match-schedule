package main

import (
	"match-schedule/app"
	"log"
	"os"
	"os/signal"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	
	app.RunServer(":8000")

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
}
