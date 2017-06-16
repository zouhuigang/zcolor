/*
作用：颜色转换
作者：邹慧刚
联系方式：952750120@qq.com
github:github.com/zouhuigang

rgba->hex
hex->rgba
*/
package zcolor

import (
	"errors"
	"strconv"
	"strings"
)

/**
 * 颜色代码转换为RGB
 * input int
 * output int red, green, blue
 **/
func hexToRGB(color int) (red, green, blue int) {
	red = color >> 16
	green = (color & 0x00FF00) >> 8
	blue = color & 0x0000FF
	return
}

func t2x(t int64) string {
	result := strconv.FormatInt(t, 16)
	if len(result) == 1 {
		result = "0" + result
	}
	return result
}

func rgbToHex(r, g, b int) string {
	red, green, blue := t2x(int64(r)), t2x(int64(g)), t2x(int64(b))
	rgbcolor := red + green + blue
	return rgbcolor
}

//color_str可以为：0x3254d5或者#3254d5
func Hex2Rgb(color_str string) (int, int, int, error) {
	color_str = strings.TrimPrefix(color_str, "0x")     //过滤掉16进制前缀
	color_str = strings.TrimPrefix(color_str, "#")      //过滤掉16进制前缀
	color64, err := strconv.ParseInt(color_str, 16, 32) //字串到数据整型
	if err != nil {
		err = errors.New("#" + color_str + ":Hex2Rgb fail ..\n")
		return 0, 0, 0, err
	}
	color32 := int(color64) //类型强转
	r, g, b := hexToRGB(color32)
	return r, g, b, nil
}

func Rgb2Hex(r, g, b int) (string, error) {
	hexString := rgbToHex(r, g, b)
	return hexString, nil
}
