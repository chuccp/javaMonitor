package config

import (
	"github.com/robfig/config"
	log "github.com/sirupsen/logrus"
	"strings"
)

type Configure struct {
	Path    string
	Cg      *config.Config
	section string
}

func newConfigure() (*config.Config, error) {
	cg, err_ := config.ReadDefault("configure.ini")
	if err_ != nil {
		cg, _ = config.ReadDefault("config/configure.ini")
	}
	return cg, err_
}
func (cfg *Configure) PutValue(key string, value string) {
	if !cfg.Cg.HasSection(cfg.section) {
		cfg.Cg.AddSection(cfg.section)
	}
	cfg.Cg.AddOption(cfg.section, key, value)

}

func (cfg *Configure) GetValue(key string) string {

	if cfg.Cg.HasSection(cfg.section) {
		if cfg.Cg.HasOption(cfg.section, key) {
			v, err := cfg.Cg.RawString(cfg.section, key)
			if err == nil {
				v := strings.Trim(v, " ")
				return v

			}

		}
	}
	return ""
}

type HttpConfig struct {
	Port string
}

func (hc *HttpConfig) SetHttpPort(port string) *HttpConfig {
	hc.Port = strings.TrimSpace(port)
	return hc
}

func NewHttpConfig() *HttpConfig {
	return &HttpConfig{Port: "7373"}
}

type JavaConfig struct {
	JavaHome   string
	JavaBin    string
	JavaExec   string
	JavaOps    string
	Start      bool
	JarConfigs []*JarConfig
}

func (javaConfig *JavaConfig) SetStart(start string) *JavaConfig {
	if start == "true" {
		javaConfig.Start = true
	} else {
		javaConfig.Start = false
	}
	return javaConfig
}
func (javaConfig *JavaConfig) SetJavaBin(javaBin string) *JavaConfig {
	javaConfig.JavaBin = javaBin
	return javaConfig
}

func (javaConfig *JavaConfig) SetJavaOps(javaOps string) *JavaConfig {
	javaConfig.JavaOps = javaOps
	return javaConfig
}
func (javaConfig *JavaConfig) SetJavaHome(javaHome string) *JavaConfig {
	javaConfig.JavaHome = javaHome
	return javaConfig
}

func (javaConfig *JavaConfig) SetJavaExec(javaExec string) *JavaConfig {
	javaConfig.JavaExec = javaExec
	return javaConfig
}
func (javaConfig *JavaConfig) AddJarConfig(jarConfig *JarConfig) *JavaConfig {

	if len(jarConfig.JavaOps) == 0 {
		jarConfig.JavaOps = javaConfig.JavaOps
	}
	if len(jarConfig.JavaHome) == 0 {
		jarConfig.JavaHome = javaConfig.JavaHome
	}
	if len(jarConfig.JavaBin) == 0 {
		jarConfig.JavaBin = javaConfig.JavaBin
	}

	javaConfig.JarConfigs = append(javaConfig.JarConfigs, jarConfig)
	return javaConfig
}

func NewJavaConfig() *JavaConfig {
	return &JavaConfig{JarConfigs: make([]*JarConfig, 0), Start: false}
}

type JarConfig struct {
	Name      string
	Run       bool
	JarPath   string
	JarArgs   string
	JarDecode string
	JavaHome  string
	JavaBin   string
	JavaOps   string
}

func NewJarConfig(Name string) *JarConfig {
	return &JarConfig{Name: Name, Run: false}
}
func (JarConfig *JarConfig) SetRun(run string) *JarConfig {

	if run == "true" {
		JarConfig.Run = true
	} else {
		JarConfig.Run = false
	}
	return JarConfig
}

func (JarConfig *JarConfig) SetJarPath(jarPath string) *JarConfig {
	JarConfig.JarPath = jarPath
	return JarConfig
}
func (JarConfig *JarConfig) SetJarArgs(jarArgs string) *JarConfig {
	JarConfig.JarArgs = jarArgs
	return JarConfig
}
func (JarConfig *JarConfig) SetJarDecode(jarDecode string) *JarConfig {
	JarConfig.JarDecode = jarDecode
	return JarConfig
}
func (JarConfig *JarConfig) SetJavaBin(javaBin string) *JarConfig {
	JarConfig.JavaBin = javaBin
	return JarConfig
}

func (JarConfig *JarConfig) SetJavaOps(javaOps string) *JarConfig {
	JarConfig.JavaOps = javaOps
	return JarConfig
}
func (JarConfig *JarConfig) SetJavaHome(javaHome string) *JarConfig {
	JarConfig.JavaHome = javaHome
	return JarConfig
}

type ScheduleConfig struct {
	Start bool
	Jobs string
	JobConfigs []*JobConfig

}
func(scheduleConfig *ScheduleConfig) SetStart(start string)*ScheduleConfig{

	if len(start)>0&&strings.Contains(start,"true"){
		scheduleConfig.Start=true
		log.Println(scheduleConfig.Start,"!!!!!!!!!!!!!!!!!!!!!")
	}
	return scheduleConfig
}
func(scheduleConfig *ScheduleConfig) SetJobs(jobs string)*ScheduleConfig{
	scheduleConfig.Jobs=jobs
	return scheduleConfig
}
func(scheduleConfig *ScheduleConfig) AddJobConfig(jobConfig *JobConfig)*ScheduleConfig{
	scheduleConfig.JobConfigs = append(scheduleConfig.JobConfigs, jobConfig)
	return scheduleConfig
}
func NewScheduleConfig() *ScheduleConfig {
	return &ScheduleConfig{JobConfigs: make([]*JobConfig, 0), Start: false}
}

type JobConfig struct {
	Run bool
	Cron string
	Exec string
	Name string
	Mode string
}
func(jobConfig *JobConfig) SetRun(start string)*JobConfig{
	if len(start)>0&&start=="true"{
		jobConfig.Run=true
	}
	return jobConfig
}
func(jobConfig *JobConfig) SetExec(exec string)*JobConfig{
	jobConfig.Exec=exec
	return jobConfig
}
func(jobConfig *JobConfig) SetCron(cron string)*JobConfig{
	jobConfig.Cron = cron
	return jobConfig
}
func(jobConfig *JobConfig) SetMode(mode string)*JobConfig{

	if strings.Contains(mode,FOLLOW){
		jobConfig.Mode = FOLLOW
	}else{
		jobConfig.Mode = REPLACE
	}
	return jobConfig
}
func NewJobConfig(name string) *JobConfig {
	return &JobConfig{ Run: false,Name:name}
}