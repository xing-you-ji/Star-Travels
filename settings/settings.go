package settings

import (
	"fmt"

	"github.com/spf13/viper"
)

//全局配置结构体
var Conf = new(AppConfig)

type AppConfig struct {
	Name               string               `mapstructure:"name"` //mapstructure：通用structTag
	Mode               string               `mapstructure:"mode"`
	Version            string               `mapstructure:"version"`
	Port               int                  `mapstructure:"port"`
	StartTime          string               `mapstructure:"start_time"`
	MachineID          int                  `mapstructure:"machine_id"`
	FileInterval       int64                `mapstructure:"fill_interval"` // 添加令牌的速度
	Cap                int64                `mapstructure:"cap"`           // 令牌桶的容量
	*LogConfig         `mapstructure:"log"` //tag需要与配置文件中的名字对应
	*MysqlMasterConfig `mapstructure:"mysql_master"`
	*MysqlSlaveConfig  `mapstructure:"mysql_slave"`
	*RedisConfig       `mapstructure:"redis"`
	*JaegerConfig `mapstructure:"jaeger"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MysqlMasterConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Dbname       string `mapstructure:"dbname"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type MysqlSlaveConfig struct {
	Count        int      `mapstructure:"count"`
	Host         []string `mapstructure:"host"`
	Port         []int    `mapstructure:"port"`
	User         []string `mapstructure:"user"`
	Password     []string `mapstructure:"password"`
	Dbname       []string `mapstructure:"dbname"`
	MaxOpenConns int      `mapstructure:"max_open_conns"`
	MaxIdleConns int      `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Sentinels []string    `mapstructure:"sentinels"`
	Password  string      `mapstructure:"password"`
	Masters   RedisMaster `mapstructure:"masters"`
	Replicas  int         `mapstructure:"replicas"`
	PoolSize  int         `mapstructure:"pool_size"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
}

type RedisMaster struct {
	MasterName []string `mapstructure:"master_name"`
	Passwords  []string `mapstructure:"passwords"`
	Counts     int      `mapstructure:"counts"`
}

// 链路追踪
type JaegerConfig struct {
	ServiceName string `mapstructure:"service_name"`
	AgentHostPort string `mapstructure:"agent_host_port"`
}


func Init(configFile string) (err error) {
	viper.SetConfigFile(configFile)
	// viper.SetConfigFile("./config.yaml") //指定配置文件（带后缀，可写绝对路径和相对路径两种）
	//viper.SetConfigName("config") //指定配置文件的名字（不带后缀）
	// 基本上是配合远程配置中心使用的，告诉viper当前的数据使用什么格式去解析
	viper.SetConfigType("yaml") //远程配置文件传输 确定配置文件的格式
	viper.AddConfigPath(".")    //指定配置文件的一个寻找路径
	err = viper.ReadInConfig()  //读取配置信息
	if err != nil {
		//读取配置信息错误
		fmt.Printf("viper.ReadInConfig() failed: %v\n", err)
		return
	}
	//把读取到的信息反序列化到 Conf 变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed: %v\n", err)
	}

	return
}
