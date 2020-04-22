package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Divya....")
	//	run.Hello()
	log.Info("INfo message...")
	log.Warn("Warn message....")
	log.Fatal("Fatal message....")
}
