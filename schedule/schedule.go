package schedule

import (
	"github.com/chuccp/javaMonitor/config"
	"github.com/chuccp/javaMonitor/store"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"strings"
)

type Schedule struct {
	jobConfig *config.JobConfig
	Status    string
	cron      *cron.Cron
	Cron      string
	Name      string
	Exec      string
}

func NewSchedule(jobConfig *config.JobConfig) *Schedule {
	var schedule Schedule
	schedule.jobConfig = jobConfig
	schedule.Status = config.NoOpen
	schedule.Name = jobConfig.Name
	schedule.Cron = jobConfig.Cron
	schedule.Exec = jobConfig.Exec
	return &schedule
}
func (schedule *Schedule) Load() {
	log.Println("导入开始：",schedule.Name,schedule.Cron)
	schedule.cron = cron.New(cron.WithSeconds())
	_,err:=schedule.cron.AddFunc(schedule.Cron, func() {
		log.Println("开始执行任务：",schedule.Name,schedule.Cron)
		schedule.run()
	})
	if err==nil{
		if schedule.jobConfig.Run{
			schedule.Status = config.OPEN
			log.Println("准备执行任务：",schedule.Name,schedule.Cron)
			schedule.cron.Start()
		}
		log.Println("导入完成：",schedule.Name,schedule.Cron)
	}else{
		log.Println("导入任务失败：",schedule.Name,schedule.Cron,err)
	}


}
func (schedule *Schedule) run() {
	es:=strings.Split(schedule.Exec,",")
	for _,v:=range es{
		javaExec,ok:=store.GetJavaExec(v)
		if ok{

			if schedule.jobConfig.Mode==config.REPLACE{
				r:=javaExec.Stop()
				log.Println("停止任务",v,"======",r)
				r=javaExec.Start()
				log.Println("开始任务",v,"======",r)
			}else{
				r:=javaExec.Start()
				log.Println("开始任务",v,"======",r)
			}
		}
	}
}

func (schedule *Schedule) GetName() string {
	return schedule.Name
}
func (schedule *Schedule) Open() string {

	if schedule.jobConfig.Run{
		return config.AlreadyOpen
	}
	schedule.jobConfig.Run = true
	schedule.cron.Start()
	return config.SUCCESS
}
func (schedule *Schedule) Close() string {
	if !schedule.jobConfig.Run{
		return  config.NO_OPEN
	}
	schedule.Status = config.NO_OPEN
	schedule.jobConfig.Run = false
	schedule.cron.Stop()
	return config.SUCCESS
}
