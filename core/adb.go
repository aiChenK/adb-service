package core

import (
	"adb-service/dto/request"
	"fmt"
	"os/exec"
	"strings"
)

func Exec(commandForm request.CommandForm) error {
	//连接设备
	targetAddr := commandForm.Ip + ":" + commandForm.Port
	err := deviceConnect(targetAddr, false)
	if err != nil {
		return err
	}

	//处理命令
	commandStr := "-s " + targetAddr + " "
	switch commandForm.Cmd {
	case "setProxy":
		commandStr += "shell settings put global http_proxy " + commandForm.ProxyAddr
		break
	case "delProxy":
		commandStr += "shell settings put global http_proxy :0"
		break
	case "clear":
		commandStr += "shell pm clear " + commandForm.PackageName
		break
	case "stop":
		commandStr += "shell am force-stop " + commandForm.PackageName
		break
	}

	cmd := exec.Command("adb", strings.Split(commandStr, " ")...)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
	if err != nil {
		return err
	}
	return nil
}

func getDevices() (deviceMap map[string]string, err error) {
	//执行命令
	cmd := exec.Command("adb", "devices")
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
	if err != nil {
		fmt.Println("命令执行失败:", err)
		return
	}

	// 过滤掉第一行
	lines := strings.Split(string(output), "\n")
	lines = lines[1:]

	// 分割每行并存储到map中
	deviceMap = make(map[string]string)
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			deviceID := fields[0]
			deviceStatus := fields[1]
			deviceMap[deviceID] = deviceStatus
		}
	}

	return
}

func deviceConnect(device string, isRetry bool) error {
	deviceMap, err := getDevices()
	if err != nil {
		return err
	}

	//判断状态是否正常,否则重连
	if status, ok := deviceMap[device]; isRetry == false && (!ok || status != "device") {
		//状态异常则全部断开
		if ok && status != "device" {
			_, err = exec.Command("adb", "disconnect").Output()
			if err != nil {
				return err
			}
		}
		//重连
		output, err := exec.Command("adb", "connect", device).Output()
		fmt.Println(string(output))
		if err != nil {
			return err
		}

		//重新检查
		err = deviceConnect(device, true)
		if err != nil {
			return err
		}
	}
	return nil
}
