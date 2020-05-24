package qiansi

import (
	"fmt"
	"gitee.com/zhimiao/qiansi-telegraf"
	"gitee.com/zhimiao/qiansi-telegraf/plugins/outputs"
	"gitee.com/zhimiao/qiansi-telegraf/plugins/serializers"
	"time"
)

type Qiansi struct {
	Ok bool `toml:"ok"`
}

func (s *Qiansi) Description() string {
	return "qiansi client output"
}

func (s *Qiansi) SampleConfig() string {
	return ` 
  ok = true
`
}

func (s *Qiansi) Init() error {
	return nil
}

func (s *Qiansi) Connect() error {
	// Make a connection to the URL here
	return nil
}

func (s *Qiansi) Close() error {
	// Close connection to the URL here
	return nil
}

func (s *Qiansi) Write(metrics []telegraf.Metric) error {
	ser, _ := serializers.NewJsonSerializer(time.Second)
	r, _ := ser.SerializeBatch(metrics)
	fmt.Println(string(r))
	/*	if common.Config.ENV() == "dev" {
			fmt.Println(string(r))
		}
		// 推送到千丝服务端
		request.MetricPush(r)*/
	return nil
}

func init() {
	outputs.Add("qiansi", func() telegraf.Output { return &Qiansi{} })
}
