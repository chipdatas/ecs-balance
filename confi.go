package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	AccessKeyId     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
	RegionId        string `json:"region_id"`
	InstanceId      string `json:"instance_id"`
	DeployDate      string `json:"deploy_date"`
}

func SetConfig() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(`
	请输入阿里云配置信息:
	开通api 权限，请参考https://help.aliyun.com/zh/ram/user-guide/create-an-accesskey-pair
	请注意：开通后，记得给他配置权限，不熟悉就全选即可`)

	fmt.Print("请输入阿里云 Access Key ID: ")
	accessKeyId, _ := reader.ReadString('\n')
	accessKeyId = strings.TrimSpace(accessKeyId)

	fmt.Print("请输入阿里云 Access Key Secret: ")
	accessKeySecret, _ := reader.ReadString('\n')
	accessKeySecret = strings.TrimSpace(accessKeySecret)

	fmt.Print("请输入阿里云区域 ID: ")
	fmt.Print("区域 ID 可以打开服务器的信息页面，看浏览器的网址，里面有，比如 cn-hangzhou \n")
	regionId, _ := reader.ReadString('\n')
	regionId = strings.TrimSpace(regionId)

	fmt.Print("请输入阿里云实例 ID: ")
	fmt.Print("实例 ID 可以打开服务器的信息页面获取\n")
	instanceId, _ := reader.ReadString('\n')
	instanceId = strings.TrimSpace(instanceId)

	fmt.Print("请输入最早部署日期 (格式: YYYY-MM-DD，一定要这个格式！): ")
	deployDate, _ := reader.ReadString('\n')
	deployDate = strings.TrimSpace(deployDate)

	config := Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		RegionId:        regionId,
		InstanceId:      instanceId,
		DeployDate:      deployDate,
	}

	file, err := os.Create("config.json")
	if err != nil {
		fmt.Println("无法创建配置文件:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(config)
	if err != nil {
		fmt.Println("无法写入配置文件:", err)
		return
	}

	fmt.Println("配置文件已成功生成: config.json")
}

// 获取解析后的配置
func GetConfig() Config {
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println("无法打开配置文件:", err)
		return Config{}
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("无法解析配置文件:", err)
		return Config{}
	}

	return config
}

func NeedConfig() bool {
	_, err := os.Stat("config.json")
	return os.IsNotExist(err)
}
