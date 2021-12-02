package main

import (
	"flag"
	"github.com/lfxnxf/school/api-gateway/ws"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	"github.com/lfxnxf/frame/logic/inits"
	"github.com/lfxnxf/school/api-gateway/conf"
	"github.com/lfxnxf/school/api-gateway/server/http"
	"github.com/lfxnxf/school/api-gateway/service"
)

func init() {
	configS := flag.String("config", "config/config.toml", "Configuration file")
	appS := flag.String("app", "", "App dir")
	flag.Parse()

	inits.Init(
		inits.ConfigPath(*configS),
	)

	if *appS != "" {
		inits.InitNamespace(*appS)
	}

}

func main() {

	defer inits.Shutdown()

	// init local config
	cfg, err := conf.Init()
	if err != nil {
		logging.Fatalf("service config init error %s", err)
	}

	// create a service instance
	srv := service.New(cfg)

	wsSrv := ws.New(cfg)

	go ws.Manager.Start()

	// init and start http server
	http.Init(srv, wsSrv, cfg)

	defer http.Shutdown()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-sigChan
		log.Printf("get a signal %s\n", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Println("api-gateway server exit now...")
			return
		case syscall.SIGHUP:
		default:
		}
	}
}
