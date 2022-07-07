package main

import (
	"github.com/chuccp/javaMonitor/config"
	"github.com/chuccp/javaMonitor/execute"
	"github.com/chuccp/javaMonitor/rest"
	"github.com/chuccp/javaMonitor/schedule"
	"github.com/chuccp/javaMonitor/store"
	log "github.com/sirupsen/logrus"
	"os"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Println("当前进程", os.Getpid())
	config.Init()
	bc := config.GetJavaConfig()
	if bc.Start {
		jc := bc.JarConfigs
		for _, v := range jc {
			javaExec := execute.NewJavaExec(v)
			store.AddJavaExec(javaExec)
			javaExec.Run()
		}
	}
	sc := config.GetScheduleConfig()
	log.Println(sc.Start,"===================================")
	if sc.Start {
		jc := sc.JobConfigs
		for _, v := range jc {
			sc:=schedule.NewSchedule(v)
			store.AddSchedule(sc)
			sc.Load()
		}
	}
	rest.Start()
}
