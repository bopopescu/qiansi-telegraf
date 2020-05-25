// +build !windows

package telegraf

func run(inputFilters, outputFilters, aggregatorFilters, processorFilters []string) {
	stop = make(chan struct{})
	reloadLoop(
		inputFilters,
		outputFilters,
		aggregatorFilters,
		processorFilters,
	)
}
