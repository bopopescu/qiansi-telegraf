// +build !linux

package dmcache

import (
	"gitee.com/zhimiao/qiansi-telegraf"
)

func (c *DMCache) Gather(acc telegraf.Accumulator) error {
	return nil
}

func dmSetupStatus() ([]string, error) {
	return []string{}, nil
}
