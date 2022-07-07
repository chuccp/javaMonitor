package execute

import (
	"github.com/StackExchange/wmi"
	"os/exec"
	"strconv"
)

type Win32_Process struct {
	Name string
	ProcessId uint32
	CommandLine string
}

func FindJavaProcess(CommandLine string)([]*Win32_Process,error)  {
	ws:=make([]*Win32_Process,0)
	var dst []Win32_Process
	q := wmi.CreateQuery(&dst, "where Name like '%java%' and CommandLine like '%"+CommandLine+"%'")
	err := wmi.Query(q, &dst)
	if err != nil {
		return nil,err
	}
	for _, v := range dst {
		ws = append(ws, &v)
	}
	return ws,nil
}

func KillJavaProcess(ProcessId uint32)error{
	cmd:=exec.Command("taskkill",`/T`,`/F`,`/pid`,strconv.Itoa((int)(ProcessId)))
	return cmd.Run()
}

type Win32_NetworkAdapter struct {
	AdapterType string
	Speed uint64
	AdapterTypeID  uint16
	GUID string
	Status string
	MaxSpeed uint64
}