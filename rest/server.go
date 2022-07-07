package rest

import (
	"gitee.com/cooge/javaMonitor/config"
	"gitee.com/cooge/javaMonitor/store"
	"github.com/pquerna/ffjson/ffjson"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func apiExecList(w http.ResponseWriter, re *http.Request) {
	list := store.GetJavaExecList()
	data, err := ffjson.Marshal(list)
	if err == nil {
		w.Write(data)
	} else {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}
}
func apiExecStop(w http.ResponseWriter, re *http.Request) {
	name := re.FormValue("name")
	if len(name) > 0 {
		javaExec, ok := store.GetJavaExec(name)
		if ok {
			ok := javaExec.Stop()
			w.Write([]byte(ok))
			return
		}
	}
	w.Write([]byte(config.FAIL))
}
func apiExecStart(w http.ResponseWriter, re *http.Request) {
	name := re.FormValue("name")
	if len(name) > 0 {
		javaExec, ok := store.GetJavaExec(name)
		if ok {
			ok := javaExec.Start()
			w.Write([]byte(ok))
			return
		}
	}
	w.Write([]byte(config.FAIL))
}

func apiScheduleList(w http.ResponseWriter, re *http.Request) {
	list := store.GetScheduleList()
	data, err := ffjson.Marshal(list)
	if err == nil {
		w.Write(data)
	} else {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}
}
func apiScheduleStop(w http.ResponseWriter, re *http.Request) {
	name := re.FormValue("name")
	if len(name) > 0 {
		schedule, ok := store.GetSchedule(name)
		if ok {
			ok := schedule.Close()
			w.Write([]byte(ok))
			return
		}
	}
	w.Write([]byte(config.FAIL))
}
func apiScheduleStart(w http.ResponseWriter, re *http.Request) {
	name := re.FormValue("name")
	if len(name) > 0 {
		schedule, ok := store.GetSchedule(name)
		if ok {
			ok :=schedule.Open()
			w.Write([]byte(ok))
			return
		}
	}
	w.Write([]byte(config.FAIL))
}

func Start() {
	//workPath, _ := os.Getwd()
	//workPath = strings.ReplaceAll(workPath, "\\", "/")
	//webPath := workPath + "/web"
	//log.Println("web目录：", webPath)
	//http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(webPath))))
	http.HandleFunc("/api/exec/list", apiExecList)
	http.HandleFunc("/api/exec/stop", apiExecStop)
	http.HandleFunc("/api/exec/start", apiExecStart)

	http.HandleFunc("/api/schedule/list", apiScheduleList)
	http.HandleFunc("/api/schedule/stop", apiScheduleStop)
	http.HandleFunc("/api/schedule/start", apiScheduleStart)

	httpConfig := config.GetHttpConfig()

	log.Println("启动端口号：", httpConfig.Port)
	server := &http.Server{
		Addr: ":" + httpConfig.Port,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Panic("启动失败", err)
	}
}
