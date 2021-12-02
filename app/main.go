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
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
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

// @title 校车通`
// @version 1.0`
// @description 校车通接口文档`
// @termsOfService [http://swagger.io/terms/](http://swagger.io/terms/)`
// @contact.name xuefeng`
// @contact.url [http://www.swagger.io/support](http://www.swagger.io/support)`
// @contact.email xuefeng6329@126.com`
// @license.url [http://www.apache.org/licenses/LICENSE-2.0.html](http://www.apache.org/licenses/LICENSE-2.0.html)`
// @host 47.241.77.253:10000`
// @BasePath /api/v1/`
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
