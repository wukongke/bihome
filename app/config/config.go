package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode/utf8"
	"work-codes/bihome/app/common"
)

var jsonData map[string]interface{}

func initJSON() {
	bytes, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		os.Exit(-1)
	}

	configStr := string(bytes[:])
	reg := regexp.MustCompile(`/\*.*\*/`)

	configStr = reg.ReplaceAllString(configStr, "")
	bytes = []byte(configStr)

	if err := json.Unmarshal(bytes, &jsonData); err != nil {
		fmt.Println("invalid config: ", err.Error())
		os.Exit(-1)
	}
}

type dBConfig struct {
	Url          string
	DbName       string
	MaxIdleConns int
	MaxOpenConns int
}

// DBConfig 数据库相关配置
var DBConfig dBConfig

func initDB() {
	common.SetStructByJSON(&DBConfig, jsonData["database"].(map[string]interface{}))
}

type redisConfig struct {
	Url string
}

// RedisConfig redis相关配置
var RedisConfig redisConfig

func initRedis() {
	common.SetStructByJSON(&RedisConfig, jsonData["redis"].(map[string]interface{}))
	// url := fmt.Sprintf("%s:%d", RedisConfig.Host, RedisConfig.Port)
	// RedisConfig.URL = url
}

type appConfig struct {
	APIPoweredBy   string
	SiteName       string
	Host           string
	ImgHost        string
	Env            string
	APIPrefix      string
	UploadImgDir   string
	ImgPath        string
	Port           int
	SessionID      string
	SessionTimeout int
	PassSalt       string
	MailUser       string //域名邮箱账号
	MailPass       string //域名邮箱密码
	MailHost       string //smtp邮箱域名
	MailPort       int    //smtp邮箱端口
	MailFrom       string //邮件来源
	Github         string
}

// appConfig 服务器相关配置
var AppConfig appConfig

func initApp() {
	common.SetStructByJSON(&AppConfig, jsonData["app"].(map[string]interface{}))

	if AppConfig.UploadImgDir == "" {
		sep := string(os.PathSeparator)
		execPath := filepath.Dir(os.Args[0])
		pathArr := []string{"website", "static", "upload", "img"}
		length := utf8.RuneCountInString(execPath)
		lastChar := execPath[length-1:]
		if lastChar != sep {
			execPath = execPath + sep
		}
		execPath = execPath + strings.Join(pathArr, sep)
		fmt.Println(execPath)
		AppConfig.UploadImgDir = execPath
	}
}

type apiConfig struct {
	Prefix string
}

// APIConfig api相关配置
var APIConfig apiConfig

func initAPI() {
	common.SetStructByJSON(&APIConfig, jsonData["api"].(map[string]interface{}))
}

func init() {
	initJSON()
	initDB()
	initRedis()
	initApp()
	initAPI()
}
