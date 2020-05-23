package riemann_legacy

import (
	"testing"

	"gitee.com/zhimiao/qiansi-telegraf/testutil"
	"github.com/stretchr/testify/require"
)

func TestConnectAndWrite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	url := testutil.GetLocalHost() + ":5555"

	r := &Riemann{
		URL:       url,
		Transport: "tcp",
	}

	err := r.Connect()
	require.NoError(t, err)

	err = r.Write(testutil.MockMetrics())
	require.NoError(t, err)
}
