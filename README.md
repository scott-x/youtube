# youtube
share youtube video

### API
- `func ShareFrameWithURL(height, width, url string) string`
- `func RetriveYoutubeURLSFromStr(str string) []string`


### test

```go
func main() {
	s := ShareWithURL("665", "375", "https://www.youtube.com/watch?v=m0xLvtLfwrQ")
	fmt.Println(s)
}
```