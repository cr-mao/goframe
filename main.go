package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"goframe/cmd"
	"goframe/cmd/http_serve"
)

var (
	// build 时间
	BuildTime string
	// git commit
	Version string
)

var rootCmd = &cobra.Command{
	Use:   "goframe",
	Short: "this is web project",
}

// 注入git版本和编译时间
func init() {
	log.Printf("Version: %s", Version)
	log.Printf("BuildTime: %s", BuildTime)
}

func main() {
	// http web 服务
	rootCmd.AddCommand(http_serve.HttpServeCmd)

	// 注册全局参数，--env  是什么环境
	cmd.RegisterGlobalFlags(rootCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
