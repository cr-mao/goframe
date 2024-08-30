package http

import (
	"fmt"
	"goframe/global/config"
	"net/http"
	"time"

	"goframe/app/http/routers"
	"goframe/infra/console"
	"goframe/infra/helpers"
)

func NewServe() *http.Server {
	router := routers.NewRouter()
	addr := fmt.Sprintf("%s:%d", config.AppConfig.App.HttpHost, config.AppConfig.App.HttpPort)
	s := &http.Server{
		Addr:    addr,
		Handler: router,
	}
	console.Success(time.Now().Format(helpers.CSTLayout) + " http listening on " + addr)
	return s
}
