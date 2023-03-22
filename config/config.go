package config

import (
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
	"github.com/zgwit/iot-master/v3/pkg/web"
	"gopkg.in/yaml.v3"
	"modbus/args"
	"os"
)

// Configure 配置
type Configure struct {
	Node     string       `yaml:"node" json:"node"`
	Web      web.Options  `yaml:"web" json:"web"`
	Database db.Options   `yaml:"database" json:"database"`
	Mqtt     mqtt.Options `yaml:"mqtt" json:"mqtt"`
	Log      log.Options  `yaml:"log" json:"log"`
}

// Config 全局配置
var Config = Configure{
	Node:     "root",
	Web:      web.Default(),
	Database: db.Default(),
	Mqtt:     mqtt.Default(),
	Log:      log.Default(),
}

func init() {
	Config.Node, _ = os.Hostname()
	Config.Web.Addr = ":8088"
	Config.Database.URL = "root:root@tcp(git.zgwit.com:3306)/modbus?charset=utf8"
	//TODO get imei sn
}

// Load 加载
func Load() error {
	//log.Println("加载配置")
	//从参数中读取配置文件名
	filename := args.ConfigPath

	// 如果没有文件，则使用默认信息创建
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return Store()
		//return nil
	} else {
		y, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
			return err
		}
		defer y.Close()

		d := yaml.NewDecoder(y)
		err = d.Decode(&Config)
		if err != nil {
			log.Fatal(err)
			return err
		}

		return nil
	}
}

func Store() error {
	//log.Println("保存配置")
	//从参数中读取配置文件名
	filename := args.ConfigPath

	y, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755) //os.Create(filename)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer y.Close()

	e := yaml.NewEncoder(y)
	defer e.Close()

	err = e.Encode(&Config)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
