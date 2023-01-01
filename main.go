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
)

const height   int = 1080
const width    int = 1920
const channel  int = 4

var (
	fb [height * width * channel] byte
)


func ShowImgOnFrambeBuffer(img image.Image){
	for y:=img.Bounds().Min.Y; y<img.Bounds().Max.Y; y++ {
		for x:=img.Bounds().Min.X; x<img.Bounds().Max.X; x++ {
			tmp := img.At(x,y)
			col := color.RGBAModel.Convert(tmp).(color.RGBA)
			fb[y * width * channel + x * channel    ] = col.B 
			fb[y * width * channel + x * channel + 1] = col.G
			fb[y * width * channel + x * channel + 2] = col.R
			fb[y * width * channel + x * channel + 3] = col.A
		}
	}

	err := os.WriteFile("/dev/fb0", fb[:], 0644) //not sure about 0644
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
	fname := os.Args[1]   // for the command line arguments
	fname2fb(fname)
}
