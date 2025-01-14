package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"

	"github.com/machinefi/sprout/cmd/enode/api"
	"github.com/machinefi/sprout/coordinator"
)

func main() {
	initLogger()
	bindEnvConfig()

	coordinator, err := coordinator.NewCoordinator(viper.GetString(DatabaseDSN), viper.GetString(BootNodeMultiaddr), viper.GetInt(IotexChainID))
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := api.NewHttpServer(coordinator).Run(viper.GetString(HttpServiceEndpoint)); err != nil {
			log.Fatal(err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
