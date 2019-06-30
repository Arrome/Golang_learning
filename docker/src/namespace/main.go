package main

import (
	"os/exec"
	"syscall"
	"os"
	"log"
)

/**
	验证UTS：pstree -pl 命令查看进程树，确认main.go进程和子进程sh进程，不在同一个uts namespace；hostname修改不了主进程的
	验证IPC：ipcs -q 查看队列，ipcmk -Q创建队列，在另一个运行中查看ipcs -q
	验证PID：pstree -pl 查看宿主实际PID，echo $$ 确认当前PID
	验证NS: 直接ls /proc看到的还是宿主机的，通过mount -t proc proc /proc查看到的即是当前，ps -ef
 */
func main() {
	cmd := exec.Command("sh");
	//封装了Namespace clone()函数的调用，执行后进入sh运行环境中

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET,// | syscall.CLONE_NEWUSER,

	}
	//cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uint32(1), Gid: uint32(1)} //CLONE_NEWUSER
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run();err != nil {
		log.Fatal(err)
	}

	os.Exit(-1)
}
