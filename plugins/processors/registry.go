package processors

import "gitee.com/zhimiao/qiansi-telegraf"

type Creator func() telegraf.Processor

var Processors = map[string]Creator{}

func Add(name string, creator Creator) {
	Processors[name] = creator
}
