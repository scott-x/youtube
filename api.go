/*
* @Author: scottxiong
* @Date:   2021-09-27 03:55:47
* @Last Modified by:   scottxiong
* @Last Modified time: 2021-09-27 04:35:42
 */
package youtube

import (
	"regexp"
	"strings"
)

var str = `<iframe width="mywidth" height="myheight" src="myurl" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>`
var youtubeRe = regexp.MustCompile(`(https://www.youtube.com\S*)`)

func ShareFrameWithURL(height, width, url string) string {
	var res string
	res = strings.Replace(str, "mywidth", width, -1)
	res = strings.Replace(res, "myheight", height, -1)
	res = strings.Replace(res, "myurl", url, -1)
	return res
}

func RetriveYoutubeURLSFromStr(str string) []string {
	return youtubeRe.FindAllString(str, -1)
}
