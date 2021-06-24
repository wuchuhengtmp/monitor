/**
 * @Desc    The server is part of http-api
 * @Author  wuchuheng<root@wuchuheng.com>
 * @Blog    https://wuchuheng.com
 * @wechat  wc20030318
 * @DATE    2021/6/23
 * @Listen  MIT
 */
package server

import (
	"bytes"
	"github.com/urfave/cli"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

/**
 * 启动监测服务
 */
func RunServer(c *cli.Context) {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <- ticker.C:
			getCPUSample()
		}
	}
}

type Process struct {
	pid int
	cpu float64
}

/**
 * 获取cpu 样本
 */
func getCPUSample() {
	cmd := exec.Command("ps", "aux")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	processes := make([]*Process, 0)
	for {
		line, err := out.ReadString('\n')
		if err!=nil {
			break;
		}
		tokens := strings.Split(line, " ")
		ft := make([]string, 0)
		for _, t := range(tokens) {
			if t!="" && t!="\t" {
				ft = append(ft, t)
			}
		}
		log.Println(len(ft), ft)
		pid, err := strconv.Atoi(ft[1])
		if err!=nil {
			continue
		}
		cpu, err := strconv.ParseFloat(ft[2], 64)
		if err!=nil {
			log.Fatal(err)
		}
		processes = append(processes, &Process{pid, cpu})
	}
	for _, p := range(processes) {
		log.Println("Process ", p.pid, " takes ", p.cpu, " % of the CPU")
	}
}

