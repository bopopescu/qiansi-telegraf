// +build !linux

package infiniband

import (
	"gitee.com/zhimiao/qiansi-telegraf"
	"gitee.com/zhimiao/qiansi-telegraf/plugins/inputs"
)

func (i *Infiniband) Init() error {
	i.Log.Warn("Current platform is not supported")
	return nil
}

func (_ *Infiniband) Gather(acc telegraf.Accumulator) error {
	return nil
}

func init() {
	inputs.Add("infiniband", func() telegraf.Input {
		return &Infiniband{}
	})
}
