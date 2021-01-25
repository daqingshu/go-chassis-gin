package main

import (
	"github.com/go-chassis/go-chassis/v2"
	"github.com/go-chassis/openlog"

	_ "github.com/daqingshu/go-chassis-gin"
)

func main() {
	chassis.RegisterSchema("gin", &RestFulHello{})
	if err := chassis.Init(); err != nil {
		openlog.Fatal("Init failed." + err.Error())
		return
	}
	chassis.Run()
}
