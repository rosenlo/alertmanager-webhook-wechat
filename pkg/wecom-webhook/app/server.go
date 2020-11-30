package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/rosenlo/wecom-webhook/pkg/wecom-webhook/service"

	"github.com/rosenlo/toolkits/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	AppID = "wecom-webhook"
)

func Run(cmd *cobra.Command, args []string) {
	// 0.Init logging
	log.Init(viper.GetString("log_level"), AppID, nil)

	log.Infof("args: %v", args)
	// 1.Init Web Service
	routes := service.NewService().WebService()
	go routes.Run(viper.GetString("address"))

	// 2.Catch the signal and then shut down the processes
	sigs := make(chan os.Signal, 1)
	log.Printf("pid: %d register signal notify", os.Getpid())
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for {
		signal := <-sigs
		log.Printf("receive %s signal", signal)
		switch signal {
		case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			log.Info("graceful exit.")
			os.Exit(0)
		}
	}
}
