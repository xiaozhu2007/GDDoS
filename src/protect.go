package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
	//"github.com/gonutz/ide/w32"
)

/*
func ShowConsoleAsync(commandShow uintptr) {
	console := w32.GetConsoleWindow()
	if console != 0 {
		_, consoleProcID := w32.GetWindowThreadProcessId(console)
		if w32.GetCurrentProcessId() == consoleProcID {
			w32.ShowWindowAsync(console, commandShow)
		}
	}
}
*/
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func check_virtual() (bool, error) { // 识别虚拟机
	model := ""
	var cmd *exec.Cmd
	cmd = exec.Command("cmd", "/C", "wmic path Win32_ComputerSystem get Model")
	stdout, err := cmd.Output()
	if err != nil {
		return false, err
	}
	model = strings.ToLower(string(stdout))
	if strings.Contains(model, "VirtualBox") || strings.Contains(model, "virtual") || strings.Contains(model, "VMware") ||
		strings.Contains(model, "KVM") || strings.Contains(model, "Bochs") || strings.Contains(model, "HVM domU") || strings.Contains(model, "Parallels") {
		return true, nil // 如果是虚拟机则返回true
	}
	return false, nil
}

func fvm(path string) {
	pe, _ := PathExists(path)
	if pe {
		log.Printf("We found virtual machine environment.")
		os.Exit(1)
	}
}

func check_file() {
	fvm("C:\\Windows\\System32\\Drivers\\Vmmouse.sys")
	fvm("C:\\Windows\\System32\\Drivers\\vmtray.dll")
	fvm("C:\\Windows\\System32\\Drivers\\VMToolsHook.dll")
	fvm("C:\\Windows\\System32\\Drivers\\vmmousever.dll")
	fvm("C:\\Windows\\System32\\Drivers\\vmhgfs.dll")
	fvm("C:\\Windows\\System32\\Drivers\\vmGuestLib.dll")
	fvm("C:\\Windows\\System32\\Drivers\\VBoxMouse.sys")
	fvm("C:\\Windows\\System32\\Drivers\\VBoxGuest.sys")
	fvm("C:\\Windows\\System32\\Drivers\\VBoxSF.sys")
	fvm("C:\\Windows\\System32\\Drivers\\VBoxVideo.sys")
	fvm("C:\\Windows\\System32\\vboxdisp.dll")
	fvm("C:\\Windows\\System32\\vboxhook.dll")
	fvm("C:\\Windows\\System32\\vboxoglerrorspu.dll")
	fvm("C:\\Windows\\System32\\vboxoglpassthroughspu.dll")
	fvm("C:\\Windows\\System32\\vboxservice.exe")
	fvm("C:\\Windows\\System32\\vboxtray.exe")
	fvm("C:\\Windows\\System32\\VBoxControl.exe")
	fvm("C:\\Program Files\\VMware\\VMware Tools\\")
}

func protectMain() {
	check_file()
	check_virtual()
	// ShowConsoleAsync(w32.SW_HIDE) // 隐藏控制台窗口
}
