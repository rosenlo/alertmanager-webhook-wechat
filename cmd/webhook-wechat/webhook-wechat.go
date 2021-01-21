package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/rosenlo/alertmanager-webhook-wechat/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Version   = "Unkown"
	GitHash   = "Unkown"
	BuildTime = "Unkown"
)

// root represents the base command when called without any subcommands
var root = &cobra.Command{
	Use:     "webhook-wechat",
	Short:   "Alertmanager Webhook for Wechat",
	Long:    `Alertmanager Webhook for Wechat`,
	Version: fmt.Sprintf("%s-%s-%s", Version, GitHash, BuildTime),
	Run:     run,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return root.Execute()
}

func init() {
	log.SetFlags(log.Lmicroseconds | log.Ltime | log.Lshortfile)
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetDefault("address", ":9090")
	viper.SetDefault("robot_url", "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=$key")

	viper.AutomaticEnv() // read in environment variables that match

}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}

func run(cmd *cobra.Command, args []string) {
	log.Printf("args: %v", args)
	// 1.Init Web Service
	routes := api.New().WebService()
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
			log.Printf("graceful exit.")
			os.Exit(0)
		}
	}
}
