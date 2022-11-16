package main

import (
	"context"
	"getblock/internal"
	"getblock/pkg/client"
	"getblock/pkg/handler"
	"getblock/pkg/service"
	"getblock/router"
	"github.com/caarlos0/env"
	"github.com/ethereum/go-ethereum/rpc"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"os/signal"
	"syscall"
)

func main() {
	var err error

	mainCtx, cancel := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	yamlFile, err := ioutil.ReadFile("C:/Users/User/GoProject/src/getblock/resources/config/config.yaml")
	if err != nil {
		log.Fatalf("couldn't read config file err: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &internal.Configuration)
	if err != nil {
		log.Fatalf("couldn't parse config err: %v", err)
	}

	err = env.Parse(&internal.Configuration.Integration.HTTP.Block)
	if err != nil {
		log.Fatalf("couldn't parse env file err: %v", err)
	}

	gbClient, err := rpc.DialContext(mainCtx, "https://eth.getblock.io/mainnet")

	gbClient.SetHeader("x-api-key", internal.Configuration.Integration.HTTP.Block.APIKey)
	gbClient.SetHeader("Content-Type", internal.Configuration.Integration.HTTP.Block.ContentType)

	getBlockClient := client.NewGetBlockClient(gbClient, internal.Configuration.Integration.HTTP.Block.URL)
	getBlockSvc := service.NewGetBlockService(getBlockClient)

	handlerCtx := handler.NewHandlerCtx(mainCtx, handler.WithGetBlockContext(getBlockSvc))
	httpServer := &http.Server{Addr: internal.Configuration.Server.HTTP.Port, Handler: router.Router(handlerCtx)}

	log.Printf("http server is listening at %v", httpServer.Addr)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	if err = httpServer.ListenAndServe(); err != nil {
		log.Fatalf("cannot listen and serve http server: %v", err)
	}
}