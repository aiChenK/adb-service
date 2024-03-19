package core

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// 获取本机ip地址
func GetLocalIp() (string, error) {
	// 获取本机所有网络接口的列表
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	netPrefix := ""
	if runtime.GOOS == "darwin" {
		netPrefix = "en"
	} else if runtime.GOOS == "windows" {
		netPrefix = "以太网"
	} else {
		return "仅支持win及mac获取ip", nil
	}

	// 遍历每个网络接口
	for _, iface := range interfaces {
		// 跳过环回接口和没有IP地址的接口
		if iface.Flags&net.FlagLoopback != 0 || iface.Flags&net.FlagUp == 0 {
			continue
		}

		// 跳过非以太网接口
		if !strings.HasPrefix(iface.Name, netPrefix) {
			continue
		}

		// 获取接口上的所有地址
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// 遍历接口上的每个地址
		for _, addr := range addrs {
			// 将地址转换为IPNet结构，这样可以更容易地提取IP地址
			var ipNet *net.IPNet
			switch v := addr.(type) {
			case *net.IPNet:
				ipNet = v
			case *net.IPAddr:
				ipNet = &net.IPNet{IP: v.IP, Mask: net.CIDRMask(32, 32)}
			}

			if ipNet == nil || ipNet.IP == nil {
				continue
			}

			// 只打印IPv4地址
			if ip4 := ipNet.IP.To4(); ip4 != nil {
				return ip4.String(), nil
			}
		}
	}
	return "", nil
}

// CheckAdb 函数用于检查adb环境变量配置是否正确
func CheckAdb() {
	cmd := exec.Command("adb", "version")
	err := cmd.Run()
	if err != nil {
		PrintErr("adb命令检查失败，请检查adb环境变量配置")
		os.Exit(0)
	}
}

// PrintErr 函数用于在终端打印红色错误信息
// 参数msg为需要打印的错误信息字符串
func PrintErr(msg string) {
	fmt.Print("\033[31m")
	fmt.Println(msg)
	fmt.Print("\033[0m")
}

// OpenBrowser 函数用于打开浏览器，并跳转到指定url地址
func OpenBrowser(url string) {
	fmt.Println("执行打开地址：" + url)

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default: // Linux, FreeBSD, etc.
		cmd = exec.Command("xdg-open", url)
	}

	err := cmd.Run()
	if err != nil {
		PrintErr("浏览器打开失败，请手动打开！")
	}
}
