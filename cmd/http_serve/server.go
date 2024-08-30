package http_serve

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"

	httpServe "goframe/app/http"
	"goframe/bootstrap"
	"goframe/cmd"
	"goframe/infra/logger"
)

var HttpServeCmd = &cobra.Command{
	Use: "http_serve",
	Run: func(cmd *cobra.Command, args []string) {
		if err := Run(cmd.Context()); err != nil {
			logger.Errorf("cmd http serve run error: %v", err)
			os.Exit(1)
		}
	},
	// rootCmd 的所有子命令都会执行以下代码
	PersistentPreRun: func(command *cobra.Command, args []string) {
		bootstrap.HttpServerBootstrap(cmd.Env)
	},
}

func Run(ctx context.Context) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//http start
	srv := httpServe.NewServe()
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			logger.Errorf("http ListenAndServe error: %v", err)
		}
	}()
	<-quit
	// shutdown http server
	newCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(newCtx); err != nil {
		logger.Errorf("http server gracefully shutdown err :%v", err)
	}
	return nil
}
