package output

import "github.com/fatih/color"

func Highlight(text string) string {
	return color.New(color.FgGreen).Sprint(text)
}

func Danger(text string) string {
	return color.New(color.FgRed).Sprint(text)
}

func Error(text string) string {
	return color.New(color.FgWhite, color.BgRed).Sprint(text)
}

func Emphasis(text string) string {
	return color.New(color.FgCyan).Sprint(text)
}

func DarkEmphasis(text string) string {
	return color.New(color.FgBlue).Sprint(text)
}

func Note(text string) string {
	return color.New(color.FgMagenta).Sprint(text)
}

//func Warning(text string) string {
//	return color.New(color.FgYellow).Sprint(text)
//}
