/*
* @Author: scottxiong
* @Date:   2021-09-27 22:29:36
* @Last Modified by:   scottxiong
* @Last Modified time: 2021-09-28 01:10:11
 */
package youtube

import (
	"strconv"
	"testing"
)

func TetsGetConfig(t *testing.T) {
	conf := GetConfig()
	if conf.Width != strconv.Itoa(defaultWidth) || conf.Height == strconv.Itoa(defaultHeight) {
		t.Errorf("GetConfig() failed, got (%s,%s), expected (%d,%d)\n", conf.Width, conf.Height, defaultWidth, defaultHeight)
	}
}

func TestShareFrameWithURL(t *testing.T) {
	url := "https://www.youtube.com/watch?v=hXHrhnt2TEI"
	var expected = `<iframe width="745" height="419" src="https://www.youtube.com/embed/hXHrhnt2TEI" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>`
	res := ShareFrameWithURL(url)
	if res != expected {
		t.Errorf("ShareFrameWithURL(height, width, url) failed, got %s, expected %s\n", res, expected)
	}
}

func TestRetriveYoutubeURLSFromStr(t *testing.T) {
	str := `aa https://www.youtube.com/watch?v=hXHrhnt2TEI bb https://www.youtube.com/watch?v=hXHrhnt2TEI&tt=11 cc https://www.youtube.com/watch?v=hXHrhnt2TEI&tt=33&dd=22`
	res := RetriveYoutubeURLSFromStr(str)
	expected := "https://www.youtube.com/watch?v=hXHrhnt2TEI"
	for _, v := range res {
		if v != expected {
			t.Errorf("RetriveYoutubeURLSFromStr(str) failed, got %v, expected %v\n", v, expected)
		}
	}
}

func TestUpdateWidth(t *testing.T) {
	UpdateWidth(645)
	expected := "645"
	if video.Width != expected {
		t.Errorf("UpdateWidth(str) failed, got %v, expected %v\n", video.Width, expected)
	}
}

func TestParseYTBUrlInRaw(t *testing.T) {
	str := `aa https://www.youtube.com/watch?v=aa bb https://www.youtube.com/watch?v=bb&tt=11 cc https://www.youtube.com/watch?v=cc&dd=22`
	res := ParseYTBUrlInRaw(str)
	expected := `aa <iframe width="645" height="363" src="https://www.youtube.com/embed/aa" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe> bb <iframe width="645" height="363" src="https://www.youtube.com/embed/bb" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe> cc <iframe width="645" height="363" src="https://www.youtube.com/embed/cc" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>`
	if res != expected {
		t.Errorf("ParseYTBUrlInRaw(str) failed, got %v, expected %v\n", res, expected)
	}
}
