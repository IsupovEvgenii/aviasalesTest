package main

import (
	"aviasalesTest/cmd/app"
)

func main() {
	application := app.New()

	if err := application.Init(); err != nil {
		application.Logger.Fatalf("can not init application:%v", err)
	}
}