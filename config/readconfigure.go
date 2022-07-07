package config

import (
	log "github.com/sirupsen/logrus"
	"strings"
)

var javaConfig = NewJavaConfig()

var httpConfig = NewHttpConfig()

var scheduleConfig = NewScheduleConfig()

func GetJavaConfig() *JavaConfig {
	return javaConfig
}
func GetHttpConfig() *HttpConfig {
	return httpConfig
}
func GetScheduleConfig() *ScheduleConfig {
	return scheduleConfig
}


func Init() {
	log.Println("初始化系统配置")

	cfg, err := newConfigure()
	if err == nil {
		if cfg.HasSection(HTTP) {
			log.Println("读取HTTP配置文件中")
			if cfg.HasOption(HTTP, HTTP_PORT) {
				v, err := cfg.RawString(HTTP, HTTP_PORT)
				if err == nil {
					httpConfig.SetHttpPort(v)
				}
			}
		}

		if cfg.HasSection(JAVA) {
			log.Println("读取JAVA配置文件中")
			if cfg.HasOption(JAVA, JAVA_HOME) {
				v, err := cfg.RawString(JAVA, JAVA_HOME)
				if err == nil {
					v = strings.ReplaceAll(v, "\\", "/")
					javaConfig.SetJavaHome(v)
					log.Println("读取JAVA_HOME", v)
				}
			}

			if cfg.HasOption(JAVA, JAVA_START) {
				v, err := cfg.RawString(JAVA, JAVA_START)
				if err == nil {
					javaConfig.SetStart(v)
					log.Println("读取JAVA_START", v)
				}
			}

			if cfg.HasOption(JAVA, JAVA_BIN) {
				v, err := cfg.RawString(JAVA, JAVA_BIN)
				if err == nil {
					javaConfig.SetJavaBin(v)
					log.Println("读取JAVA_BIN", v)
				}
			} else {
				javaConfig.SetJavaBin("java")
			}

			if cfg.HasOption(JAVA, JAVA_OPS) {
				v, err := cfg.RawString(JAVA, JAVA_OPS)
				if err == nil {
					javaConfig.SetJavaOps(v)
					log.Println("读取JAVA_OPS", v)
				}
			}
			if cfg.HasOption(JAVA, JAVA_EXEC) {
				v, err := cfg.RawString(JAVA, JAVA_EXEC)
				if err == nil {
					javaConfig.SetJavaExec(v)
					log.Println("读取JAVA_EXEC", v)
				}
			}

			if len(javaConfig.JavaExec) > 0 {
				jars := strings.Split(strings.ToUpper(javaConfig.JavaExec), ",")
				for _, val := range jars {
					v := strings.TrimSpace(strings.ToUpper(val))
					if cfg.HasSection(v) {
						log.Println("读取JAR", v)
						jarConfig := NewJarConfig(v)
						javaConfig.AddJarConfig(jarConfig)
						jarRun, err := cfg.RawString(v, JAR_RUN)
						if err == nil {
							jarConfig.SetRun(jarRun)
							log.Println("读取JAR", v, "JAR_RUN", jarRun)
						}

						jarPath, err := cfg.RawString(v, JAR_PATH)
						if err == nil {
							jarPath = strings.ReplaceAll(jarPath, "\\", "/")
							jarConfig.SetJarPath(jarPath)
							log.Println("读取JAR", v, "JAR_PATH", jarPath)
						}

						jarArgs, err := cfg.RawString(v, JAR_ARGS)
						if err == nil {
							jarConfig.SetJarArgs(jarArgs)
							log.Println("读取JAR", v, "JAR_ARGS", jarArgs)
						}

						jarDecode, err := cfg.RawString(v, JAR_DECODE)
						if err == nil {
							jarConfig.SetJarDecode(jarDecode)
							log.Println("读取JAR", v, "JAR_DECODE", jarDecode)
						} else {
							jarConfig.SetJarDecode("UTF-8")
						}

						javaHome, err := cfg.RawString(v, JAVA_HOME)
						if err == nil {
							javaHome = strings.ReplaceAll(javaHome, "\\", "/")
							jarConfig.SetJavaHome(javaHome)
							log.Println("读取JAR", v, "JAVA_HOME", javaHome)
						}

						javaBin, err := cfg.RawString(v, JAVA_BIN)
						if err == nil {
							jarConfig.SetJavaBin(javaBin)
							log.Println("读取JAR", v, "JAVA_BIN", javaBin)
						}

						jarOps, err := cfg.RawString(v, JAVA_OPS)
						if err == nil {
							jarConfig.SetJavaOps(jarOps)
							log.Println("读取JAR", v, "JAVA_OPS", jarOps)
						}

					} else {
						log.Error("找不到配置:", v)
					}

				}

			}
		} else {
			log.Panicln("java配置为空启动失败")
		}

		if cfg.HasSection(SCHEDULE) {
			log.Println("读取定时配置文件中")
			if cfg.HasOption(SCHEDULE, SCHEDULE_START) {
				scheduleStart, err := cfg.RawString(SCHEDULE, SCHEDULE_START)
				if err == nil {
					log.Println("读取 SCHEDULE_START", scheduleStart)
					scheduleConfig.SetStart(scheduleStart)
				}
			}
			if cfg.HasOption(SCHEDULE, SCHEDULE_JOBS) {
				v, err := cfg.RawString(SCHEDULE, SCHEDULE_JOBS)
				if err == nil {
					log.Println("读取 SCHEDULE_JOBS", v)
					scheduleConfig.SetJobs(v)
				}
			}

			if len(scheduleConfig.Jobs) > 0 {
				jobs := strings.Split(strings.ToUpper(scheduleConfig.Jobs), ",")
				for _, v := range jobs {
					v = strings.TrimSpace(strings.ToUpper(v))
					jobConfig := NewJobConfig(v)
					scheduleConfig.AddJobConfig(jobConfig)
					if cfg.HasOption(v, JOB_RUN) {
						jobRun, err := cfg.RawString(v, JOB_RUN)
						if err == nil {
							log.Println("读取 JOB_RUN", v,jobRun)
							jobConfig.SetRun(jobRun)
						}
					}
					if cfg.HasOption(v, JOB_CRON) {
						jobCron, err := cfg.RawString(v, JOB_CRON)
						if err == nil {
							log.Println("读取 JOB_CRON", v,jobCron)
							jobConfig.SetCron(jobCron)
						}
					}
					if cfg.HasOption(v, JOB_EXEC) {
						jobExec, err := cfg.RawString(v, JOB_EXEC)
						if err == nil {
							log.Println("读取 JOB_EXEC", v,jobExec)
							jobConfig.SetExec(jobExec)
						}
					}

					if cfg.HasOption(v, JOB_MODE) {
						jobMode, err := cfg.RawString(v, JOB_MODE)
						if err == nil {
							log.Println("读取 JOB_MODE", v,jobMode)
							jobConfig.SetMode(jobMode)
						}
					}

				}
			}

		}

	} else {
		log.Panicln("读取系统配置失败", err)
	}
}
