package temp

import (
	"testing"

	"github.com/shirou/gopsutil/host"
	"github.com/stretchr/testify/require"

	"gitee.com/zhimiao/qiansi-telegraf/plugins/inputs/system"
	"gitee.com/zhimiao/qiansi-telegraf/testutil"
)

func TestTemperature(t *testing.T) {
	var mps system.MockPS
	var err error
	defer mps.AssertExpectations(t)
	var acc testutil.Accumulator

	ts := host.TemperatureStat{
		SensorKey:   "coretemp_sensor1_crit",
		Temperature: 60.5,
	}

	mps.On("Temperature").Return([]host.TemperatureStat{ts}, nil)

	err = (&Temperature{ps: &mps}).Gather(&acc)
	require.NoError(t, err)

	expectedFields := map[string]interface{}{
		"temp": float64(60.5),
	}

	expectedTags := map[string]string{
		"sensor": "coretemp_sensor1_crit",
	}
	acc.AssertContainsTaggedFields(t, "temp", expectedFields, expectedTags)

}
