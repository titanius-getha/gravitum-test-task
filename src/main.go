package main

import (
	"github.com/titanius-getha/gravitum-test-task/app"
	"github.com/titanius-getha/gravitum-test-task/pkg/config"
)

func main() {
	conf, err := config.New()
	if err != nil {
		panic("load config error: " + err.Error())
	}

	a := app.New(conf)
	a.Start()
}
