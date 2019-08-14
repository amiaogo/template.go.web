package main

import (
	"flag"
	"fmt"
	"gitee.com/piupuer/go/tools"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"path/filepath"
	"strings"
	_ "template.go.web/web"
)

func main() {
	// 通过flag从运行环境中读取配置参数
	// app.conf绝对路径
	conf := flag.String("conf", "conf/app.conf", "app.conf file path")
	// 静态文件目录
	staticDir := flag.String("static.dir", "static", "static dir path")
	// 视图目录
	viewsDir := flag.String("views.dir", "views", "views dir path")

	// 转换环境变量
	flag.Parse()
	if *conf != "" {
		// 读取配置写入beego
		beego.LoadAppConfig("ini", *conf)
	}
	if *viewsDir != "" {
		beego.SetViewsPath(*viewsDir)
	}
	if *staticDir != "" {
		beego.SetStaticPath("/static", *staticDir)
	}
	logMode, err := beego.AppConfig.Int("logmode")
	if err == nil {
		logs.SetLevel(logMode)
	} else {
		logs.SetLevel(logs.LevelInformational)
	}
	// 是否开启orm调试
	orm.Debug, _ = beego.AppConfig.Bool("orm.debug")
	// 调试模式
	if beego.AppConfig.String("runmode") == "dev" {
		logs.SetLogger(logs.AdapterConsole)
	} else {
		logFile := filepath.Join("logs", beego.AppConfig.String("appname")+".log")
		_, err = tools.CreateFileIfNotExists(logFile)
		if err != nil {
			logs.Error(fmt.Sprintf("[日志初始化] 创建log文件失败 异常 %s", err.Error()))
			panic(err)
		}
		log := fmt.Sprintf(`{"filename":"%s"}`, logFile)
		// 这里统一使用正斜杠, 不然json解析可能报错
		log = strings.Replace(log, "/", "\\", -1)
		log = strings.Replace(log, "\\", "/", -1)
		beego.SetLogger("file", log)
	}

	// beego启动
	beego.Run()
}
