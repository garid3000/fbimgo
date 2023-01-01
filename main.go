package main

import (
	"fmt"
	"os"
	"image/color"
	"image/png"
	"log"
)



func main(){
	fname := os.Args[1]   // for the command line arguments

	imgfile, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
		
	}
	defer imgfile.Close()


	img, err := png.Decode(imgfile)
	if err != nil {
		log.Fatal(err)
	}

	const height   int = 1080
	const width    int = 1920
	const channel  int = 4
	var fb [height * width * channel] byte

	for y:=img.Bounds().Min.Y; y<img.Bounds().Max.Y; y++ {
		for x:=img.Bounds().Min.X; x<img.Bounds().Max.X; x++ {
			tmp := img.At(x,y)
			col := color.RGBAModel.Convert(tmp).(color.RGBA)
			fb[y * width * channel + x * channel    ] = col.B
			fb[y * width * channel + x * channel + 1] = col.G
			fb[y * width * channel + x * channel + 2] = col.G
			fb[y * width * channel + x * channel + 3] = col.A
		}
	}

	err = os.WriteFile("/dev/fb0", fb[:], 0644) //not sure about 0644
	if err!=nil{
		fmt.Printf("Something went wront, err code: %v\n", err)
		panic(err)
	}
	
}
