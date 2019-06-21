package wallmap

import (
	"encoding/json"
	"fmt"
	"image/png"
	"io/ioutil"
	"os"
)

type wallmap struct {
	wall string
}

var wall wallmap

func init() {
}

func (wm *wallmap) parse(src, dst string, rgb string) (err error) {
	wm.wall = rgb
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Printf("Open png error:%s\n", err.Error())
		return err
	}
	defer srcFile.Close()
	img, err := png.Decode(srcFile)
	if err != nil {
		fmt.Printf("Decode png error:%s\n", err.Error())
		return err
	}
	fmt.Printf("bounds:%v\n", img.Bounds())
	bounds := img.Bounds()
	var bitmap [][]uint32
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		var line []uint32

		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			rgb = fmt.Sprintf("#%02d%02d%02d", r, g, b)
			if rgb == wm.wall {
				line = append(line, 1)
				//img.SetRGBA(x,y, image.RGBA{R:256,G:0,B:0,A:1})
			} else {
				line = append(line, 0)
			}
		}
		bitmap = append(bitmap, line)
	}
	//fmt.Printf("bitmap:%v", bitmap)
	body, err := json.Marshal(bitmap)
	if err != nil {
		fmt.Printf("Marsh JSON error:%s\n", err.Error())
		return err
	}
	err = ioutil.WriteFile(dst, body, 0644)
	if err != nil {
		fmt.Printf("Write error:%s\n", err.Error())
		return err
	}
	return
}

// Parse parse the src png to dst josn
func Parse(src, dst string, rgb string) (err error) {
	return wall.parse(src, dst, rgb)
}
