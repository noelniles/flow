package imageio


import (
	"image"
	//"image/jpeg"
	"fmt"
	//"github.com/noelniles/flow/modules/strings"
	"github.com/noelniles/flow/constants"
	"github.com/noelniles/flow/processing"
)

// ImageViewer takes an image and displays it.
type ImageViewer struct {
	processing.Process
	In  chan image.Image
	//Out chan image.Image
}

// NewImageViewer creates a new image viewer.
func NewImageViewer() *ImageViewer {
	return &ImageViewer{In: make(chan image.Image, constants.BUFSIZE)}
}

// Run the image reader.
func (proc *ImageViewer) Run() {
	fmt.Println("Viewing")

	for image := range proc.In {
		fmt.Println(image)
	}
}