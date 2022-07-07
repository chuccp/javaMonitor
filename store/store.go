package store

import (
	"github.com/chuccp/javaMonitor/execute"
	"sync"
)

var javaMap = new(sync.Map)

var scheduleMap = new(sync.Map)

type JavaExec interface {
	GetName() string
	Start() string
	Stop() string
}

func AddJavaExec(javaExec JavaExec) {
	javaMap.Store(javaExec.GetName(), javaExec)
}
func GetJavaExec(name string) (JavaExec, bool) {
	je, ok := javaMap.Load(name)
	if ok {
		return je.(JavaExec), ok
	}
	return nil, false
}
func GetJavaExecList() []*execute.JavaExec {
	array := make([]*execute.JavaExec, 0)
	javaMap.Range(func(key, value interface{}) bool {
		val, ok := value.(*execute.JavaExec)
		if ok {
			val.Status = val.HasRun()
			array = append(array, val)
		}
		return true
	})
	return array
}

type Schedule interface {
	GetName() string
	Open() string
	Close() string
}

func AddSchedule(schedule Schedule) {
	scheduleMap.Store(schedule.GetName(), schedule)
}
func GetSchedule(name string) (Schedule, bool) {
	je, ok := scheduleMap.Load(name)
	if ok {
		return je.(Schedule), ok
	}
	return nil, false
}
func GetScheduleList() []Schedule {
	array := make([]Schedule, 0)
	scheduleMap.Range(func(key, value interface{}) bool {
		val, ok := value.(Schedule)
		if ok {
			array = append(array, val)
		}
		return true
	})
	return array
}
