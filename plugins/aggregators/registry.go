package aggregators

import "gitee.com/zhimiao/qiansi-telegraf"

type Creator func() telegraf.Aggregator

var Aggregators = map[string]Creator{}

func Add(name string, creator Creator) {
	Aggregators[name] = creator
}
