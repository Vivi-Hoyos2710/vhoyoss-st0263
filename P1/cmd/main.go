package main

import (
	"fmt"
	"github.com/Vivi-Hoyos2710/vhoyoss-st0263/P1/config"
)

func main() {
	config, err := config.Bootstrap("./config/.env")
	if err != nil {
		panic(err)
	}
	fmt.Println(config)
	err = runClient()
	if err != nil {
		panic(err)
	}
	err = runServer()
	if err != nil {
		panic(err)
	}
}

func runClient() error {
	panic("oo")
}
func runServer() error {
	panic("aaa")
}
