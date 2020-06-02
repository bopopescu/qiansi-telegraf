// +build windows

package telegraf

import (
	"runtime"
)

func run(inputFilters, outputFilters, aggregatorFilters, processorFilters []string) {
	if runtime.GOOS == "windows" && windowsRunAsService() {
		runAsWindowsService(
			inputFilters,
			outputFilters,
			aggregatorFilters,
			processorFilters,
		)
	} else {
		stop = make(chan struct{})
		reloadLoop(
			inputFilters,
			outputFilters,
			aggregatorFilters,
			processorFilters,
		)
	}
}
