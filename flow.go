package main

import (
	"fmt"
	"github.com/noelniles/flow/modules/imageio"
	//"github.com/noelniles/flow/modules/strings"
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

	//timesayer := strings.NewTimeSayer()
	//printer := strings.NewPrinter()
	imagereader := imageio.NewImageReader()
	imagereader.In <- "/home/me/Pictures/lena.jpg"
	imageviewer := imageio.NewImageViewer()

	// Connect the network.
	imageviewer.In = imagereader.Out

	// Add the processes to the pipeline.
	pipeline.AddProcesses(imagereader, imageviewer)

	pipeline.Run()
}
