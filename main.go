package main

import (
	"fmt"
	"os"
	"image"
	"image/color"
	"image/png"
	"image/jpeg"
	"log"
	"strings"
	"golang.org/x/term"
)

const channel  int = 4

var (
	fb_height, fb_width int;
)


func getSizeOfScreen()  {
	bytes, err := os.ReadFile("/sys/class/graphics/fb0/virtual_size")
	if err != nil {
		fmt.Printf("Something went wrong: %v\n", err)
		panic(err)
	}
	fmt.Sscanf(string(bytes), "%v,%v\n",  &fb_width, &fb_height)
}


func ShowImgOnFrambeBuffer(img image.Image){
	fb := make([]byte, fb_height * fb_width * channel)
	for y:=img.Bounds().Min.Y; y<img.Bounds().Max.Y; y++ {
		for x:=img.Bounds().Min.X; x<img.Bounds().Max.X; x++ {
			tmp := img.At(x,y)
			col := color.RGBAModel.Convert(tmp).(color.RGBA)
			fb[y * fb_width * channel + x * channel    ] = col.B 
			fb[y * fb_width * channel + x * channel + 1] = col.G
			fb[y * fb_width * channel + x * channel + 2] = col.R
			fb[y * fb_width * channel + x * channel + 3] = col.A
		}
	}
	err := os.WriteFile("/dev/fb0", fb, 0644) //not sure about 0644
	if err!=nil{
		fmt.Printf("Something went wront, err code: %v\n", err)
		panic(err)
	}
}


func fname2fb(fname string) {
	imgfile, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer imgfile.Close()

	var img image.Image;
	fnameParts := strings.Split(fname, ".")  //splits fname with dot . 
	extStr     := strings.ToLower(
		fnameParts[len(fnameParts)-1])       //last item from slice,& make it lowercase


	switch (extStr) {
	case "png":
		img, err = png.Decode(imgfile)
	case "jpeg", "jpg":
		img, err = jpeg.Decode(imgfile)
	default:
		panic(123)
	} 
	

	if err != nil {
		log.Fatal(err)
	}
	ShowImgOnFrambeBuffer(img)
}




func main(){
	// set cursor to the end of line.
	// if not, display image will overwritten with shell's input line with $
	_, height, err := term.GetSize(0)
    if err != nil {
        panic(2)
    }
	fmt.Printf("\033[%d;%dH", height, 0)

	fname := os.Args[1]   // for the command line arguments
	getSizeOfScreen()     // gets screen size
	fname2fb(fname)       //\033[%d;%dH%c
}


