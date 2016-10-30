package config

import (
    "github.com/kataras/iris"
    "os"
    "encoding/json"
    "fmt"
)


func init(){
    //定义一个salt, 用于混入原始密码，然后计算多次hash,更加不容易被cracker通过彩虹表泄漏原始密码
    salt := "mysaltstring"

    //定义一个存储配置结构体的map，分别存储开发，测试，生产的不同配置
    var configMap = map[string]iris.Configuration{}
    //从配置文件加载配置数据
    cf, errOpen := os.Open(`config.json`)
    if errOpen != nil {
        fmt.Println("Open config file error:", errOpen)
    }
    defer func(){
        if err := cf.Close(); err != nil {
            panic(err)
        }
    }()

    //配置文件格式为json，需要一个json Decoder，并调用其Decode方法
    decoder := json.NewDecoder(cf)
    errDecode := decoder.Decode(&configMap)
    if errDecode != nil {
        fmt.Println("Decoding config.json:", errDecode)
    }

    //加载配置之后，先合并基础配置，再合并特定环境的配置
    configMap["BASIC"].Set(iris.Config)
    configMap["DEV"].Set(iris.Config)
    fmt.Println("########当前使用配置#########")
    fmt.Println("iris.Config.Gzip:",iris.Config.Gzip)
    fmt.Println("iris.Config.IsDevelopment:",iris.Config.IsDevelopment)
    fmt.Println(`iris.Config.Other["MySQLDSN"]:`,iris.Config.Other["MySQLDSN"])
    iris.Config.Other["salt"] = salt
    fmt.Println(`iris.Config.Other["salt"]:`,iris.Config.Other["salt"])


}

