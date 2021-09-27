# youtube

Parse your youtube link to html string, which can be used for sharing.

### API

- `func GetConfig() *Video`
- `func UpdateWidth(width int)`: update the width for the video, default width/height is 745/149.
- `func ShareFrameWithURL(url string) string`: get the sharing html with just a youtube url
- `func RetriveYoutubeURLSFromStr(str string) []string`
- `func ParseYTBUrlInRaw(raw string) string`: parse all youtube url in raw to frame