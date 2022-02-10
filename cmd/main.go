package main

import (
	"flag"
	"runtime"
)

// config file
var configFile = flag.String("f", "../configs/dev.yml", "set config file which viper will loading.")

func main() {
	runtime.GOMAXPROCS(1)

	flag.Parse()

	app, err := CreateApp(*configFile)
	if err != nil {
		panic(err)
	}

	if err := app.Start(); err != nil {
		panic(err)
	}

	app.AwaitSignal()
}
