// +build !linux,!freebsd

package zfs

import (
	"gitee.com/zhimiao/qiansi-telegraf"
	"gitee.com/zhimiao/qiansi-telegraf/plugins/inputs"
)

func (z *Zfs) Gather(acc telegraf.Accumulator) error {
	return nil
}

func init() {
	inputs.Add("zfs", func() telegraf.Input {
		return &Zfs{}
	})
}
