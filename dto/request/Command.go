package request

type CommandForm struct {
	Ip          string `form:"ip" binding:"required"`
	Port        string `form:"port"`                   //默认5555
	Cmd         string `form:"cmd" binding:"required"` //执行命令，setProxy|delProxy|clear|stop
	ProxyAddr   string `form:"proxyAddr"`              //代理地址 setProxy可用
	PackageName string `form:"packageName"`            //包名 clear|stop可用
}
