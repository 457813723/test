package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"github.com/urfave/cli/v2"
	"os"
	"sync"
	"time"
)

var (
	ctxServer, cancel = context.WithCancel(context.Background())
	wgServer          = sync.WaitGroup{}
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Load configuration from `FILE`",
			},
			&cli.StringFlag{
				Name:    "mode",
				Aliases: []string{"m"},
				Usage:   "Server Type, ``(default): api and timer server, `api`: api server, `timer`: timer server",
			},
		},
		Action: runServer,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
	wgServer.Wait()
}

func runServer(ctx *cli.Context) error {
	configFilePath := ctx.String("config")
	fmt.Println("config path: ", configFilePath)
	//service mode
	mode := ctx.String("mode")
	if mode == "api" {
		fmt.Println("init api server")
		initApiServer()
	} else if mode == "timer" {
		fmt.Println("init timer server")
		initTimerServer()
	} else {
		fmt.Println("init api server and timer server")
		initTimerServer()
		initApiServer()

	}
	return nil
}
func initApiServer() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, "hello world v11.0.0")
	})
	r.Run("0.0.0.0:8899") // 监听并在 0.0.0.0:8899 上启动服务
}
func initTimerServer() {
	ticker := time.NewTicker(time.Second * 2)
	wgServer.Add(1)
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("timer start ...")
			case <-ctxServer.Done():
				fmt.Println("timer done")
				wgServer.Done()
				return
			}
		}
	}()
}

//CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
//docker build -f ./docker/Dockerfile -t go-test:v13 .
//111111
