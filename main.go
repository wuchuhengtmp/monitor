package main

import (
	"github.com/urfave/cli"
	"log"
	"monitor/server"
	"monitor/tests"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "tests-监测程序"
	app.Usage = "用于监测cup限定的阈值以及可疑的启动命令"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		cli.Command{
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "commandSample",
					Usage: "可疑进程启动行样本",
				},
				&cli.StringFlag{
					Name:  "touchCpu",
					Usage: "cpu触发值",
					Required: true,
				},
				&cli.StringFlag{
					Name:  "email",
					Usage: "接收邮箱",
					Required: true,
				},

			},

			Action: server.RunServer,
			Usage:  "启动监测程序",
			Name:   "start",
		},
		cli.Command{
			Action: tests.RunTest,
			Name:   "tests",
			Usage:  "启动测试",
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Println(err)
	}
}

