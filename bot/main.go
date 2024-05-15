package main

import (
	"github.com/fairytale5571/awesomeProject1/config"
	"github.com/fairytale5571/awesomeProject1/internal/app"
)

func main() {
	if err := config.ParseConfig(); err != nil {
		panic(err)
	}
	instanceApp, err := app.New()
	if err != nil {
		panic(err)
	}
	if err := instanceApp.Run(); err != nil {
		panic(err)
	}

	select {}
}
