### golang颜色转换库

实现hex->rgb,rgb->hex的互转

	package main

	import (
		"fmt"
		"github.com/zouhuigang/zcolor"
	)
	
	func main() {
		color_str := "#0393ff" //颜色代码的值
		r, g, b, err := zcolor.Hex2Rgb(color_str)
		if err != nil {
			fmt.Println(err)
		}
	
		rgb := fmt.Sprintf("RGB(%d,%d,%d)", r, g, b)
		println(rgb)//#0393ff
	
		hexString, _ := zcolor.Rgb2Hex(3, 147, 255)
	
		fmt.Println("#" + hexString) //#0393ff
	
	}
