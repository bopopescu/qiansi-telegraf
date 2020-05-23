package outputs

import (
	"gitee.com/zhimiao/qiansi-telegraf"
)

type Creator func() telegraf.Output

var Outputs = map[string]Creator{}

func Add(name string, creator Creator) {
	Outputs[name] = creator
}
