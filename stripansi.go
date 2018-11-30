package stripansi

import (
	"regexp"
	"strings"
)

const ansi = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"
const f = "[<>]"

var re = regexp.MustCompile(ansi)
var reF = regexp.MustCompile(f)

// Strip удаляет ANSI символы
func Strip(str string) string {
	return re.ReplaceAllString(str, "")
}

// StripFullR также даляем символы которые не входят в GTK3 markup
func StripFullR(str string) string {
	str = reF.ReplaceAllString(str, "")
	return re.ReplaceAllString(str, "")
}

// StripFull также даляем символы которые не входят в GTK3 markup
func StripFull(str string) string {
	str = strings.Replace(str, "<", "", -1)
	str = strings.Replace(str, ">", "", -1)
	return re.ReplaceAllString(str, "")
}
