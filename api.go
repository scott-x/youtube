/*
* @Author: scottxiong
* @Date:   2021-09-27 03:55:47
* @Last Modified by:   scottxiong
* @Last Modified time: 2021-09-28 01:04:43
 */
package youtube

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

//global value
var (
	video         Video
	defaultWidth  = 745
	defaultHeight = 419
	str           = `<iframe width="mywidth" height="myheight" src="myurl" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>`
	youtubeRe     = regexp.MustCompile(`(https://www.youtube.com\S*)`)
)

type Video struct {
	Width  string
	Height string
}

func init() {
	video.Width = strconv.Itoa(defaultWidth)   //default width
	video.Height = strconv.Itoa(defaultHeight) //default height
}

func UpdateWidth(width int) {
	video.Width = strconv.Itoa(width)

	x := float64(defaultWidth) / float64(width)
	video.Height = fmt.Sprintf("%.f", math.Ceil(float64(defaultHeight)/x))
}

func GetConfig() *Video {
	return &video
}

func ShareFrameWithURL(url string) string {
	var res string
	res = strings.Replace(str, "mywidth", video.Width, -1)
	res = strings.Replace(res, "myheight", video.Height, -1)
	res = strings.Replace(res, "myurl", url, -1)
	res = strings.Replace(res, "watch?v=", "embed/", -1)
	return res
}

func RetriveYoutubeURLSFromStr(str string) []string {
	str = updateRaw(str)
	return youtubeRe.FindAllString(str, -1)
}

func updateRaw(raw string) string {
	urls := youtubeRe.FindAllString(raw, -1)

	for _, url := range urls {
		if strings.Contains(url, "&") {
			idx := strings.Index(url, "&")
			raw = strings.ReplaceAll(raw, url, url[:idx])
		}
	}
	return raw
}

//parse all youtube url in raw to frame
func ParseYTBUrlInRaw(raw string) string {
	raw = updateRaw(raw)
	urls := RetriveYoutubeURLSFromStr(raw)

	for _, url := range urls {
		raw = strings.ReplaceAll(raw, url, ShareFrameWithURL(url))
	}
	return raw
}
