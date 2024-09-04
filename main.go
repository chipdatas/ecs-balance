// This file is auto-generated, don't edit it. Thanks.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v4/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

type InstanceMonitorData struct {
	BPSRead                      int     `json:"BPSRead"`
	BPSWrite                     int     `json:"BPSWrite"`
	CPU                          int     `json:"CPU"`
	CPUAdvanceCreditBalance      int     `json:"CPUAdvanceCreditBalance"`
	CPUCreditBalance             float64 `json:"CPUCreditBalance"`
	CPUCreditUsage               float64 `json:"CPUCreditUsage"`
	CPUNotpaidSurplusCreditUsage int     `json:"CPUNotpaidSurplusCreditUsage"`
	IOPSRead                     int     `json:"IOPSRead"`
	IOPSWrite                    int     `json:"IOPSWrite"`
	InstanceId                   string  `json:"InstanceId"`
	InternetBandwidth            int     `json:"InternetBandwidth"`
	InternetRX                   int     `json:"InternetRX"`
	InternetTX                   int     `json:"InternetTX"`
	IntranetBandwidth            int     `json:"IntranetBandwidth"`
	IntranetRX                   int     `json:"IntranetRX"`
	IntranetTX                   int     `json:"IntranetTX"`
	TimeStamp                    string  `json:"TimeStamp"`
}

type MonitorData struct {
	InstanceMonitorData []InstanceMonitorData `json:"InstanceMonitorData"`
}

type ResponseBody struct {
	MonitorData MonitorData `json:"MonitorData"`
	RequestId   string      `json:"RequestId"`
}

type APIResponse struct {
	Headers    map[string]string `json:"headers"`
	StatusCode int               `json:"statusCode"`
	Body       ResponseBody      `json:"body"`
}

func CreateClient() (_result *ecs20140526.Client, _err error) {

	conf := GetConfig()

	// 设置阿里云的 AccessKey ID 和 AccessKey Secret
	accessKeyId := conf.AccessKeyId
	accessKeySecret := conf.AccessKeySecret
	regionId := conf.RegionId

	// 工程代码泄露可能会导致 AccessKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考。
	// 建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html。
	config := &openapi.Config{
		// 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_ID。
		AccessKeyId: tea.String(accessKeyId),
		// 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_SECRET。
		AccessKeySecret: tea.String(accessKeySecret),
	}
	// Endpoint 请参考 https://api.aliyun.com/product/Ecs
	config.Endpoint = tea.String("ecs." + regionId + ".aliyuncs.com")
	_result = &ecs20140526.Client{}
	_result, _err = ecs20140526.NewClient(config)
	return _result, _err
}

func _main(startTime, endTime string) (msi map[string]interface{}, _err error) {
	client, _err := CreateClient()
	if _err != nil {
		return nil, _err
	}
	conf := GetConfig()
	describeInstanceMonitorDataRequest := &ecs20140526.DescribeInstanceMonitorDataRequest{
		InstanceId: tea.String(conf.InstanceId),
		Period:     tea.Int32(3600),
		StartTime:  tea.String(startTime), // "2024-09-03T01:04:05Z"
		EndTime:    tea.String(endTime),
	}
	runtime := &util.RuntimeOptions{}
	return func() (msi map[string]interface{}, _e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		res, _err := client.DescribeInstanceMonitorDataWithOptions(describeInstanceMonitorDataRequest, runtime)
		if _err != nil {
			return nil, _err
		}

		// 将响应转换为 JSON 字符串
		jsonString := res.String()

		// 解析 JSON 字符串
		var apiResponse APIResponse
		err := json.Unmarshal([]byte(jsonString), &apiResponse)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			return nil, err
		}

		totalInternetRX := 0
		totalInternetTX := 0
		totalIntranetRX := 0
		totalIntranetTX := 0
		totalInternetBandwidth := 0
		totalIntranetBandwidth := 0
		balance_cpu := 0.0

		for _, data := range apiResponse.Body.MonitorData.InstanceMonitorData {
			totalInternetRX += data.InternetRX
			totalInternetTX += data.InternetTX
			totalIntranetRX += data.IntranetRX
			totalIntranetTX += data.IntranetTX
			totalInternetBandwidth += data.InternetBandwidth
			totalIntranetBandwidth += data.IntranetBandwidth
			balance_cpu = data.CPUCreditBalance
		}

		// bytes to Mbps
		totalInternetRX = totalInternetRX / 1024 / 1024
		totalInternetTX = totalInternetTX / 1024 / 1024
		totalIntranetRX = totalIntranetRX / 1024 / 1024
		totalIntranetTX = totalIntranetTX / 1024 / 1024

		// fmt.Printf("公网带宽: %d Mbps\n", totalInternetBandwidth)
		// fmt.Printf("公网接收的字节数: %d Mb\n", totalInternetRX)
		// fmt.Printf("公网发送的字节数: %d Mb\n", totalInternetTX)
		// fmt.Printf("内网接收的字节数: %d Mb\n", totalIntranetRX)
		// fmt.Printf("内网发送的字节数: %d Mb\n", totalIntranetTX)
		// fmt.Printf("内网带宽: %d Mbps\n", totalIntranetBandwidth)
		// fmt.Printf("[突发实例]百分百CPU剩余运行分钟数: %f分钟", balance_cpu)

		return map[string]interface{}{
			"totalInternetRX":        totalInternetRX,
			"totalInternetTX":        totalInternetTX,
			"totalIntranetRX":        totalIntranetRX,
			"totalIntranetTX":        totalIntranetTX,
			"totalInternetBandwidth": totalInternetBandwidth,
			"totalIntranetBandwidth": totalIntranetBandwidth,
			"balance_cpu":            balance_cpu,
		}, nil

	}()

}

func main() {

	// 如何有传一个 auto 参数，则自动添加到 shell 配置文件
	if len(os.Args) > 1 && os.Args[1] == "auto" {
		SetShell()
		println("已执行添加任务，如果失败需要手动添加到 shell 配置文件.")
		os.Exit(0)
	}

	if NeedConfig() {
		SetConfig()
		fmt.Println("配置文件已成功生成,如果需要重新配置请先删除config.json文件再重新运行程序.\n配置完成后，查看最新消息请再次运行程序.")
		os.Exit(0)
	}

	// 获取配置文件
	config := GetConfig()
	if config.AccessKeyId == "" || config.AccessKeySecret == "" || config.RegionId == "" || config.InstanceId == "" || config.DeployDate == "" {
		fmt.Println("配置文件不完整，请删除或者 config.json 后重新运行程序.")
		os.Exit(0)
	}

	// 获取从今天往前推 15天和 30天的 2 个日期,
	dates := []time.Time{}
	today := time.Now()
	before30 := today.AddDate(0, 0, -30)
	before15 := today.AddDate(0, 0, -15)
	// 如果 before15 在部署日期之前，则使用部署日期
	deployDate, _ := time.Parse("2006-01-02", config.DeployDate)
	if before15.Before(deployDate) {
		before15 = deployDate
		dates = append(dates, before15)
	} else if before30.Before(deployDate) {
		before30 = deployDate
		dates = append(dates, before30)
		dates = append(dates, before15)
	} else {
		dates = append(dates, before30)
		dates = append(dates, before15)
	}

	bigData := map[string]interface{}{
		"totalInternetRX":        0,
		"totalInternetTX":        0,
		"totalIntranetRX":        0,
		"totalIntranetTX":        0,
		"totalInternetBandwidth": 0,
		"totalIntranetBandwidth": 0,
		"balance_cpu":            0.0,
	}
	for _, date := range dates {
		startTime := date.Format("2006-01-02") + "T00:00:00Z"
		endTime := date.AddDate(0, 0, 15).Format("2006-01-02") + "T23:59:59Z"
		msi, err := _main(startTime, endTime)
		if err != nil {
			fmt.Println("Error:", err)
		}
		for k, v := range msi {
			if k == "balance_cpu" {
				bigData[k] = v.(float64)
			} else {
				bigData[k] = bigData[k].(int) + v.(int)
			}
		}
	}

	// 创建表格
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"指标", "值"})

	// 定义颜色
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	magenta := color.New(color.FgMagenta).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	// 添加数据到表格
	println("\n")
	table.Append([]string{green("公网带宽"), fmt.Sprintf("%d Mbps", bigData["totalInternetBandwidth"])})
	table.Append([]string{red("公网接收的字节数"), fmt.Sprintf("%d Mb", bigData["totalInternetRX"])})
	table.Append([]string{yellow("公网发送的字节数"), fmt.Sprintf("%d Mb", bigData["totalInternetTX"])})
	table.Append([]string{blue("内网接收的字节数"), fmt.Sprintf("%d Mb", bigData["totalIntranetRX"])})
	table.Append([]string{magenta("内网发送的字节数"), fmt.Sprintf("%d Mb", bigData["totalIntranetTX"])})
	table.Append([]string{cyan("内网带宽"), fmt.Sprintf("%d Mbps", bigData["totalIntranetBandwidth"])})
	table.Append([]string{green("[突发实例]百分百CPU剩余运行分钟数"), fmt.Sprintf("%f 分钟", bigData["balance_cpu"])})

	// 渲染表格
	table.Render()

}
