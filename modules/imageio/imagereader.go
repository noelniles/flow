package imageio

import (
	"image"
	_ "image/jpeg" // Register jpeg decoder
	"fmt"
	"os"
	//"image/jpeg"
	//"github.com/noelniles/flow/modules/strings"
	"github.com/noelniles/flow/constants"
	"github.com/noelniles/flow/processing"
	//"runtime"
)

// ImageReader turns a filename into an image.
type ImageReader struct {
	processing.Process
	In  chan string
	Out chan image.Image
}

// NewImageReader creates a new image reader.
func NewImageReader() *ImageReader {
	return &ImageReader{In: make(chan string, constants.BUFSIZE)}
}

// Run the image reader.
func (proc *ImageReader) Run() {
	fmt.Println("Reading")
	//defer close(proc.Out)

	for filename := range proc.In {
		fmt.Println(filename)
		fh, err := os.Open(filename)
		defer fh.Close()

		if err != nil {
			fmt.Println("Couldn't open the image file")
		}
		imagedata, imagetype, err := image.Decode(fh)

		if err != nil {
			fmt.Println("Couldn't create image. Error ", err)
		}
		fmt.Println(imagedata)
		fmt.Println(imagetype)
		proc.Out <- imagedata
	}
	fmt.Println(proc.In)
}