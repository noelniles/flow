package main

import (
	"fmt"
	"github.com/noelniles/flow/modules/strings"
	"github.com/noelniles/flow/processing"
	"runtime"
)

func init_threads() {
	nthreads := runtime.NumCPU() - 1
	fmt.Println("Starting ", nthreads, " threads...")
	runtime.GOMAXPROCS(nthreads)
}

func main() {
	init_threads()

	pipeline := processing.NewPipeline()

	timesayer := strings.NewTimeSayer()
	printer := strings.NewPrinter()

	// Connect the network.
	printer.In = timesayer.Out

	// Add the processes to the pipeline.
	pipeline.AddProcesses(timesayer, printer)

	pipeline.Run()
}
