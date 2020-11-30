package main

import (
	"runtime"

	"github.com/rosenlo/toolkits/log"
	"github.com/rosenlo/wecom-webhook/pkg/wecom-webhook/app"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if err := app.Execute(); err != nil {
		log.Fatal(err)
	}
}
