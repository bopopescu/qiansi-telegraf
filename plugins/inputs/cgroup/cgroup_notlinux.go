// +build !linux

package cgroup

import (
	"gitee.com/zhimiao/qiansi-telegraf"
)

func (g *CGroup) Gather(acc telegraf.Accumulator) error {
	return nil
}
