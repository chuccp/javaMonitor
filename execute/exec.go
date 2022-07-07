package execute

import (
	"fmt"
	"github.com/chuccp/javaMonitor/config"
	"github.com/chuccp/javaMonitor/stream"
	"github.com/chuccp/javaMonitor/util"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
	"sync"
	"time"
)

func NewJavaExec(javaConfigs *config.JarConfig) *JavaExec {
	var javaExec JavaExec
	javaExec.JavaConfigs = javaConfigs
	javaExec.Name = javaConfigs.Name
	return &javaExec
}

type JavaExec struct {
	JavaConfigs *config.JarConfig
	Process     *os.Process
	LastRunTime string
	Status      bool
	Name        string
}

var lock sync.RWMutex

func (je *JavaExec) HasRun() bool{
	process,err:=FindJavaProcess(je.JavaConfigs.JarArgs)
	if err==nil{
		if len(process)>0{
			return true
		}
	}else {
		log.Panicln("不支持 WMI",err)
	}
	return false
}

func (je *JavaExec) Run() {
	lock.Lock()
	defer lock.Unlock()
	if je.JavaConfigs.Run {
		if !je.HasRun() {
			dir, _ := os.Getwd()
			pPath, _ := path.Split(je.JavaConfigs.JarPath)
			os.Chdir(pPath)
			defer os.Chdir(dir)
			cmd := je.command(je.JavaConfigs)
			output, err := cmd.StdoutPipe()
			if err == nil {
				err := cmd.Start()
				je.Process = cmd.Process
				if err == nil {
					go je.run(output)
				}else{
					log.Println("执行失败",err,je.JavaConfigs.JarArgs)
				}
			}else{
				log.Println("执行失败",err,je.JavaConfigs.JarArgs)
			}
		}
	}
}
func (je *JavaExec) command(javaConfigs *config.JarConfig) *exec.Cmd {
	_, file := path.Split(javaConfigs.JarPath)
	javaBinPath := javaConfigs.JavaHome + "/bin/" + javaConfigs.JavaBin
	args := []string{javaBinPath}
	if len(javaConfigs.JavaOps) > 0 {
		ops := strings.Split(javaConfigs.JavaOps, " ")
		for _, v := range ops {
			args = append(args, v)
		}
	}
	args = append(args, "-jar", file)

	if len(javaConfigs.JarArgs) > 0 {
		ops := strings.Split(javaConfigs.JarArgs, " ")
		for _, v := range ops {
			args = append(args, v)
		}
	}
	cmd := exec.Command(javaBinPath)
	cmd.Args = args
	return cmd
}

func (je *JavaExec) run(reader io.Reader) {
	now  := time.Now()
	je.LastRunTime =now.Format("2006-01-02 15:04:05")
	defer je.recordRecover()
	var stringStream = stream.NewStringStream(reader)
	for {
		data, err := stringStream.ReadLine()
		if err == nil {
			runtime.Gosched()
			fmt.Println(util.StringDecode(data, je.JavaConfigs.JarDecode))
		} else {
			log.Println(je.JavaConfigs.Name, "退出程序")
			break
		}
	}
	if je.Process != nil {
		je.Process.Pid=-1
	}

}

func (je *JavaExec) recordRecover() {
	err := recover()
	if err != nil {
		je.Stop()
		log.Println("系统级错误:  %s", err)
	}
}

func (je *JavaExec) GetName() string {
	return je.Name
}
func (je *JavaExec) Start() string {
	if je.HasRun() {
		return config.AlreadyRunning
	}
	je.JavaConfigs.Run = true
	je.Run()
	return config.SUCCESS
}
func (je *JavaExec) Stop() string {
	process,err:=FindJavaProcess(je.JavaConfigs.JarArgs)
	if err==nil{
		if len(process)==0{
			return config.NoRunning
		}
		for i, v := range process {
			println(i, v.Name,v.ProcessId,v.CommandLine)
			err:=KillJavaProcess(v.ProcessId)
			if err!=nil{
				return config.FAIL
			}else{
				return config.SUCCESS
			}
		}
	}else {
		log.Panicln("不支持 WMI",err)
	}
	return config.NoRunning
}
