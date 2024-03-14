package core

import (
	"fmt"
	"net"
	"strings"
)

func GetLocalIp() (string, error) {
	// 获取本机所有网络接口的列表
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// 遍历每个网络接口
	for _, iface := range interfaces {
		// 跳过环回接口和没有IP地址的接口
		if iface.Flags&net.FlagLoopback != 0 || iface.Flags&net.FlagUp == 0 {
			continue
		}

		// 跳过非以太网接口
		if !strings.HasPrefix(iface.Name, "以太网") {
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
